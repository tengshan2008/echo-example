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

// get cat handlers
// include get one or get mutiple

type catGetOneRequest struct {
	Fields string `json:"fields" form:"fields" query:"fields"`
}

func (r *catGetOneRequest) bind(c echo.Context, cat *models.Cat) (err error) {
	var id uint64
	id, err = strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return
	}
	cat.ID = uint(id)
	return
}

type catGetResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func newCatGetResponse(c echo.Context, data interface{}) (r catGetResponse) {
	r.Message = "ok"
	r.Data = data
	return
}

func (h *CatHandler) GetOne(c echo.Context) (err error) {
	req := new(catGetOneRequest)
	cat := new(models.Cat)
	if err = req.bind(c, cat); err != nil {
		return
	}

	cat.ReadOne(req.Fields)

	resp := newCatGetResponse(c, cat)
	return c.JSON(http.StatusOK, resp)
}

type catGetMoreRequest struct {
	Name    string `json:"name" form:"name" query:"name"`
	Type    string `json:"type" form:"type" query:"type"`
	Search  string `json:"search" form:"search" query:"search"`
	Sort    string `json:"sort" form:"sort" query:"sort"`
	Fields  string `json:"fields" form:"fields" query:"fields"`
	Page    int64  `json:"page" form:"page" query:"page"`
	PerPage int64  `json:"per_page" form:"per_page" query:"per_page"`
}

func (r *catGetMoreRequest) bind(c echo.Context, cat *models.Cat) (err error) {
	if err = c.Bind(r); err != nil {
		return
	}
	cat.Name = r.Name
	cat.Type = r.Type
	return
}

func (h *CatHandler) GetMore(c echo.Context) (err error) {
	req := new(catGetMoreRequest)
	cat := new(models.Cat)
	if err = req.bind(c, cat); err != nil {
		return
	}
	cats := cat.ReadMore(req.Search, req.Sort, req.Fields, req.Page, req.PerPage)
	resp := newCatGetResponse(c, cats)
	return c.JSON(http.StatusOK, resp)
}

func (h *CatHandler) V2GetMore(c echo.Context) (err error) {
	return c.String(http.StatusOK, "coming soon!")
}

// add new cat handler

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
	if notExist := cat.Insert(); notExist {
		return errors.New("insert new cat failed")
	}
	return c.JSON(http.StatusCreated, req)
}

// delete cat handler

type catDeleteRequest struct{}

func (r *catDeleteRequest) bind(c echo.Context, cat *models.Cat) (err error) {
	var id uint64
	id, err = strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return
	}
	cat.ID = uint(id)
	return
}

func (h *CatHandler) Delete(c echo.Context) (err error) {
	req := new(catDeleteRequest)
	cat := new(models.Cat)
	if err = req.bind(c, cat); err != nil {
		return
	}
	if notExist := cat.Delete(); !notExist {
		return errors.New("delete record failed")
	}
	return c.JSON(http.StatusOK, req)
}
