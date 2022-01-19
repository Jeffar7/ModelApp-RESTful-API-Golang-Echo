package controller

import (
	"myapp/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FetchAllCatalogue(c echo.Context) error {

	result, err := models.FetchAllCatalogue()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreModelCatalogues(c echo.Context) error {
	Model_name := c.FormValue("Model_name")
	Image_url := c.FormValue("Image_url")
	Description := c.FormValue("Description")

	result, err := models.StoreModelCatalogues(Model_name,Image_url,Description)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func FetchIdSchedule(c echo.Context) error{

	Model_id := c.Param("model_id")
	conv_id, err := strconv.Atoi(Model_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result , err := models.FetchScheduleByModelId(conv_id) 
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	
	return c.JSON(http.StatusOK,result)
}

func StoreModelSchedules(c echo.Context) error {
	Model_id := c.FormValue("Model_id")
	Schedule_date := c.FormValue("Schedule_date")
	conv_id, err := strconv.Atoi(Model_id)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.StoreModelSchedules(conv_id,Schedule_date)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	return c.JSON(http.StatusOK,result)
}

func FetchAllModelById(c echo.Context) error {
	Model_id := c.Param("model_id")
	conv_id, err := strconv.Atoi(Model_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	result, err := models.FetchAllModelById(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK,result)
}

func UpdateModelCatalogues(c echo.Context) error {
	id := c.Param("model_id")
	Model_name := c.FormValue("Model_name")
	Image_url := c.FormValue("Image_url")
	Description := c.FormValue("Description")

	conv_id, err :=strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	result, err := models.UpdateModelCatalogues(conv_id,Model_name,Image_url,Description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	return c.JSON(http.StatusOK, result)
}

func UpdateScheduleByScheduleID(c echo.Context) error{
	id := c.Param("id")
	Model_id := c.FormValue("Model_id")
	Schedule_date := c.FormValue("Schedule_date")
	
	conv_id, err :=strconv.Atoi(id)
	conv_model_id, err :=strconv.Atoi(Model_id)

	
	result, err := models.UpdateScheduleByScheduleID(conv_id,conv_model_id,Schedule_date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK,result)

}

func DeleteModelSchedule(c echo.Context) error {
	id := c.Param("model_id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	result, err := models.DeleteModelSchedule(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
func DeleteModelCatalogues(c echo.Context) error {
	id := c.Param("model_id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	result, err := models.DeleteModelCatalogues(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}