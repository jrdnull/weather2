weather2
========

weather2 provides easy access to the [Weather2](http://www.myweather2.com) api, to use this 
api you must sign up for a key and agree to the terms of service which can be done [here](http://www.myweather2.com/developer/).

The 7 day forecast is limited to one global location by using a reference to a location which can be generated in 
the Developer Zone of your account.

Installation
------------

With Go 1 and git installed:

    go get github.com/jrdnull/weather2

This will download, compile, and install the package into your `$GOROOT`
directory hierarchy. Alternatively, you can import it into a
project:

    import "github.com/jrdnull/weather2"

and when you build that project with `go build`, it will be
downloaded and installed automatically.

Usage
-----

To get the two day forecast for 90210, with the temp in celcius and wind speed in miles per hour:
```go
client, err := weather2.NewClient("6-f827v6vb", weather2.CELCIUS, weather2.MILES_PER_HOUR)
if err != nil {
	fmt.Println(err)
}

weather, err := client.Get2DayForecast("90210")
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println(weather)
}
```
Outputs:
> Today (Clear skies)
>         Temp: 11c :: Wind: 2mph (Direction: NNE) :: Humidity: 22pct :: Pressure: 1023mb
> 2012-12-20
>         High: 18c Low: 8c
>         Day: Sunny skies :: Wind: 2mph (Direction: E)
>         Night: Clear skies :: Wind: 2mph (Direction: N)
> 2012-12-21
>         High: 19c Low: 9c
>         Day: Sunny skies :: Wind: 2mph (Direction: E)
>         Night: Clear skies :: Wind: 2mph (Direction: N)
> Weather provided by www.MyWeather2.com

Check the generated go doc for more info.

License
-------

weather2 is distributed under the Simplified BSD License:

> Copyright © 2012 Jordon Smith. All rights reserved.
> 
> Redistribution and use in source and binary forms, with or without modification, are
> permitted provided that the following conditions are met:
> 
>    1. Redistributions of source code must retain the above copyright notice, this list of
>       conditions and the following disclaimer.
> 
>    2. Redistributions in binary form must reproduce the above copyright notice, this list
>       of conditions and the following disclaimer in the documentation and/or other materials
>       provided with the distribution.
> 
> THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER ``AS IS'' AND ANY EXPRESS OR IMPLIED
> WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
> FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> OR
> CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
> CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
> SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
> ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
> NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
> ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
> 
> The views and conclusions contained in the software and documentation are those of the
> authors and should not be interpreted as representing official policies, either expressed
> or implied, of the copyright holder.
