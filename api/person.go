package api

import (
	"brok/navnetjener/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllForetakForPerson returns all foretak for a person
// user provides fnr as a parameter
func GetAllForetakForPerson(context *gin.Context) {
	fnr := context.Param("fnr")

	if len(fnr) != 11 {
		context.JSON(http.StatusBadRequest, gin.H{"error": fnr + " must be 11 valid digits"})
		return
	}

	if _, err := strconv.Atoi(fnr); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fnr + " must be 11 valid digits"})
		return
	}

	capTables, err := model.FindAllCapTablesForPerson(fnr)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "finner ikke noen aksjebok for denne personen"})
		return
	}

	context.JSON(http.StatusOK, capTables)
}
