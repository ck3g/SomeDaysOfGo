package services

import (
	"github.com/ck3g/SomeDaysOfGo/testing-course/src/api/domain/locations"
	"github.com/ck3g/SomeDaysOfGo/testing-course/src/api/providers/locations_provider"
	"github.com/ck3g/SomeDaysOfGo/testing-course/src/api/utils/errors"
)

type locationsService struct{}

type locationsServiceInterface interface {
	GetCountry(countryID string) (*locations.Country, *errors.APIError)
}

var LocationsService locationsServiceInterface

func init() {
	LocationsService = &locationsService{}
}

// GetCountry returns a country by ID
func (s *locationsService) GetCountry(countryID string) (*locations.Country, *errors.APIError) {
	return locations_provider.GetCountry(countryID)
}
