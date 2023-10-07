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

type UserResponse struct {
	Message string
	Data    []model.Users
}

func TestIndex(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get user normal",
			path:       "/users",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	var config = config.InitConfig()
	var gorm = model.InitModel(*config)
	var model = model.UsersModel{}
	model.Init(gorm)
	var ctl = UserController{}
	ctl.InitUserController(model, *config)

	var e = echo.New()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if assert.NoError(t, ctl.Index()(c)) {
				assert.Equal(t, testCase.expectCode, rec.Code)

				var user UserResponse
				if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &user)) {
					assert.Equal(t, testCase.sizeData, len(user.Data))
				}
			}
		})
	}
}
