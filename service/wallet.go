package service

import (
	"github.com/vnFuhung2903/postgresql/api"
	"github.com/vnFuhung2903/postgresql/model"
	"github.com/vnFuhung2903/postgresql/repository"
)

type WalletService interface {
	CreateNewWallet(req *api.WalletReqBody) (*model.Wallet, error)
	GetWallets(req *api.WalletReqBody) ([]*model.Wallet, error)
}

type walletService struct {
	Wr repository.WalletRepository
}

func NewWalletService(wr *repository.WalletRepository) WalletService {
	return &walletService{Wr: *wr}
}

func (walletService *walletService) CreateNewWallet(req *api.WalletReqBody) (*model.Wallet, error) {
	wallet, err := walletService.Wr.CreateWallet(req.UserId)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (walletService *walletService) GetWallets(req *api.WalletReqBody) ([]*model.Wallet, error) {
	wallets, err := walletService.Wr.FindByUserId(req.UserId)
	if err != nil {
		return nil, err
	}

	return wallets, nil
}
