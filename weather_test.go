package weather2

import (
	"testing"
)

const uac = "<your-uac>"

func TestClientCreation(t *testing.T) {
	_, err := NewClient(uac, "badinput", MILES_PER_HOUR)
	if err == nil {
		t.Error("Expected bad temperature unit error")
	}

	_, err = NewClient(uac, CELCIUS, "badinput")
	if err == nil {
		t.Error("Expected bad wind speed unit error")
	}
}

func BenchmarkForecast(b *testing.B) {
	c, err := NewClient(uac, CELCIUS, MILES_PER_HOUR)
	if err != nil {
		b.Error(err)
	}

	w, err := c.Get2DayForecast("90210")
	if err != nil {
		b.Error(err)
	}
	b.Logf("%+v", w)
}
