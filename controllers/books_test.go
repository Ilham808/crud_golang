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

func TestBookController_Index(t *testing.T) {

	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.BooksModel{}
	mdl.Init(gorm)
	var ctl = BookController{}
	ctl.InitBookController(mdl)

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	err := ctl.Index()(c)

	if assert.NoError(t, err) {
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		books := response["books"].([]interface{})

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success get all books", response["message"])
		assert.NotNil(t, books)
	}
}

func TestBookController_Create(t *testing.T) {

	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.BooksModel{}
	mdl.Init(gorm)
	var ctl = BookController{}
	ctl.InitBookController(mdl)

	reqBody := `{
		"ID": 20,
		"judul": "Buku B",
		"penulis": "Penulis B",
		"penerbit": "Penerbit B"
	}`

	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	err := ctl.Create()(c)

	if assert.NoError(t, err) {
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success create new book", response["message"])

		books := response["book"].(map[string]interface{})
		assert.NotNil(t, books)
		assert.Equal(t, "Buku B", books["judul"])
		assert.Equal(t, "Penulis B", books["penulis"])

	}
}

func TestBookController_Show(t *testing.T) {

	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.BooksModel{}
	mdl.Init(gorm)
	var ctl = BookController{}
	ctl.InitBookController(mdl)

	req := httptest.NewRequest(http.MethodGet, "/books/20", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("20")
	err := ctl.Show()(c)

	if assert.NoError(t, err) {
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		books := response["book"].(map[string]interface{})
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success get book by id", response["message"])
		assert.NotNil(t, books)
		assert.Equal(t, "Buku B", books["judul"])
		assert.Equal(t, "Penulis B", books["penulis"])
	}
}

func TestBookController_Update(t *testing.T) {
	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.BooksModel{}
	mdl.Init(gorm)
	var ctl = BookController{}
	ctl.InitBookController(mdl)

	reqBody := `{
		"judul": "Buku B",
		"penulis": "Penulis B",
		"penerbit": "Penerbit B"
	}`

	req := httptest.NewRequest(http.MethodPut, "/books/20", strings.NewReader(reqBody))
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
		assert.Equal(t, "Book updated successfully", response["message"])

		book := response["book"].(map[string]interface{})
		assert.NotNil(t, book)
		assert.Equal(t, "Buku B", book["judul"])
		assert.Equal(t, "Penulis B", book["penulis"])
		assert.Equal(t, "Penerbit B", book["penerbit"])
	}
}

func TestBookController_Delete(t *testing.T) {
	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.BooksModel{}
	mdl.Init(gorm)
	var ctl = BookController{}
	ctl.InitBookController(mdl)

	req := httptest.NewRequest(http.MethodDelete, "/books/20", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("20")

	err := ctl.Delete()(c)
	if assert.NoError(t, err) {
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		books := response["book"].(map[string]interface{})

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Book deleted successfully", response["message"])
		assert.NotNil(t, books)
		assert.Equal(t, "Buku B", books["judul"])
		assert.Equal(t, "Penulis B", books["penulis"])
	}
}
