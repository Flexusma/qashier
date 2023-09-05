package main

import (
	"flexusma.de/quashier-server/database"
	apiRouters "flexusma.de/quashier-server/routers"
	"flexusma.de/quashier-server/util"
	"github.com/labstack/echo-jwt/v4"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

var echoInstance *echo.Echo
var Logger echo.Logger

var routers []apiRouters.APIRouter = []apiRouters.APIRouter{
	new(apiRouters.IndexRouter),
	new(apiRouters.ItemRouter),
}

func main() {
	//Init Echo
	echoInstance = echo.New()
	echoInstance.HideBanner = true
	//get Logger
	Logger = echoInstance.Logger

	//PRINT INIT
	println("\n\n                       __     _            \n  ____ _ ____ _ _____ / /_   (_)___   _____\n / __ `// __ `// ___// __ \\ / // _ \\ / ___/\n/ /_/ // /_/ /(__  )/ / / // //  __// /    \n\\__, / \\__,_//____//_/ /_//_/ \\___//_/     \n  /_/        server  " + util.IfThenStr(util.Dev, "dev", "    ") + "  v" + util.Version + "   \n\n")

	//IF Dev set debug level
	if util.Dev {
		Logger.SetLevel(log.DEBUG)
	}
	println("Connecting to database...")
	database.InitDB()

	println("Registering routers")
	//adding middleware
	echoInstance.Use(echojwt.JWT([]byte("wTDGpyAsJwEzoiz8xQGz5NM2afl0JB7jGmX005ulfP21BScn/P3WRVJH/w8CFf9BWT8RrS++M3kRZfpRkYE7dw==")))

	//init routers
	for _, router := range routers {
		router.Init(echoInstance)
	}

	echoInstance.Use(middleware.Logger())
	echoInstance.Logger.Fatal(echoInstance.Start(":3000"))

}
