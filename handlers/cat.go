package handlers

import (
	"echo-example/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func GetCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")
	dataType := c.Param("data")
	if dataType == "string" {
		desc := fmt.Sprintf("your cat name is: %s\nand cat type is: %s\n", catName, catType)
		return c.String(http.StatusOK, desc)
	} else if dataType == "json" {
		cat := models.Cat{
			Name: catName,
			Type: catType,
		}
		return c.JSON(http.StatusOK, cat)
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Please specify the data type as String or JSON",
		})
	}
}

func AddCat(c echo.Context) error {
	cat := models.Cat{}

	body := c.Request().Body
	defer body.Close()

	if err := json.NewDecoder(body).Decode(&cat); err != nil {
		log.Fatalf("Failed reading the request body %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}

	log.Printf("this is your cat %#v", cat)
	return c.String(http.StatusOK, "We got your Cat!!!")
}
