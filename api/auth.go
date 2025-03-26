package api

import "github.com/vnFuhung2903/postgresql/model"

type LoginReqBody struct {
	Email    string
	Password string
}

type RegisterReqBody struct {
	Email    string
	Password string
	Name     string
}

type LoginResBody struct {
	Id      uint   `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Wallets []uint `json:"wallet"`
	Token   string `json:"token"`
}

func FormatLoginResBody(user *model.User, wallets []*model.Wallet, token string) LoginResBody {
	var walletNumbers []uint
	for _, wallet := range wallets {
		walletNumbers = append(walletNumbers, wallet.WalletNumber)
	}
	return LoginResBody{
		Id:      user.Id,
		Email:   user.Email,
		Name:    user.Name,
		Wallets: walletNumbers,
		Token:   token,
	}
}
