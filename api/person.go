package api

import (
	"brok/navnetjener/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllForetakForPerson returns all foretak for a person
// user provides pnr as a parameter
func GetAllForetakForPerson(context *gin.Context) {
	pnr := context.Param("pnr")

	capTables, err := model.FindAllCapTablesForPerson(pnr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, capTables)
}
