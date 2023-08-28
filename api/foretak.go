package api

import (
	"brok/navnetjener/model"
	"net/http"
	"strconv"

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
		context.JSON(http.StatusNotFound, gin.H{"error": "finner ikke Aksjebok for orgnr " + orgnr})
		return
	}

	context.JSON(http.StatusOK, capTable)
}

// GetForetak returns 25 foretak from TheGraph and the database
// user can use the queryparameter "page" to paginate
func GetForetak(context *gin.Context) {
	page := context.Query("page")
	safePage := model.SanitizeString(page)

	// convert page to int
	// if page is empty, set it to 0
	// if page is not a number, set it to 0
	// if page is less than 0, set it to 0
	if safePage == "" {
		safePage = "0"
	}

	safePageInt, err := strconv.Atoi(safePage)

	if err != nil {
		safePageInt = 0
	}

	capTables, err := model.FindForetak(safePageInt)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, capTables)
}
