package book

import (
	"flag"
	"fmt"
	"testing"
)

var integration = flag.Bool("integration", false, "Use to run integration tests or not")
var apiKey = flag.String("apikey", "", "Pass your open weather map api key")

func init() {
	flag.Parse()
}


func Test_Integration_OpenWeatherMap_GetWeatherByCityName(t *testing.T) {
	if !*integration {
		t.SkipNow()
	}

	if *apiKey == "" {
		t.Error("Not api key passed")
	}

	weatherMap := CurrentWeatherData{*apiKey}

	weather, err := weatherMap.GetByCityAndCountryCode("Madrid", "ES")
	if err != nil {
		t.Fatal(err)
	}

	if weather.Coord.Lon != -3.7 {
		t.Errorf("Lon was not -3.7 as expected. Lon=%f", weather.Coord.Lon)
	}

	fmt.Printf("Temperature in Madrid is %f celsius\n", weather.Main.Temp-273.15)
}

func Test_Integration_OpenWeatherMap_GetWeatherByGeographicalCoordinates(t *testing.T) {
	if !*integration {
		t.SkipNow()
	}

	if *apiKey == "" {
		t.Error("Not api key passed")
	}

	weatherMap := CurrentWeatherData{*apiKey}

	weather, err := weatherMap.GetByGeoCoordinates(-3.7, 40.42)
	if err != nil {
		t.Fatal(err)
	}

	if weather.Cod != 200 {
		t.Errorf("Cod was not 200 as expected. Code: %d\n", weather.Cod)
	}
}


func TestOpenWeatherMap_responseParser(t *testing.T) {
	r := getMockData()
	openWeatherMap := CurrentWeatherData{APIkey: ""}

	weather, err := openWeatherMap.responseParser(r)
	if err != nil {
		t.Fatal(err)
	}

	if weather.ID != 3117735 {
		t.Errorf("Madrid id is 3117735, not %d\n", weather.ID)
	}
}
