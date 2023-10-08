package controller

import (
	"OrmGo/config"
	model "OrmGo/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var createdUserID uint = 20

func TestUsersController_Index(t *testing.T) {

	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.UsersModel{}
	mdl.Init(gorm)
	var ctl = UserController{}
	ctl.InitUserController(mdl, *config)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	err := ctl.Index()(c)

	if assert.NoError(t, err) {
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		users := response["users"].([]interface{})

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success get all users", response["message"])
		assert.NotNil(t, users)
	}
}

func TestUsersController_Show(t *testing.T) {

	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.UsersModel{}
	mdl.Init(gorm)
	var ctl = UserController{}
	ctl.InitUserController(mdl, *config)

	req := httptest.NewRequest(http.MethodGet, "/users/15", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("15")
	err := ctl.Show()(c)

	if assert.NoError(t, err) {
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)
		user := response["user"].(map[string]interface{})
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success get user by id", response["message"])
		assert.NotNil(t, user)
		assert.Equal(t, "budiawanilham04@gmail.com", user["email"])
		assert.Equal(t, "12345", user["password"])
	}
}

func TestUsersController_Create(t *testing.T) {

	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.UsersModel{}
	mdl.Init(gorm)
	var ctl = UserController{}
	ctl.InitUserController(mdl, *config)

	reqBody := `{
		"ID": 20,
		"name": "Coba Unit Testing",
		"email": "coba@gmail.com",
		"password": "12345"
	}`

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	err := ctl.Create()(c)

	if assert.NoError(t, err) {
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success create new user", response["message"])

		user := response["user"].(map[string]interface{})
		assert.NotNil(t, user)
		assert.Equal(t, "coba@gmail.com", user["email"])
		assert.Equal(t, "12345", user["password"])

	}
}

func TestUsersController_Update(t *testing.T) {
	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.UsersModel{}
	mdl.Init(gorm)
	var ctl = UserController{}
	ctl.InitUserController(mdl, *config)

	reqBody := `{
		"name": "Coba Unit Testing",
		"email": "coba@gmail.com",
		"password": "12345"
	}`

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/users/%d", createdUserID), strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", createdUserID))

	err := ctl.Update()(c)

	if assert.NoError(t, err) {
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "User updated successfully", response["message"])

		user := response["user"].(map[string]interface{})
		assert.NotNil(t, user)
		assert.Equal(t, "coba@gmail.com", user["email"])
		assert.Equal(t, "12345", user["password"])
	}
}

func TestUsersController_Delete(t *testing.T) {
	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.UsersModel{}
	mdl.Init(gorm)
	var ctl = UserController{}
	ctl.InitUserController(mdl, *config)

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/users/%d", createdUserID), nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", createdUserID))

	err := ctl.Delete()(c)

	if assert.NoError(t, err) {
		var response map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "User deleted successfully", response["message"])

		user := response["user"].(map[string]interface{})
		assert.NotNil(t, user)
		assert.Equal(t, "coba@gmail.com", user["email"])
		assert.Equal(t, "12345", user["password"])
	}
}

func TestUsersController_Login(t *testing.T) {
	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.UsersModel{}
	mdl.Init(gorm)
	var ctl = UserController{}
	ctl.InitUserController(mdl, *config)

	// Jika Behasil Login
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{
		"email": "budiawanilham04@gmail.com",
		"password": "12345"
	}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	err := ctl.Login()(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	token, tokenExists := response["token"].(string)
	refreshToken, refreshTokenExists := response["refresh_token"].(string)

	assert.True(t, tokenExists)
	assert.True(t, refreshTokenExists)
	assert.NotEmpty(t, token)
	assert.NotEmpty(t, refreshToken)
}

func TestUsersController_LoginIncorrectCredentials(t *testing.T) {
	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.UsersModel{}
	mdl.Init(gorm)
	var ctl = UserController{}
	ctl.InitUserController(mdl, *config)

	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{
		"email": "budiawanilham04@gmail.com",
		"password": "wrongpassword"
	}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	err := ctl.Login()(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	errorMessage, errorExists := response["error"].(string)

	assert.True(t, errorExists)
	assert.Equal(t, "Invalid credentials", errorMessage)
}

func TestUsersController_LoginInvalidRequest(t *testing.T) {

	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.UsersModel{}
	mdl.Init(gorm)
	var ctl = UserController{}
	ctl.InitUserController(mdl, *config)

	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{
		"username": "budiawanilham04@gmail.com",
		"fieldkosong": "test",
	}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	err := ctl.Login()(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	errorMessage, errorExists := response["error"].(string)

	assert.True(t, errorExists)
	assert.Equal(t, "Invalid request", errorMessage)
}
