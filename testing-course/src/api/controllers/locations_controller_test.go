package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ck3g/SomeDaysOfGo/testing-course/src/api/domain/locations"
	"github.com/ck3g/SomeDaysOfGo/testing-course/src/api/services"
	"github.com/ck3g/SomeDaysOfGo/testing-course/src/api/utils/errors"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

var getCountryFn func(countryID string) (*locations.Country, *errors.APIError)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

type locationsServiceMock struct {
}

func (s *locationsServiceMock) GetCountry(countryID string) (*locations.Country, *errors.APIError) {
	return getCountryFn(countryID)
}

func TestGetCountry_Integration_NotFound(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/DE",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Country not found", "error": "not_found", "status": 404, "cause": []}`,
	})

	res := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(res)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "DE"},
	}
	GetCountry(c)

	assert.EqualValues(t, http.StatusNotFound, res.Code)

	var apiErr errors.APIError
	err := json.Unmarshal(res.Body.Bytes(), &apiErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "Country not found", apiErr.Message)
}

func TestGetCountry_Mocked_NotFound(t *testing.T) {
	getCountryFn = func(countryID string) (*locations.Country, *errors.APIError) {
		return nil, &errors.APIError{Status: http.StatusNotFound, Message: "Country not found"}
	}
	services.LocationsService = &locationsServiceMock{}

	res := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(res)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "DE"},
	}
	GetCountry(c)

	assert.EqualValues(t, http.StatusNotFound, res.Code)

	var apiErr errors.APIError
	err := json.Unmarshal(res.Body.Bytes(), &apiErr)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "Country not found", apiErr.Message)
}

func TestGetCountry_Mocked_NoError(t *testing.T) {
	getCountryFn = func(countryID string) (*locations.Country, *errors.APIError) {
		return &locations.Country{ID: "DE", Name: "Germany"}, nil
	}
	services.LocationsService = &locationsServiceMock{}

	res := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(res)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "DE"},
	}
	GetCountry(c)

	assert.EqualValues(t, http.StatusOK, res.Code)

	var country locations.Country
	err := json.Unmarshal(res.Body.Bytes(), &country)
	assert.Nil(t, err)
	assert.EqualValues(t, "DE", country.ID)
	assert.EqualValues(t, "Germany", country.Name)
}
