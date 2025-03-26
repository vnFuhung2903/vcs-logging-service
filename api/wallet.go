package api

import "github.com/vnFuhung2903/postgresql/model"

type WalletReqBody struct {
	UserId uint
}

type WalletResBody struct {
	Id           uint `json:"id"`
	UserId       uint `json:"userId"`
	WalletNubmer uint `json:"walletNumber"`
	Balance      uint `json:"balance"`
}

func FormatWalletResBody(wallet *model.Wallet) WalletResBody {
	return WalletResBody{
		Id:      wallet.Id,
		UserId:  wallet.UserId,
		Balance: wallet.Balance,
	}
}
