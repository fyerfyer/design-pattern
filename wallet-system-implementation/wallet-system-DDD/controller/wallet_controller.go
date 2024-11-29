package controller

import (
	"net/http"
	"strconv"
	"wallet_system/service"

	"github.com/gin-gonic/gin"
)

type WalletController struct {
	Service *service.WalletService
}

func (c *WalletController) GetBalance(ctx *gin.Context) {
	walletID, err := strconv.ParseInt(ctx.Param("wallet_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet_id"})
		return
	}

	balance, err := c.Service.GetBalance(walletID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"balance": balance})
}

// others are the same
