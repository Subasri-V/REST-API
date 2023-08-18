package controllers

import (
	"net/http"
	"rest_api_app/interfaces"
	"rest_api_app/models"

	"github.com/gin-gonic/gin"
)

type TransactController struct {
	TransactService interfaces.ITransact
}

func InitTransactController(transactService interfaces.ITransact) TransactController {
	return TransactController{transactService}
}

func (tc *TransactController) CreateTransaction(ctx *gin.Context) {
	var transact *models.Transact
	if err := ctx.ShouldBindJSON(&transact); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newTransact, err := tc.TransactService.CreateTransaction(transact)

	if err != nil {

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newTransact})
}

func (tc *TransactController) GetAllTransaction(ctx *gin.Context) {

}
