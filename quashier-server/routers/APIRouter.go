package routers

import "github.com/labstack/echo/v4"

type APIRouter interface {
	Init(echo *echo.Echo)
}
