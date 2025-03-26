package controller

import "github.com/vnFuhung2903/postgresql/service"

type Controller struct {
	authService   service.AuthService
	userService   service.UserService
	walletService service.WalletService
}

func NewControlle(
	as service.AuthService,
	us service.UserService,
	ws service.WalletService,
) *Controller {
	return &Controller{
		authService:   as,
		userService:   us,
		walletService: ws,
	}
}
