package main

import (
	"brok/navnetjener/api"
	"brok/navnetjener/database"
	"brok/navnetjener/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.Wallet{})
}

func serveApplication() {
	router := routerConfig()
	router.Run(":9000")
	fmt.Println("Server running at port 9000")
}

func routerConfig() *gin.Engine {
	router := gin.Default()
	router.GET("/wallet", api.GetAllWallets)
	router.GET("/wallet/:walletAddress", api.GetWalletByID)
	router.POST("/wallet", api.CreateWallet)

	return router
}

// // getAlbums responds with the list of all albums as JSON
// func getWallets(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, walletOwners)
// }

// // postAlbums adds an album form JSON received in the request body
// func postAlbums(c *gin.Context) {
// 	var newAlbum walletOwner

// 	// Call BindJSON to bind the received JSON to newAlbum
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	// Add the new album to the slice
// 	walletOwners = append(walletOwners, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

// // getAlbumByID locates the album whose ID value matches the id
// // parameter sendt by the client, then returns that album as a response
// func getAlbumByID(c *gin.Context) {
// 	walletAddress := c.Param("walletAddress")

// 	// Loop over the list of albums, looking for
// 	// an album whose ID value matches the parameter
// 	for _, a := range walletOwners {
// 		if a.WalletAddress == walletAddress {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "wallet not found"})
// }
