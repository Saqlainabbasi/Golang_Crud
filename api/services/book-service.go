package services

import (
	"net/http"

	"github.com/Saqlainabbasi/api/models"
	"github.com/Saqlainabbasi/api/utils"
)

// func to get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	//pass to Database layer
	resp, data := models.GetBooks()
	if resp.Error != nil {
		utils.WriteJson(w, http.StatusInternalServerError, resp.Error)
	}
	// send the responce back
	utils.WriteJson(w, http.StatusOK, data)
}
