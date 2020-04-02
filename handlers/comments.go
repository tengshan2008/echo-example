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

type CatCommentHandler struct{}

type catCommentGetOneRequest struct {
	Fields string `json:"fields" form:"fields" query:"fields"`
}

func (r *catCommentGetOneRequest) bind(c echo.Context, comment *models.Comment) (err error) {
	var id uint64
	var cid uint64
	id, err = strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return
	}
	cid, err = strconv.ParseUint(c.Param("cid"), 10, 64)
	if err != nil {
		return
	}
	comment.ID = uint(cid)
	comment.CatID = uint(id)
	return
}

type catCommentGetResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func newCatCommentGetResponse(c echo.Context, data interface{}) (r catCommentGetResponse) {
	r.Message = "ok"
	r.Data = data
	return
}

func (h *CatCommentHandler) GetOne(c echo.Context) (err error) {
	req := new(catCommentGetOneRequest)
	comment := new(models.Comment)
	if req.bind(c, comment); err != nil {
		return
	}
	comment.ReadOne(req.Fields)

	resp := newCatGetResponse(c, comment)
	return c.JSON(http.StatusOK, resp)
}

type catCommentGetMoreRequest struct {
	Author  string `json:"author" form:"author" query:"author"`
	Content string `json:"content" form:"content" query:"content"`
	Search  string `json:"search" form:"search" query:"search"`
	Sort    string `json:"sort" form:"sort" query:"sort"`
	Fields  string `json:"fields" form:"fields" query:"fields"`
	Page    int64  `json:"page" form:"page" query:"page"`
	PerPage int64  `json:"per_page" form:"per_page" query:"per_page"`
}

func (r *catCommentGetMoreRequest) bind(c echo.Context, comment *models.Comment) (err error) {
	if err = c.Bind(r); err != nil {
		return
	}
	comment.Author = r.Author
	comment.Content = r.Content
	return
}

func (h *CatCommentHandler) GetMore(c echo.Context) (err error) {
	req := new(catCommentGetMoreRequest)
	comment := new(models.Comment)
	if err = req.bind(c, comment); err != nil {
		return
	}
	comments := comment.ReadMore(req.Search, req.Sort, req.Fields, req.Page, req.PerPage)
	resp := newCatCommentGetResponse(c, comments)
	return c.JSON(http.StatusOK, resp)
}

type catCommentAddRequest struct {
	Author  string
	Content string
	CatID   uint
}

func (r *catCommentAddRequest) bind(c echo.Context, comment *models.Comment) (err error) {
	body := c.Request().Body
	defer body.Close()

	if err := json.NewDecoder(body).Decode(r); err != nil {
		log.Fatalf("Failed reading the request body %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	comment.Author = r.Author
	comment.Content = r.Content
	comment.CatID = r.CatID
	return
}

func (h *CatCommentHandler) Add(c echo.Context) (err error) {
	req := new(catCommentAddRequest)
	comment := new(models.Comment)
	if err = req.bind(c, comment); err != nil {
		return
	}
	if notExist := comment.Insert(); notExist {
		return errors.New("insert new cat failed")
	}
	return c.JSON(http.StatusCreated, req)
}

type catCommentDeleteRequest struct{}

func (r *catCommentDeleteRequest) bind(c echo.Context, comment *models.Comment) (err error) {
	var cid uint64
	cid, err = strconv.ParseUint(c.Param("cid"), 10, 64)
	if err != nil {
		return
	}
	comment.ID = uint(cid)
	return
}

func (h *CatCommentHandler) Delete(c echo.Context) (err error) {
	req := new(catCommentDeleteRequest)
	comment := new(models.Comment)
	if err = req.bind(c, comment); err != nil {
		return
	}
	if notExist := comment.Delete(); !notExist {
		return errors.New("delete record failed")
	}
	return c.JSON(http.StatusOK, req)
}
