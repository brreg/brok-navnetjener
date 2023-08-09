package model

import (
	"brok/navnetjener/database"

	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	FirstName     string `gorm:"size:255;not null" json:"first_name"`
	LastName      string `gorm:"size:255;not null" json:"last_name"`
	Orgnr         string `gorm:"size:9;not null" json:"orgnr"`
	Pnr           string `gorm:"size:11;not null" json:"pnr"`
	YearBorn      string `gorm:"size:2;not null" json:"year_born"`
	WalletAddress string `gorm:"size:42;not null;unique" json:"wallet_address"`
}

type PublicWalletInfo struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Orgnr         string `json:"orgnr"`
	YearBorn      string `json:"year_born"`
	WalletAddress string `json:"wallet_address"`
}

func (wallet *Wallet) Save() (*Wallet, error) {
	err := database.Database.Create(&wallet).Error
	if err != nil {
		return &Wallet{}, err
	}
	return wallet, nil
}

func FindWalletByPnr(pnr string) ([]PublicWalletInfo, error) {
	var wallets []Wallet
	var publicWallets []PublicWalletInfo
	err := database.Database.Where("pnr=?", pnr).Find(&wallets).Error
	if err != nil {
		return []PublicWalletInfo{}, err
	}

	for _, wallet := range wallets {
		publicWallets = append(publicWallets, parseWalletToPublicInfo(wallet))
	}

	return publicWallets, nil
}

func FindWalletByWalletAddress(walletAddress string) (PublicWalletInfo, error) {
	var wallet Wallet
	err := database.Database.Where("wallet_address=?", walletAddress).Find(&wallet).Error
	if err != nil {
		return PublicWalletInfo{}, err
	}

	return parseWalletToPublicInfo(wallet), nil
}

func FindAllWallets() ([]PublicWalletInfo, error) {
	var wallets []Wallet
	var publicWallets []PublicWalletInfo
	err := database.Database.Find(&wallets).Error

	if err != nil {
		return []PublicWalletInfo{}, err
	}

	for _, wallet := range wallets {
		publicWallets = append(publicWallets, parseWalletToPublicInfo(wallet))
	}

	return publicWallets, nil
}

func parseWalletToPublicInfo(wallet Wallet) PublicWalletInfo {
	// return empty if empty
	if wallet == (Wallet{}) {
		return PublicWalletInfo{}
	}
	return PublicWalletInfo{
		FirstName:     wallet.FirstName,
		LastName:      wallet.LastName,
		Orgnr:         wallet.Orgnr,
		YearBorn:      wallet.YearBorn,
		WalletAddress: wallet.WalletAddress,
	}
}
