package endpoints

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

const IndexRoute = "/"

func IndexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Mailcollector is online")
}
