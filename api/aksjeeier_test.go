package api_test

import (
	"brok/navnetjener/api"
	"brok/navnetjener/model"
	"brok/navnetjener/utils"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**

!-------OBS!-----OBS!-----OBS!-----OBS!-----OBS!-----OBS!-------!
!																																!
! 	Dette settet med tester fungerer KUN når du har deployet		!
!		en lokal versjon av The Graph																!
!   https://github.com/brreg/brok                               !
!																																!
!-------OBS!-----OBS!-----OBS!-----OBS!-----OBS!-----OBS!-------!

*/

func TestFindWalletsForPerson(t *testing.T) {
	router := utils.Setup()
	router.GET("/v1/aksjeeier/:id", api.GetAllForetakForAksjeeier)

	fnr := "21058000000"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/aksjeeier/"+fnr, nil)
	router.ServeHTTP(w, req)

	var receivedCapTable []model.CapTable
	json.Unmarshal([]byte(w.Body.String()), &receivedCapTable)

	assert.Equal(t, http.StatusOK, w.Code)

	if result := contains(receivedCapTable, fnr); result != true {
		t.Errorf("Expected receivedCaptable %v to contain a owner with birth date %s", receivedCapTable, fnr)
	}

	assert.Equal(t, 1, len(receivedCapTable))

}

func TestFindWalletsForCompany(t *testing.T) {
	router := utils.Setup()
	router.GET("/v1/aksjeeier/:id", api.GetAllForetakForAksjeeier)

	orgnr := "310812277"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/aksjeeier/"+orgnr, nil)
	router.ServeHTTP(w, req)

	var receivedCapTable []model.CapTable
	json.Unmarshal([]byte(w.Body.String()), &receivedCapTable)

	assert.Equal(t, http.StatusOK, w.Code)

	if result := contains(receivedCapTable, orgnr); result != true {
		t.Errorf("Expected receivedCaptable %v to contain a owner with orgnr %s", receivedCapTable, orgnr)
	}

	assert.Equal(t, 1, len(receivedCapTable))
}

func TestShouldFailWithBadRequest(t *testing.T) {
	router := utils.Setup()
	router.GET("/v1/aksjeeier/:id", api.GetAllForetakForAksjeeier)

	badvalue := "3108122771111"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/aksjeeier/"+badvalue, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestShouldFailWithBadRequest2(t *testing.T) {
	router := utils.Setup()
	router.GET("/v1/aksjeeier/:id", api.GetAllForetakForAksjeeier)

	badvalue := "aaaaaaaaa"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/aksjeeier/"+badvalue, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// Contains checks if the owner is present in the list of TokenHolders
func contains(captables []model.CapTable, id string) bool {
	for _, captable := range captables {
		for _, tokenHolder := range captable.TokenHolders {
			if tokenHolder.Owner.Person.BirthDate == id[0:6] {
				return true
			}
			if tokenHolder.Owner.Company.Orgnr == id {
				return true
			}
		}
	}
	return false
}
