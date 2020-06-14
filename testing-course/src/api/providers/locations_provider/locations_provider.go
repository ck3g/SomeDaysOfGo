package locations_provider

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ck3g/SomeDaysOfGo/testing-course/src/api/domain/locations"
	"github.com/ck3g/SomeDaysOfGo/testing-course/src/api/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

const (
	urlGetCountry = "https://api.mercadolibre.com/countries/%s"
)

// GetCountry retrieves a country by ID
func GetCountry(countryID string) (*locations.Country, *errors.APIError) {
	res := rest.Get(fmt.Sprintf(urlGetCountry, countryID))
	if res == nil || res.Response == nil {
		return nil, &errors.APIError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid resetclient error when getting country %s", countryID),
		}
	}

	if res.StatusCode > 299 {
		var apiErr errors.APIError
		if err := json.Unmarshal(res.Bytes(), &apiErr); err != nil {
			return nil, &errors.APIError{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid interface when getting country %s", countryID),
			}
		}
		return nil, &apiErr
	}

	var result locations.Country
	if err := json.Unmarshal(res.Bytes(), &result); err != nil {
		return nil, &errors.APIError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal country data for %s", countryID),
		}
	}

	return &result, nil
}
