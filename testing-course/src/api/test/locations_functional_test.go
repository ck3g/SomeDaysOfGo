package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

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
