package repository

import (
	"github.com/vnFuhung2903/vcs-logging-service/model"
	"gorm.io/gorm"
)

type WalletRepository interface {
	FindAll() ([]*model.Wallet, error)
	FindById(id uint) ([]*model.Wallet, error)
	FindByUserId(userId uint) ([]*model.Wallet, error)
	CreateWallet(userId uint, walletNumber uint) (*model.Wallet, error)
	UpdateBalance(wallet *model.Wallet, balance uint) error
}

type walletRepository struct {
	Db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepository{Db: db}
}

func (wr *walletRepository) FindAll() ([]*model.Wallet, error) {
	var wallets []*model.Wallet
	res := wr.Db.Find(wallets)
	if res != nil {
		return nil, res.Error
	}
	return wallets, nil
}

func (wr *walletRepository) FindById(id uint) ([]*model.Wallet, error) {
	var wallets []*model.Wallet
	res := wr.Db.Find(wallets, model.Wallet{Id: id})
	if res != nil {
		return nil, res.Error
	}
	return wallets, nil
}

func (wr *walletRepository) FindByUserId(userId uint) ([]*model.Wallet, error) {
	var wallets []*model.Wallet
	res := wr.Db.Find(wallets, model.Wallet{UserId: userId})
	if res != nil {
		return nil, res.Error
	}
	return wallets, nil
}

func (wr *walletRepository) CreateWallet(userId uint, walletNumber uint) (*model.Wallet, error) {
	res := wr.Db.Create(model.Wallet{
		UserId:       userId,
		WalletNumber: walletNumber,
	})
	if res.Error != nil {
		return nil, res.Error
	}

	var wallet *model.Wallet
	res = wr.Db.Find(wallet, model.Wallet{UserId: userId})
	if res.Error != nil {
		return nil, res.Error
	}

	return wallet, nil
}

func (wr *walletRepository) UpdateBalance(wallet *model.Wallet, balance uint) error {
	res := wr.Db.Save(model.Wallet{
		Id:      wallet.Id,
		UserId:  wallet.UserId,
		Balance: balance,
	})
	return res.Error
}
