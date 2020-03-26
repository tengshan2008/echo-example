package handlers

import (
	"echo-example/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type CatHandler struct{}

type catGetRequest struct {
	Name    string `json:"name" form:"name" query:"name"`
	Type    string `json:"type" form:"type" query:"type"`
	Sort    string `json:"sort" form:"sort" query:"sort"`
	Fields  string `json:"fields" form:"fields" query:"fields"`
	Page    int64  `json:"page" form:"page" query:"page"`
	PerPage int64  `json:"per_page" form:"per_page" query:"per_page"`
}

type catGetResponse struct {
	Message string `json:"message,omitempty"`
	Data    struct {
		Name string `json:"name,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"data,omitempty"`
}

func (r *catGetRequest) bind(c echo.Context, cat *models.Cat) (err error) {
	if err = c.Bind(r); err != nil {
		return
	}
	if err = c.Validate(r); err != nil {
		return
	}
	cat.Name = r.Name
	cat.Type = r.Type
	var id uint64
	id, err = strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return
	}
	cat.Model.ID = uint(id)
	return
}

func newCatGetResponse(c echo.Context, cat []models.Cat) catGetResponse {
	return catGetResponse{}
}

func (h *CatHandler) Get(c echo.Context) (err error) {
	req := new(catGetRequest)
	cat := new(models.Cat)
	if err = req.bind(c, cat); err != nil {
		return
	}
	var cats []models.Cat
	if cat.ID == 0 {
		cats = cat.ReadMore(req.Sort, req.Fields, req.Page, req.PerPage)
	} else {
		cat.ReadOne(req.Fields)
		cats = append(cats, *cat)
	}

	resp := newCatGetResponse(c, cats)
	return c.JSON(http.StatusOK, resp)
}

type catAddRequest struct {
	Name string `json:"name" form:"name" query:"name"`
	Type string `json:"type" form:"type" query:"type"`
}

func (r *catAddRequest) bind(c echo.Context, cat *models.Cat) (err error) {
	body := c.Request().Body
	defer body.Close()

	if err := json.NewDecoder(body).Decode(r); err != nil {
		log.Fatalf("Failed reading the request body %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	if err = c.Validate(r); err != nil {
		return
	}
	cat.Name = r.Name
	cat.Type = r.Type
	return
}

func (h *CatHandler) Add(c echo.Context) (err error) {
	req := new(catAddRequest)
	cat := new(models.Cat)
	if err = req.bind(c, cat); err != nil {
		return
	}
	if ok := cat.Insert(); !ok {
		return errors.New("insert new cat failed")
	}
	return c.JSON(http.StatusCreated, req)
}

type catDeleteRequest struct{}

func (r *catDeleteRequest) bind(c echo.Context, cat *models.Cat) (err error) {
	var id uint64
	id, err = strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return
	}
	cat.Model.ID = uint(id)
	return
}

func (h *CatHandler) Delete(c echo.Context) (err error) {
	req := new(catDeleteRequest)
	cat := new(models.Cat)
	if err = req.bind(c, cat); err != nil {
		return
	}
	cat.Delete()
	return c.JSON(http.StatusOK, req)
}
