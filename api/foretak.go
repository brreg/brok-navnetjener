package api

import (
	"brok/navnetjener/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetForetakByOrgnr(context *gin.Context) {
	orgnr := context.Param("orgnr")

	capTable, err := model.FindCaptableByOrgnr(orgnr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if foretak is empty

	if capTable.Name == "" {
		context.JSON(http.StatusNotFound, gin.H{"error": "finner ikke foretak med orgnr " + orgnr})
		return
	}

	context.JSON(http.StatusOK, capTable)
}
