package locations_provider

import (
	"net/http"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestGetCountry_RESTClientError(t *testing.T) {
	// Unfortunately rest.StartMockupServer() doesn't work with fresh go lang versions
	// https://github.com/mercadolibre/golang-restclient/issues/5
	// Wait for https://github.com/mercadolibre/golang-restclient/pull/2/files to be merged, or update the files
	rest.StartMockupServer()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/DE",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: -1,
	})

	country, err := GetCountry("DE")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid resetclient error when getting country DE", err.Message)
}

func TestGetCountry_CountryNotFound(t *testing.T) {
	rest.StartMockupServer()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/DE",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Country not found", "error": "not_found", "status": 404, "cause": []}`,
	})

	country, err := GetCountry("DE")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountry_InvalidErrorInterface(t *testing.T) {
	rest.StartMockupServer()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/DE",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Country not found", "error": "not_found", "status": "404, "cause": []}`,
	})

	country, err := GetCountry("DE")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid interface when getting country DE", err.Message)
}

func TestGetCountry_InvalidJSONResponse(t *testing.T) {
	rest.StartMockupServer()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/DE",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 1, "name": "Germany"}`,
	})

	country, err := GetCountry("DE")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal country data for DE", err.Message)
}

func TestGetCountry_NoError(t *testing.T) {
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
	rest.StartMockupServer()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/DE",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     successfulResponse,
	})

	country, err := GetCountry("DE")

	assert.NotNil(t, country)
	assert.Nil(t, err)
	assert.EqualValues(t, "DE", country.ID)
	assert.EqualValues(t, "Germany", country.Name)
	assert.EqualValues(t, "GMT+02:00", country.TimeZone)
	assert.EqualValues(t, 15, len(country.States))
}
