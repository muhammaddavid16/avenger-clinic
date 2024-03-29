package json

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type (
	jsonResponse struct {
		Code    string      `json:"responseCode"`
		Message string      `json:"responseMessage"`
		Data    interface{} `json:"data,omitempty"`
		Paging interface{} 	`json:"paging,omitempty"`
	}

	jsonErrorResponse struct {
		Code    string `json:"responseCode"`
		Message string `json:"responseMessage"`
		Error   string `json:"error,omitempty"`
	}

	ValidationField struct {
		FieldName string `json:"field"`
		Message   string `json:"message"`
	}

	jsonBadRequestResponse struct {
		Code             string            `json:"responseCode"`
		Message          string            `json:"responseMessage"`
		ErrorDescription []ValidationField `json:"error_description,omitempty"`
	}
)

func NewResponseSuccess(c *gin.Context, result interface{}, message, serviceCode, responseCode string) {
	c.JSON(http.StatusOK, jsonResponse{
		Code:    "200" + serviceCode + responseCode,
		Message: message,
		Data:    result,
	})
}

func NewResponseCreated(c *gin.Context, result interface{}, message, serviceCode, responseCode string) {
	c.JSON(http.StatusCreated, jsonResponse{
		Code:    "201" + serviceCode + responseCode,
		Message: message,
		Data:    result,
	})
}

func NewResponseError(c *gin.Context, err, serviceCode, errorCode string) {
	log.Error().Msg(err)
	c.JSON(http.StatusInternalServerError, jsonErrorResponse{
		Code:    "500" + serviceCode + errorCode,
		Message: "internal server error",
		Error:   err,
	})
}

func NewResponseBadRequest(c *gin.Context, validationField []ValidationField, message, serviceCode, errorCode string) {
	c.JSON(http.StatusBadRequest, jsonBadRequestResponse{
		Code:             "400" + serviceCode + errorCode,
		Message:          message,
		ErrorDescription: validationField,
	})
}

func NewResponseUnauthorized(c *gin.Context, message, serviceCode, errorCode string) {
	c.JSON(http.StatusUnauthorized, jsonResponse{
		Code:    "401" + serviceCode + errorCode,
		Message: message,
	})
}

func NewResponseForbidden(c *gin.Context, message, serviceCode, errorCode string) {
	c.JSON(http.StatusForbidden, jsonResponse{
		Code:    "403" + serviceCode + errorCode,
		Message: message,
	})
}
func NewResponseNotFound(c *gin.Context, message, serviceCode, errorCode string) {
	c.JSON(http.StatusNotFound, jsonResponse{
		Code:    "404" + serviceCode + errorCode,
		Message: message,
	})
}