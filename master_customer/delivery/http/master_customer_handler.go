package http

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/master_customer"

	"github.com/labstack/echo"
	"github.com/models"
	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}

type MasterVendorHandler struct {
	MasterVendorUsecase master_customer.Usecase
}

func NewMasterVendorHandler(e *echo.Echo, us master_customer.Usecase) {
	handler := &MasterVendorHandler{
		MasterVendorUsecase: us,
	}
	e.POST("master/import/master-customer", handler.ImportMasterVendor)
	e.GET("master/master-customer", handler.GetAllMasterVendor)
}

// GetByID will get article by given id
func (a *MasterVendorHandler) GetAllMasterVendor(c echo.Context) error {
	qpage := c.QueryParam("page")
	qsize := c.QueryParam("size")

	var limit = 20
	var page = 1
	var offset = 0

	page, _ = strconv.Atoi(qpage)
	limit, _ = strconv.Atoi(qsize)
	offset = (page - 1) * limit
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	art, err := a.MasterVendorUsecase.GetAll(ctx, page, limit, offset)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, art)

}

// Store will store the user by given request body
func (a *MasterVendorHandler) ImportMasterVendor(c echo.Context) error {

	filupload, image, _ := c.Request().FormFile("file-excel")
	dir, err := os.Getwd()
	if err != nil {
		return models.ErrInternalServerError
	}
	var imagePath string
	if filupload != nil {
		fileLocation := filepath.Join(dir, "files", image.Filename)
		targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			os.MkdirAll(filepath.Join(dir, "files"), os.ModePerm)
			return models.ErrInternalServerError
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, filupload); err != nil {
			return models.ErrInternalServerError
		}

		imagePath = fileLocation
		targetFile.Close()

	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	error := a.MasterVendorUsecase.Import(ctx, imagePath)

	if error != nil {
		return c.JSON(getStatusCode(error), ResponseError{Message: error.Error()})
	}
	return c.JSON(http.StatusOK, true)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrUnAuthorize:
		return http.StatusUnauthorized
	case models.ErrConflict:
		return http.StatusBadRequest
	case models.ErrBadParamInput:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
