package controller

import (
	"OrmGo/config"
	model "OrmGo/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	mockUserModel := new(model.UsersModelMock)

	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var mdl = model.UsersModel{}
	mdl.Init(gorm)
	var ctl = UserController{}
	ctl.InitUserController(mdl, *config)

	mockUserModel.On("GetDatas").Return([]model.Users{{
		Name:     "Ilham Budiawan",
		Email:    "budiawanilham04@gmail.com",
		Password: "12345"}}, nil)

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
		assert.Len(t, users, 2)
	}
}
