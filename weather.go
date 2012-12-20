// Package weather2 provides API access to the weather2 forecast.
//
// Weather2 Terms of Use (http://www.myweather2.com/developer/ - accessed 19/12/2012):-
//
// Weather2 weather API web service feed is provided for non-commercial and commercial use. 
// The data provided remains the property of Weather2 Ltd and under no circumstances can it be sold or 
// re-distributed. Under these conditions the feeds are free of charge.
//
// Anyone using the feed is required to provide a hyperlink back to www.myweather2.com wherever data from the 
// feed is displayed. The hyperlink text should state 'Weather provided by MyWeather2.com'.
//
// Weather2 reserve the right to cease the distribution of the weather API web service feed at any time and also 
// can require that any user cease use of feed content. Weather2 strictly forbids the on-selling of any data 
// contained within the feeds.
//
// This free API is restricted to a maximum of 500 server requests per day. If in any calendar month if 3 days 
// have over 500 requests we will request you upgrade your account to a paid service. Prices start from as little 
// as £30 per month.
package weather2

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseUrl             = "http://www.myweather2.com/developer/"
	CELCIUS             = "c"
	FAHRENHEIT          = "f"
	METRE_PER_SECOND    = "mps"
	MILES_PER_HOUR      = "mph"
	KILOMETERS_PER_HOUR = "kph"
)

type Weather struct {
	CurrentWeather CurrentWeather `xml:"curren_weather"` // typo in the api
	Forecast       []Forecast     `xml:"forecast"`
}

type CurrentWeather struct {
	Temp        int    `xml:"temp"`
	TempUnit    string `xml:"temp_unit"`
	Wind        Wind   `xml:"wind"`
	Humidity    int    `xml:"humidity"`
	Pressure    int    `xml:"pressure"`
	WeatherText string `xml:"weather_text"`
}

type Forecast struct {
	Date         string          `xml:"date"`
	TempUnit     string          `xml:"temp_unit"`
	DayMaxTemp   int             `xml:"day_max_temp"`
	NightMinTemp int             `xml:"night_min_temp"`
	Day          ForecastWeather `xml:"day"`
	Night        ForecastWeather `xml:"night"`
}

type Wind struct {
	Speed int    `xml:"speed"`
	Dir   string `xml:"dir"`
	Unit  string `xml:"wind_unit"`
}

type ForecastWeather struct {
	Wind        Wind   `xml:"wind"`
	WeatherCode string `xml:"weather_code"`
	WeatherText string `xml:"weather_text"`
}

type Client struct {
	uac      string
	tempUnit string
	windUnit string
}

// Pretty prints the weather including the required backlink.
func (w *Weather) String() string {
	s := w.CurrentWeather.String() + "\n"
	for _, forecast := range w.Forecast {
		s += forecast.String() + "\n"
	}
	s += "Weather provided by www.MyWeather2.com"
	return s
}

func (cw *CurrentWeather) String() string {
	return fmt.Sprintf("Today (%s)\n\tTemp: %d%s :: %s :: Humidity: %dpct :: Pressure: %dmb",
		cw.WeatherText, cw.Temp, cw.TempUnit, cw.Wind.String(), cw.Humidity, cw.Pressure)
}

func (f *Forecast) String() string {
	return fmt.Sprintf("%s\n\tHigh: %d%s Low: %d%s\n\tDay: %s\n\tNight: %s", f.Date, f.DayMaxTemp, f.TempUnit, f.NightMinTemp, f.TempUnit, f.Day.String(), f.Night.String())
}

func (w *Wind) String() string {
	return fmt.Sprintf("Wind: %d%s (Direction: %s)", w.Speed, w.Unit, w.Dir)
}

func (fw *ForecastWeather) String() string {
	return fmt.Sprintf("%s :: %s", fw.WeatherText, fw.Wind.String())
}

// Get2DayForecast queries Weather2 for a two day forecast.
//
// Query can be:-
//   UK Postcode
//   Zip code
//   Lat,Lon
//
// If sucessful a pointer to a Weather struct will be returned with a nil error.
func (c *Client) Get2DayForecast(query string) (*Weather, error) {
	return getWeather(baseUrl+"forecast.ashx?query="+url.QueryEscape(query), c.uac, c.tempUnit, c.windUnit)
}

// Get7DayForecast queries Weather2 for a seven day forecast.
//
// This is only available for one global location, get and set the uref from the
// Developer Zone of your account.
//
// If sucessful a pointer to a Weather struct will be returned with a nil error.
func (c *Client) Get7DayForecast(uref string) (*Weather, error) {
	return getWeather(baseUrl+"weather.ashx?uref="+uref, c.uac, c.tempUnit, c.windUnit)
}

// NewClient returns a client for accessing the Weather2 api.
// Use the contants provided to set the units.
//
// Read the terms of use and get your UAC from http://www.myweather2.com/developer/
func NewClient(uac, tempUnit, windUnit string) (*Client, error) {
	if tempUnit != CELCIUS && tempUnit != FAHRENHEIT {
		return nil, errors.New("Invalid temperature unit")
	}

	if windUnit != METRE_PER_SECOND && windUnit != MILES_PER_HOUR && windUnit != KILOMETERS_PER_HOUR {
		return nil, errors.New("Invalid wind speed unit")
	}
	return &Client{uac, tempUnit, windUnit}, nil
}

func getWeather(addr, uac, tempUnit, windUnit string) (*Weather, error) {
	addr += "&uac=" + uac + "&output=xml&temp_unit=" + tempUnit + "&ws_unit=" + windUnit
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Api errors are returned in plain text, return error if not XML
	if !strings.Contains(string(body), "<?xml version=\"1.0\" encoding=\"UTF-8\"?>") {
		return nil, errors.New(string(body))
	}

	var weather *Weather = new(Weather)
	if err = xml.Unmarshal(body, weather); err != nil {
		return nil, err
	}
	return weather, nil
}
