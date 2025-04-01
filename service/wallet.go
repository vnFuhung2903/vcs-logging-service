package service

import (
	"github.com/vnFuhung2903/vcs-logging-service/model"
	"github.com/vnFuhung2903/vcs-logging-service/repository"
)

type WalletService interface {
	CreateNewWallet(userId uint, walletNumber uint) (*model.Wallet, error)
	GetWallets(userId uint) ([]*model.Wallet, error)
}

type walletService struct {
	Wr repository.WalletRepository
}

func NewWalletService(wr *repository.WalletRepository) WalletService {
	return &walletService{Wr: *wr}
}

func (walletService *walletService) CreateNewWallet(userId uint, wallletNumber uint) (*model.Wallet, error) {
	wallet, err := walletService.Wr.CreateWallet(userId, wallletNumber)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (walletService *walletService) GetWallets(userId uint) ([]*model.Wallet, error) {
	wallets, err := walletService.Wr.FindByUserId(userId)
	if err != nil {
		return nil, err
	}

	return wallets, nil
}
