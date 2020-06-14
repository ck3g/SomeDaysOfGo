package locations_provider

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCountry_RESTClientError(t *testing.T) {
	country, err := GetCountry("DE")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid resetclient error when getting country DE", err.Message)
}

func TestGetCountry_CountryNotFound(t *testing.T) {
	country, err := GetCountry("DE")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountry_InvalidErrorInterface(t *testing.T) {
	country, err := GetCountry("DE")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid interface when getting country DE", err.Message)
}

func TestGetCountry_InvalidJSONResponse(t *testing.T) {
	country, err := GetCountry("DE")

	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal country data for DE", err.Message)
}

func TestGetCountry_NoError(t *testing.T) {
	country, err := GetCountry("DE")

	assert.NotNil(t, country)
	assert.Nil(t, err)
	assert.EqualValues(t, "DE", country.ID)
	assert.EqualValues(t, "Germany", country.Name)
	assert.EqualValues(t, "GMT+02:00", country.TimeZone)
	assert.EqualValues(t, 15, len(country.States))
}
