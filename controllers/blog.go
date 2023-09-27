package controller

import (
	model "OrmGo/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BlogController struct {
	model model.BlogModel
}

func (uc *BlogController) InitBlogController(um model.BlogModel) {
	uc.model = um
}

func (uc *BlogController) Index() echo.HandlerFunc {

	return func(c echo.Context) error {
		Blog, err := uc.model.GetDatas()

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all Blog",
			"Blog":    Blog,
		})
	}
}

func (uc *BlogController) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		cnv, err := strconv.Atoi(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		blog, err := uc.model.GetDataById(cnv)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get blog by id",
			"blog":    blog,
		})
	}
}

func (uc *BlogController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		blog := model.Blog{}
		c.Bind(&blog)

		err := uc.model.Create(blog)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create new blog",
			"blog":    blog,
		})
	}
}

func (uc *BlogController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		cnv, _ := strconv.Atoi(id)

		existingblog, err := uc.model.GetDataById(cnv)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": "blog not found",
			})
		}

		if err := c.Bind(existingblog); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())

		}

		updatedblog, err := uc.model.Update(cnv, existingblog)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "blog updated successfully",
			"blog":    updatedblog,
		})
	}
}

func (uc *BlogController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		cnv, _ := strconv.Atoi(id)

		existingblog, err := uc.model.GetDataById(cnv)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": "blog not found",
			})
		}

		err = uc.model.Delete(cnv)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "blog deleted successfully",
			"blog":    existingblog,
		})
	}
}
