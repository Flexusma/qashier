package routers

import "github.com/labstack/echo/v4"

const routerPath = "/"

type IndexRouter struct {
	e *echo.Echo
}

func (x IndexRouter) Init(e *echo.Echo) {
	//TODO implement me
	x.e = e
	// '/'
	e.GET(routerPath, helloWorld)
}

//Routes

// Hello world route
func helloWorld(c echo.Context) error {

	return c.JSON(200, "hello there")
}
