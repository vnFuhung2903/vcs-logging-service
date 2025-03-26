package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vnFuhung2903/postgresql/api"
	"github.com/vnFuhung2903/postgresql/model"
	"github.com/vnFuhung2903/postgresql/util"
)

func (c *Controller) Login(ctx *gin.Context) {
	inp := &api.LoginReqBody{}
	err := ctx.ShouldBindJSON(inp)
	if err != nil {
		res := api.NewErrorResponse("login failed", http.StatusUnprocessableEntity, err)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	user, err := c.authService.Login(inp)
	if err != nil {
		res := api.NewErrorResponse("login failed", http.StatusUnprocessableEntity, err)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	walletInp := &api.WalletReqBody{}
	walletInp.UserId = user.Id
	wallets, err := c.walletService.GetWallets(walletInp)
	if err != nil {
		res := api.NewErrorResponse("login failed", http.StatusUnprocessableEntity, err)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	token, err := util.GenerateJwtToken(user.Id)
	if err != nil {
		res := api.NewErrorResponse("login failed", http.StatusUnprocessableEntity, err)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	res := api.NewSuccessResponse(
		"login success",
		http.StatusOK,
		api.FormatLoginResBody(user, wallets, token),
	)
	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) Register(ctx *gin.Context) {
	inp := &api.RegisterReqBody{}
	err := ctx.ShouldBindJSON(inp)
	if err != nil {
		res := api.NewErrorResponse("register failed", http.StatusUnprocessableEntity, err)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	user, err := c.userService.Register(inp)
	if err != nil {
		res := api.NewErrorResponse("register failed", http.StatusUnprocessableEntity, err)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	walletInp := &api.WalletReqBody{}
	walletInp.UserId = user.Id
	wallet, err := c.walletService.CreateNewWallet(walletInp)
	if err != nil {
		res := api.NewErrorResponse("register failed", http.StatusUnprocessableEntity, err)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	token, err := util.GenerateJwtToken(user.Id)
	if err != nil {
		res := api.NewErrorResponse("register failed", http.StatusUnprocessableEntity, err)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	res := api.NewSuccessResponse(
		"register success",
		http.StatusOK,
		api.FormatLoginResBody(user, []*model.Wallet{wallet}, token),
	)
	ctx.JSON(http.StatusOK, res)
}
