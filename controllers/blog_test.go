package controller

import (
	"OrmGo/config"
	model "OrmGo/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestBlogController_Index(t *testing.T) {

	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.BlogModel{}
	mdl.Init(gorm)
	var ctl = BlogController{}
	ctl.InitBlogController(mdl)

	req := httptest.NewRequest(http.MethodGet, "/blogs", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	err := ctl.Index()(c)

	if assert.NoError(t, err) {
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		blog := response["blog"].([]interface{})

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success get all blogs", response["message"])
		assert.NotNil(t, blog)
	}
}

func TestBlogController_Create(t *testing.T) {

	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.BlogModel{}
	mdl.Init(gorm)
	var ctl = BlogController{}
	ctl.InitBlogController(mdl)

	reqBody := `{
		"ID": 20,
		"title": "Title A",
		"content": "Content A",
		"user": 2
	}`

	req := httptest.NewRequest(http.MethodPost, "/blogs", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	err := ctl.Create()(c)

	if assert.NoError(t, err) {
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success create new blog", response["message"])

		blog := response["blog"].(map[string]interface{})
		assert.NotNil(t, blog)
		assert.Equal(t, "Title A", blog["title"])
		assert.Equal(t, "Content A", blog["content"])
		assert.Equal(t, float64(2), blog["user"])

	}
}

func TestBlogController_Show(t *testing.T) {
	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.BlogModel{}
	mdl.Init(gorm)
	var ctl = BlogController{}
	ctl.InitBlogController(mdl)

	req := httptest.NewRequest(http.MethodGet, "/blogs/20", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("20")
	err := ctl.Show()(c)

	if assert.NoError(t, err) {
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success get blog by id", response["message"])

		blog := response["blog"].(map[string]interface{})
		assert.NotNil(t, blog)
		assert.Equal(t, "Title A", blog["title"])
		assert.Equal(t, "Content A", blog["content"])
		assert.Equal(t, float64(2), blog["user"])
	}
}

func TestBlogController_Update(t *testing.T) {
	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.BlogModel{}
	mdl.Init(gorm)
	var ctl = BlogController{}
	ctl.InitBlogController(mdl)

	reqBody := `{
		"ID": 20,
		"title": "Title A",
		"content": "Content A",
		"user": 2
	}`

	req := httptest.NewRequest(http.MethodPut, "/blogs/20", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("20")

	err := ctl.Update()(c)

	if assert.NoError(t, err) {
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "blog updated successfully", response["message"])

		blog := response["blog"].(map[string]interface{})
		assert.NotNil(t, blog)
		assert.Equal(t, "Title A", blog["title"])
		assert.Equal(t, "Content A", blog["content"])
		assert.Equal(t, float64(2), blog["user"])
	}
}

func TestBlogController_Delete(t *testing.T) {
	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.BlogModel{}
	mdl.Init(gorm)
	var ctl = BlogController{}
	ctl.InitBlogController(mdl)

	req := httptest.NewRequest(http.MethodDelete, "/books/20", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("20")

	err := ctl.Delete()(c)
	if assert.NoError(t, err) {
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "blog deleted successfully", response["message"])

		blog := response["blog"].(map[string]interface{})
		assert.NotNil(t, blog)
		assert.Equal(t, "Title A", blog["title"])
		assert.Equal(t, "Content A", blog["content"])
		assert.Equal(t, float64(2), blog["user"])
	}
}
