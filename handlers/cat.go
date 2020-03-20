package handlers

import (
	"echo-example/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type CatHandler struct {
	Cat models.Cat
}

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
	cat.ID, err = strconv.ParseInt(c.Param("id"), 10, 64)
	return
}

func newCatGetResponse(c echo.Context, cat []models.Cat) catGetResponse {
	return catGetResponse{}
}

func (h *CatHandler) Get(c echo.Context) (err error) {
	req := new(catGetRequest)
	if err = req.bind(c, &h.Cat); err != nil {
		return
	}
	var cats []models.Cat
	if h.Cat.ID == 0 {
		cats = h.Cat.ReadMore(req.Sort, req.Fields, req.Page, req.PerPage)
	} else {
		h.Cat.ReadOne(req.Fields)
		cats = append(cats, h.Cat)
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
	if err = req.bind(c, &h.Cat); err != nil {
		return
	}
	if err = h.Cat.PreInsert(); err != nil {
		return
	}
	return c.JSON(http.StatusCreated, req)
}
