package model

import (
	"brok/navnetjener/database"

	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	FirstName     string `gorm:"size:255;not null" json:"first_name" binding:"required"`
	LastName      string `gorm:"size:255;not null" json:"last_name" binding:"required"`
	Orgnr         int    `gorm:"not null" json:"orgnr" binding:"required"`
	Pnr           string `gorm:"not null" json:"pnr" binding:"required"`
	BirthDate     string `gorm:"not null" json:"birth_date"`
	WalletAddress string `gorm:"size:42;not null;unique" json:"wallet_address" binding:"required"`
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"`
}

type PublicWalletInfo struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Orgnr         int    `json:"orgnr"`
	BirthDate     string `json:"birth_date"`
	WalletAddress string `json:"wallet_address"`
}

func (wallet *Wallet) Save() (*Wallet, error) {
	if err := SanitizeWallet(wallet); err != nil {
		return &Wallet{}, err
	}

	if err := database.Database.Create(&wallet).Error; err != nil {
		return &Wallet{}, err
	}

	return wallet, nil
}

func FindWalletByOrgnr(orgnr string) ([]PublicWalletInfo, error) {
	var wallets []Wallet
	var publicWallets []PublicWalletInfo
	safeOrgnr := SanitizeString(orgnr)
	err := database.Database.Where("orgnr=?", safeOrgnr).Find(&wallets).Error
	if err != nil {
		return []PublicWalletInfo{}, err
	}

	for _, wallet := range wallets {
		publicWallets = append(publicWallets, parseWalletToPublicInfo(wallet))
	}

	return publicWallets, nil
}

func FindWalletByPnr(pnr string) ([]PublicWalletInfo, error) {
	var wallets []Wallet
	var publicWallets []PublicWalletInfo
	safePnr := SanitizeString(pnr)
	err := database.Database.Where("pnr=?", safePnr).Find(&wallets).Error
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
	safeWalletAddress := SanitizeString(walletAddress)
	err := database.Database.Where("wallet_address=?", safeWalletAddress).Find(&wallet).Error
	if err != nil {
		return PublicWalletInfo{}, err
	}

	return parseWalletToPublicInfo(wallet), nil
}

func findPersonByWalletAddress(walletAddress string) (Person, error) {
	var wallet Wallet
	var person Person
	safeWalletAddress := SanitizeString(walletAddress)
	err := database.Database.Where("wallet_address=?", safeWalletAddress).Find(&wallet).Error
	if err != nil {
		return Person{}, err
	}

	person = Person{
		FirstName: wallet.FirstName,
		LastName:  wallet.LastName,
		BirthDate: wallet.BirthDate,
	}

	return person, nil
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
		BirthDate:     wallet.BirthDate,
		WalletAddress: wallet.WalletAddress,
	}
}
