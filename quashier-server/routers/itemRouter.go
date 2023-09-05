package routers

import (
	"flexusma.de/quashier-server/database/models"
	"flexusma.de/quashier-server/util"
	"github.com/labstack/echo/v4"
)

const itemRouterPath = "/items/"

type ItemRouter struct {
	e *echo.Echo
}

func (x ItemRouter) Init(e *echo.Echo) {
	//TODO implement me
	x.e = e
	// '/'
	e.GET(itemRouterPath, getItems)
}

//Routes

// Hello world route
func getItems(c echo.Context) error {
	var items []models.Item
	query := models.AllItems(items)
	if query.Error != nil {
		err := query.Error.Error()
		return c.JSON(500, util.NewError(err, "UNKNOWN"))
	}

	return c.JSON(200, util.NewSuccess("Success getting all items", items))
}
