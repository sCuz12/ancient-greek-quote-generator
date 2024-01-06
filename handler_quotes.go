package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sCuz12/ancient-greek-quote-api/models"
)

type Response struct {
	Quotes    []models.Quote `json:"quotes"`
	Datetime  time.Time	     `json:"extraData"`
}

func GetAllGreekQuotes(w http.ResponseWriter, r *http.Request) {
    quotes := &[]models.Quote{}
	
    err := DB.Find(quotes).Error
	
	if (err != nil) {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	response := Response{
		Quotes: *quotes,
		Datetime :  time.Now(),
	}

    // Handle response...
	json.NewEncoder(w).Encode(response)
}