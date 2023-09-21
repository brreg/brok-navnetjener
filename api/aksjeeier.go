package api

import (
	"brok/navnetjener/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// GetAllForetak returns all foretak for a person or organization
// user provides fnr as a parameter
func GetAllForetakForAksjeeier(context *gin.Context) {
	id := context.Param("id")
	safeId := model.SanitizeString(id)

	if len(safeId) != 11 && len(safeId) != 9 {
		logrus.Warn("error: id må være et gyldig orgnr eller fnr")
		context.JSON(http.StatusBadRequest, gin.H{"error": safeId + " må være et gyldig orgnr eller fnr"})
		return
	}

	if _, err := strconv.Atoi(safeId); err != nil {
		logrus.Warn("error: id må være et tall")
		context.JSON(http.StatusBadRequest, gin.H{"error": safeId + " må være et tall"})
		return
	}

	capTables, err := model.FindAllCapTables(safeId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "finner ikke noen aksjebok for denne id'en"})
		return
	}

	context.JSON(http.StatusOK, capTables)
}
