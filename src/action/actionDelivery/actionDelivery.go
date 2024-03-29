package actionDelivery

import (
	"avengers-clinic/model/dto/actionDto"
	"avengers-clinic/model/dto/json"
	"avengers-clinic/pkg/constants"
	"avengers-clinic/pkg/middleware"
	"avengers-clinic/pkg/utils"
	"avengers-clinic/src/action"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type actionDelivery struct {
	actionUC action.ActionUsecase
}

func NewActionDelivery(v1Group *gin.RouterGroup, actionUC action.ActionUsecase) {
	handler := actionDelivery{actionUC: actionUC}
	actionGroup := v1Group.Group("/actions")
	{
		actionGroup.GET("", middleware.JwtAuth("ADMIN"), handler.GetAll)
		actionGroup.GET("/:id", middleware.JwtAuth("ADMIN"), handler.GetByID)
		actionGroup.POST("", middleware.JwtAuth("ADMIN"), handler.Create)
		actionGroup.PUT("/:id", middleware.JwtAuth("ADMIN"), handler.Update)
		actionGroup.DELETE("/:id", middleware.JwtAuth("ADMIN"), handler.Delete)
		actionGroup.DELETE("/:id/trash", middleware.JwtAuth("ADMIN"), handler.SoftDelete)
		actionGroup.PUT("/:id/restore", middleware.JwtAuth("ADMIN"), handler.Restore)
	}
}

func (delivery *actionDelivery) GetAll(c *gin.Context) {
	response, err := delivery.actionUC.GetAll()
	if err != nil {
		json.NewResponseError(c, err.Error(), constants.ActionService, "01")
		return
	}

	if len(response) == 0 {
		json.NewResponseForbidden(c, "Actions not found", constants.ActionService, "02")
		return
	}

	json.NewResponseSuccess(c, response, "actions successfully retrieved", constants.ActionService, "01")
}

func (delivery *actionDelivery) GetByID(c *gin.Context) {
	actionID := c.Param("id")
	response, err := delivery.actionUC.GetByID(actionID)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewResponseForbidden(c, "Action not found", constants.ActionService, "01")
			return
		}

		json.NewResponseError(c, err.Error(), constants.ActionService, "02")
		return
	}

	json.NewResponseSuccess(c, response, "action successfully retrieved", constants.ActionService, "01")
}

func (delivery *actionDelivery) Create(c *gin.Context) {
	var request actionDto.CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		json.NewResponseError(c, err.Error(), constants.ActionService, "01")
		return
	}

	if err := utils.Validated(request); err != nil {
		json.NewResponseBadRequest(c, err, "Bad request", constants.ActionService, "02")
		return
	}

	response, err := delivery.actionUC.Create(request)
	if err != nil {
		if err.Error() == "1" {
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName:"name", Message:"Name is already registered"}}, "Bad request", constants.ActionService, "03")
			return
		}

		json.NewResponseError(c, err.Error(), constants.ActionService, "04")
		return
	}

	json.NewResponseCreated(c, response, "Action created successfully", constants.ActionService, "01")
}

func (delivery *actionDelivery) Update(c *gin.Context) {
	var request actionDto.UpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		json.NewResponseError(c, err.Error(), constants.ActionService, "01")
		return
	}
	request.ID = c.Param("id")

	response, err := delivery.actionUC.Update(request)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewResponseForbidden(c, "Action not found", constants.ActionService, "02")
			return
		}

		if err.Error() == "1" {
			json.NewResponseBadRequest(c, []json.ValidationField{{FieldName:"name", Message:"Name is already registered"}}, "Bad request", constants.ActionService, "03")
			return
		}

		json.NewResponseError(c, err.Error(), constants.ActionService, "04")
		return
	}

	json.NewResponseSuccess(c, response, "Action updated successfully", constants.ActionService, "01")
}

func (delivery *actionDelivery) Delete(c *gin.Context) {
	actionID := c.Param("id")
	err := delivery.actionUC.Delete(actionID)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewResponseForbidden(c, "Action not found", constants.ActionService, "01")
			return
		}

		json.NewResponseError(c, err.Error(), constants.ActionService, "02")
		return
	}
	json.NewResponseSuccess(c, nil, "Action deleted successfully", constants.ActionService, "01")
}

func (delivery *actionDelivery) SoftDelete(c *gin.Context) {
	actionID := c.Param("id")
	err := delivery.actionUC.SoftDelete(actionID)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewResponseForbidden(c, "Action not found", constants.ActionService, "01")
			return
		}

		json.NewResponseError(c, err.Error(), constants.ActionService, "02")
		return
	}
	json.NewResponseSuccess(c, nil, "Action deleted successfully", constants.ActionService, "01")
}

func (delivery *actionDelivery) Restore(c *gin.Context) {
	actionID := c.Param("id")
	err := delivery.actionUC.Restore(actionID)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewResponseForbidden(c, "Action not found", constants.ActionService, "01")
			return
		}

		json.NewResponseError(c, err.Error(), constants.ActionService, "02")
		return
	}
	json.NewResponseSuccess(c, nil, "Action restored successfully", constants.ActionService, "01")
}