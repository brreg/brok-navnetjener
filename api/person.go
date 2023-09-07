package api

import (
	"brok/navnetjener/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllForetakForPerson returns all foretak for a person
// user provides pnr as a parameter
func GetAllForetakForPerson(context *gin.Context) {
	pnr := context.Param("pnr")

	if len(pnr) != 11 {
		context.JSON(http.StatusBadRequest, gin.H{"error": pnr + " must be 11 valid digits"})
		return
	}

	if _, err := strconv.Atoi(pnr); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": pnr + " must be 11 valid digits"})
		return
	}

	capTables, err := model.FindAllCapTablesForPerson(pnr)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "finner ikke noen aksjebok for denne personen"})
		return
	}

	context.JSON(http.StatusOK, capTables)
}
