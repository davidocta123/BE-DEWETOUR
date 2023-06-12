package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	UserRoutes(e)
	AutRoutes(e)
	CountryRouter(e)
	TripRouter(e)
	TransactionRoutes(e)
}
