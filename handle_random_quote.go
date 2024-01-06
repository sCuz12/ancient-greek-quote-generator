package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sCuz12/ancient-greek-quote-api/models"
	"github.com/sCuz12/ancient-greek-quote-api/services"
)


func getRandomQuote (counterService *services.CounterService ,w http.ResponseWriter, r *http.Request) {

	 // Increment the counter
	 if err := counterService.IncrementCounter(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    quotes := &[]models.Quote{}
	//get random , retrieve only one 
    err := DB.Order("RANDOM()").First(quotes).Error
	
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