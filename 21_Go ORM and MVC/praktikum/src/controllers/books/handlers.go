package books

import (
	"gorm-mvc-api/src/database"
	"gorm-mvc-api/src/database/factories"
	"gorm-mvc-api/src/models/book"
	"gorm-mvc-api/src/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

var db = database.Connect()
var factory = factories.BookFactory(db)

func GetAllBooks(echoContext echo.Context) error {
	books, err := factory.GetAllBooks()

	if err != nil {
		return utils.CreateResponse(echoContext, http.StatusBadRequest, "request failed", err)
	}

	return utils.CreateResponse(echoContext, http.StatusOK, "OK", books)
}

func GetBook(echoContext echo.Context) error {
	var bookID string = echoContext.Param("id")
	bookData, err := factory.GetBook(bookID)

	if err != nil {
		return utils.CreateResponse(echoContext, http.StatusBadRequest, "request failed", err)
	}

	return utils.CreateResponse(echoContext, http.StatusOK, "OK", bookData)
}

func CreateBook(echoContext echo.Context) error {
	var request book.Book
	echoContext.Bind(&request)

	if err := factory.CreateBook(request); err != nil {
		return utils.CreateResponse(echoContext, http.StatusBadRequest, "request failed")
	}

	return utils.CreateResponse(echoContext, http.StatusCreated, "book created")
}

func UpdateBook(echoContext echo.Context) error {
	var request book.Book
	echoContext.Bind(&request)

	if err := factory.UpdateBook(request, echoContext.Param("id")); err != nil {
		return utils.CreateResponse(echoContext, http.StatusBadRequest, "request failed", err)
	}

	return utils.CreateResponse(echoContext, http.StatusNoContent, "book updated")
}

func DeleteBook(echoContext echo.Context) error {
	var bookID string = echoContext.Param("id")

	if err := factory.DeleteBook(bookID); err != nil {
		return utils.CreateResponse(echoContext, http.StatusBadRequest, "request failed")
	}

	return utils.CreateResponse(echoContext, http.StatusNoContent, "book deleted")
}
