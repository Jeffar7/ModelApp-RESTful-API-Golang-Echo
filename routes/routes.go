package routes

import (
	"myapp/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo{
	e:= echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.GET("/models/list", controller.FetchAllCatalogue)
	e.POST("/models/create", controller.StoreModelCatalogues)
	e.GET("/models/schedules/:model_id", controller.FetchIdSchedule)
	e.POST("/models/schedules/create", controller.StoreModelSchedules)
	e.GET("/models/:model_id", controller.FetchAllModelById)
	e.PUT("/models/update/:model_id", controller.UpdateModelCatalogues)
	e.PUT("/models/schedules/:id", controller.UpdateScheduleByScheduleID)
	e.DELETE("/models/schedules/:model_id", controller.DeleteModelSchedule)
	e.DELETE("/models/:model_id", controller.DeleteModelCatalogues)
	
	return e;
}