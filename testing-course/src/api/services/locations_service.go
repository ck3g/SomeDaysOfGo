package services

import (
	"github.com/ck3g/SomeDaysOfGo/testing-course/src/api/domain/locations"
	"github.com/ck3g/SomeDaysOfGo/testing-course/src/api/providers/locations_provider"
	"github.com/ck3g/SomeDaysOfGo/testing-course/src/api/utils/errors"
)

// GetCountry returns a country by ID
func GetCountry(countryID string) (*locations.Country, *errors.APIError) {
	return locations_provider.GetCountry(countryID)
}
