package controller

import (
	"OrmGo/config"
	"OrmGo/helpers"
	model "OrmGo/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	config config.ProgramConfig
	model  model.UsersModel
}

func (uc *UserController) InitUserController(um model.UsersModel, c config.ProgramConfig) {
	uc.model = um
	uc.config = c
}

func (uc *UserController) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uc.model.GetDatas()

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all users",
			"users":   users,
		})
	}
}

func (uc *UserController) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		cnv, err := strconv.Atoi(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		user, err := uc.model.GetDataById(cnv)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get user by id",
			"user":    user,
		})
	}
}

func (uc *UserController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := model.Users{}
		c.Bind(&user)

		err := uc.model.Create(user)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create new user",
			"user":    user,
		})
	}
}

func (uc *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		cnv, _ := strconv.Atoi(id)

		existingUser, err := uc.model.GetDataById(cnv)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": "User not found",
			})
		}

		if err := c.Bind(existingUser); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())

		}

		updatedUser, err := uc.model.Update(cnv, existingUser)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "User updated successfully",
			"user":    updatedUser,
		})
	}
}

func (uc *UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		cnv, _ := strconv.Atoi(id)

		existingUser, err := uc.model.GetDataById(cnv)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": "User not found",
			})
		}

		// Panggil fungsi Delete dari model
		err = uc.model.Delete(cnv)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "User deleted successfully",
			"user":    existingUser,
		})
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": "Invalid request",
			})
		}

		user, err := uc.model.GetUserByEmail(req.Email)
		if err != nil || user.Password != req.Password {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error": "Invalid credentials",
			})
		}

		token, err := helpers.GenerateToken(uc.config.SECRET, user.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "Failed to generate token",
			})
		}

		refreshToken, err := helpers.GenerateRefreshToken(uc.config.SECRETREFRESH, user.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "Failed to generate refresh token",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"token":         token,
			"refresh_token": refreshToken,
		})
	}
}
