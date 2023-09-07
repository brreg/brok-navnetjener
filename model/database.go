package model

import (
	"brok/navnetjener/database"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	OwnerPersonFirstName string `gorm:"size:255;" json:"owner_person_first_name"`
	OwnerPersonLastName  string `gorm:"size:255;" json:"owner_person_last_name"`
	OwnerPersonPnr       string `gorm:"" json:"owner_person_pnr"`
	OwnerPersonBirthDate string `gorm:"" json:"owner_person_birth_date"`

	OwnerCompanyName  string `gorm:"size:255;" json:"owner_company_name"`
	OwnerCompanyOrgnr string `gorm:"size:9;" json:"owner_company_orgnr"`

	CapTableOrgnr string `gorm:"size:9;" json:"cap_table_orgnr" binding:"required"`
	WalletAddress string `gorm:"size:42;not null;unique" json:"wallet_address" binding:"required"`
}

type Owner struct {
	Person  Person  `json:"person"`
	Company Company `json:"company"`
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"`
}

type Company struct {
	Name  string `json:"name"`
	Orgnr string `json:"orgnr"`
}

type PublicWalletInfo struct {
	CapTableOrgnr string `json:"orgnr"`
	WalletAddress string `json:"wallet_address"`
	Owner         Owner  `json:"owner"`
}

func (wallet *Wallet) Save() (*Wallet, error) {
	SanitizeWallet(wallet)
	wallet.WalletAddress = strings.ToLower(wallet.WalletAddress)

	if err := database.Database.Create(&wallet).Error; err != nil {
		logrus.Error("error creating wallet in the database: ", err)
		return &Wallet{}, err
	}

	return wallet, nil
}

func FindWalletByOrgnr(orgnr string) ([]PublicWalletInfo, error) {
	var wallets []Wallet
	var publicWallets []PublicWalletInfo
	safeOrgnr := SanitizeString(orgnr)
	err := database.Database.Where("cap_table_orgnr=?", safeOrgnr).Find(&wallets).Error
	if err != nil {
		logrus.Error("could not find wallet in db with orgnr: ", orgnr)
		return []PublicWalletInfo{}, err
	}

	for _, wallet := range wallets {
		publicWallets = append(publicWallets, parseWalletToPublicInfo(wallet))
	}

	logrus.Debug("found wallets in db with orgnr: ", orgnr, " wallets: ", publicWallets)
	return publicWallets, nil
}

func FindWalletByPnr(pnr string) ([]PublicWalletInfo, error) {
	var wallets []Wallet
	var publicWallets []PublicWalletInfo
	safePnr := SanitizeString(pnr)
	err := database.Database.Where("owner_person_pnr=?", safePnr).Find(&wallets).Error
	if err != nil {
		logrus.Error(err)
		return []PublicWalletInfo{}, err
	}

	if len(wallets) == 0 {
		logrus.Warn("could not find person in db with pnr: ", pnr[0:6], "*****")
		// return empty with a new error if no person is found
		return []PublicWalletInfo{}, gorm.ErrRecordNotFound
	}

	for _, wallet := range wallets {
		publicWallets = append(publicWallets, parseWalletToPublicInfo(wallet))
	}

	logrus.Debug("found wallets in db with pnr: ", pnr[0:6], "*****", " wallets: ", publicWallets)
	return publicWallets, nil
}

func FindWalletByWalletAddress(walletAddress string) (PublicWalletInfo, error) {
	var wallet Wallet
	safeWalletAddress := SanitizeString(walletAddress)
	safeLowerCaseWalletAddress := strings.ToLower(safeWalletAddress)
	err := database.Database.Where("wallet_address=?", safeLowerCaseWalletAddress).Find(&wallet).Error

	if err != nil {
		logrus.Error("could not find person in db with wallet_address: ", walletAddress)
		return PublicWalletInfo{}, err
	}
	publicWallet := parseWalletToPublicInfo(wallet)

	logrus.Debug("found wallets in db with walletAddress: ", walletAddress, " wallets: ", publicWallet)
	return publicWallet, nil
}

func FindAllWallets() ([]PublicWalletInfo, error) {
	var wallets []Wallet
	var publicWallets []PublicWalletInfo
	err := database.Database.Find(&wallets).Error

	if err != nil {
		logrus.Error("could not find any wallets in db")
		return []PublicWalletInfo{}, err
	}

	for _, wallet := range wallets {
		publicWallets = append(publicWallets, parseWalletToPublicInfo(wallet))
	}

	logrus.Debug("found wallets in db: ", publicWallets)
	return publicWallets, nil
}

func parseWalletToPublicInfo(wallet Wallet) PublicWalletInfo {
	// return empty if empty
	if wallet == (Wallet{}) {
		return PublicWalletInfo{}
	}
	return PublicWalletInfo{
		CapTableOrgnr: wallet.CapTableOrgnr,
		WalletAddress: wallet.WalletAddress,
		Owner: Owner{
			Person: Person{
				FirstName: wallet.OwnerPersonFirstName,
				LastName:  wallet.OwnerPersonLastName,
				BirthDate: wallet.OwnerPersonBirthDate,
			},
			Company: Company{
				Name:  wallet.OwnerCompanyName,
				Orgnr: wallet.OwnerCompanyOrgnr,
			},
		},
	}
}
