package app

import "github.com/ck3g/SomeDaysOfGo/testing-course/src/api/controllers"

func mapUrls() {
	router.GET("/locations/countries/:country_id", controllers.GetCountry)
}
