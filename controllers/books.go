package controller

import (
	model "OrmGo/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	model model.BooksModel
}

func (uc *BookController) InitBookController(um model.BooksModel) {
	uc.model = um
}

func (uc *BookController) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		books, err := uc.model.GetDatas()

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all books",
			"books":   books,
		})
	}
}

func (uc *BookController) Show() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		cnv, err := strconv.Atoi(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		book, err := uc.model.GetDataById(cnv)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get book by id",
			"book":    book,
		})
	}
}

func (uc *BookController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		book := model.Books{}
		c.Bind(&book)

		err := uc.model.Create(book)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create new book",
			"book":    book,
		})
	}
}

func (uc *BookController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		cnv, _ := strconv.Atoi(id)

		existingBook, err := uc.model.GetDataById(cnv)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": "Book not found",
			})
		}

		if err := c.Bind(existingBook); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())

		}

		updatedBook, err := uc.model.Update(cnv, existingBook)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Book updated successfully",
			"book":    updatedBook,
		})
	}
}

func (uc *BookController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		cnv, _ := strconv.Atoi(id)

		existingBook, err := uc.model.GetDataById(cnv)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"error": "Book not found",
			})
		}

		err = uc.model.Delete(cnv)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Book deleted successfully",
			"book":    existingBook,
		})
	}
}
