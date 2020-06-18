package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ck3g/SomeDaysOfGo/testing-course/src/api/domain/locations"
	"github.com/ck3g/SomeDaysOfGo/testing-course/src/api/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestGetCountries_NotFound(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/DEU",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Country not found", "error": "not_found", "status": 404, "cause": []}`,
	})

	res, err := http.Get("http://localhost:8080/locations/countries/DEU")

	bytes, _ := ioutil.ReadAll(res.Body)

	assert.Nil(t, err)
	assert.NotNil(t, res)

	var apiErr errors.APIError
	err = json.Unmarshal(bytes, &apiErr)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, res.StatusCode)
	assert.EqualValues(t, "not_found", apiErr.Error)
	assert.EqualValues(t, "Country not found", apiErr.Message)
}

func TestGetCountries_OK(t *testing.T) {
	successfulResponse := `
	{
		"id": "DE",
		"name": "Germany",
		"locale": "en_US",
		"currency_id": "USD",
		"decimal_separator": ".",
		"thousands_separator": ",",
		"time_zone": "GMT+02:00",
		"geo_information": {},
		"states": [
			{ "id": "DE-NW", "name": "North Rhine West" },
			{ "id": "DE-BW", "name": "Baden-Wurtemberg" },
			{ "id": "DE-BY", "name": "Baviera" },
			{ "id": "DE-BB", "name": "Brandeburgo" },
			{ "id": "DE-HB", "name": "Bremen" },
			{ "id": "DE-HH", "name": "Hamburgo" },
			{ "id": "DE-HE", "name": "Hesse" },
			{ "id": "DE-NI", "name": "Lower Saxony" },
			{ "id": "DE-MV", "name": "Mecklenburg-Vorpo" },
			{ "id": "DE-RP", "name": "Rhineland-Palatinate" },
			{ "id": "DE-SL", "name": "Saarland" },
			{ "id": "DE-SN", "name": "Saxony" },
			{ "id": "DE-ST", "name": "Sajonia-Anhalt" },
			{ "id": "DE-SH", "name": "Schleswig-Holstein" },
			{ "id": "DE-TH", "name": "Thuringia" }
		]
		}
	`

	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/DE",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     successfulResponse,
	})

	res, err := http.Get("http://localhost:8080/locations/countries/DE")

	bytes, _ := ioutil.ReadAll(res.Body)

	assert.Nil(t, err)
	assert.NotNil(t, res)

	var country locations.Country
	err = json.Unmarshal(bytes, &country)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusOK, res.StatusCode)
	assert.EqualValues(t, "Germany", country.Name)
	assert.EqualValues(t, "GMT+02:00", country.TimeZone)
	assert.EqualValues(t, 15, len(country.States))
}
