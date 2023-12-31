package httplib

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DefaultResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

type DefaultPaginationResponse struct {
	Status     string      `json:"status"`
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Page       int         `json:"page"`
	Size       int         `json:"size"`
	TotalCount uint64      `json:"totalCount"`
	TotalPages uint64      `json:"totalPages"`
	Data       interface{} `json:"data"`
}

func SetSuccessResponse(c echo.Context, code int, message string, data interface{}) error {
	return c.JSON(code, DefaultResponse{
		Status:  http.StatusText(code),
		Code:    code,
		Data:    data,
		Message: message,
		Errors:  nil,
	})
}

func SetPaginationResponse(c echo.Context, code int, message string, data interface{}, totalCount uint64, pg *Query) error {
	return c.JSON(code, DefaultPaginationResponse{
		Status:     http.StatusText(code),
		Code:       code,
		Message:    message,
		Page:       pg.GetPage(),
		Size:       pg.GetSize(),
		TotalCount: totalCount,
		TotalPages: uint64(GetTotalPages(int(totalCount), pg.GetSize())),
		Data:       data,
	})
}

func SetErrorResponse(c echo.Context, code int, message string) error {
	return c.JSON(code, DefaultResponse{
		Status:  http.StatusText(code),
		Code:    code,
		Data:    nil,
		Message: message,
	})
}

func SetCustomResponse(c echo.Context, code int, message string, data interface{}, dataErr interface{}) error {
	return c.JSON(code, DefaultResponse{
		Status:  http.StatusText(code),
		Code:    code,
		Data:    data,
		Message: message,
		Errors:  dataErr,
	})
}

func SetBinarySuccessResponse(c echo.Context, code int, filename string, data []byte) error {
	c.Response().Header().Set("Content-Type", "text/csv")
	c.Response().Header().Set("Content-Disposition", "attachment; filename="+filename)
	c.Response().WriteHeader(code)
	_, err := c.Response().Write(data)
	if err != nil {
		return err
	}
	return c.JSON(code, c.Response())
}
