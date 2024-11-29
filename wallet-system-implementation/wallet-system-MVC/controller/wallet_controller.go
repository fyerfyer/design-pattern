package controller

import (
	"net/http"
	"strconv"
	"wallet-system/service"

	"github.com/gin-gonic/gin"
)

type VirtualWalletController struct {
	walletService *service.VirtualWalletService
}

func NewVirtualWalletController(walletService *service.VirtualWalletService) *VirtualWalletController {
	return &VirtualWalletController{walletService: walletService}
}

func (c *VirtualWalletController) GetBalance(ctx gin.Context) {
	walletID, err := strconv.ParseInt(ctx.Param("wallet_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet_id"})
		return
	}

	balance, err := c.walletService.GetBalance(walletID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"balance": balance})
}

func (c *VirtualWalletController) debit(ctx *gin.Context) {
	walletID, err := strconv.ParseInt(ctx.Param("wallet_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet_id"})
		return
	}

	var requestBody struct {
		Amount float64 `json:"amount"`
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := c.walletService.Debit(walletID, requestBody.Amount); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Debit successful"})
}

func (c *VirtualWalletController) Transfer(ctx *gin.Context) {
	fromWalletID, err := strconv.ParseInt(ctx.Param("from_wallet_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid from_wallet_id"})
		return
	}
	toWalletID, err := strconv.ParseInt(ctx.Param("to_wallet_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid to_wallet_id"})
		return
	}

	var requestBody struct {
		Amount float64 `json:"amount"`
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := c.walletService.Transfer(fromWalletID, toWalletID, requestBody.Amount); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transfer successful"})
}
