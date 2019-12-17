package nrecho

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/willnewrelic/layouttesting/v3/newrelic"
)

func init() {
	fmt.Println("echo version", echo.Version)
	fmt.Println("agent version", newrelic.Version)
}
