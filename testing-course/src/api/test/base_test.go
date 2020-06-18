package test

import (
	"os"
	"testing"

	"github.com/ck3g/SomeDaysOfGo/testing-course/src/api/app"
	"github.com/mercadolibre/golang-restclient/rest"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	go app.StartApp() // Run the app start in a Go routine to not block the work
	os.Exit(m.Run())
}
