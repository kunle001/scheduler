package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"dev.azure.com/wole0010243/scheduler/_git/scheduler/internal/database"
	"github.com/google/uuid"
)

func ( apiCfg *apiConfig) CreateTask(w http.ResponseWriter, r *http.Request) {
	
	type parameters struct{
		Type string `json:"action"`
		Expiry_date string `json:"expiry_date"`
	}

	decoder := json.NewDecoder(r.Body)

	params:= parameters{}
	err := decoder.Decode(params)
	
	expiryDate, err := time.Parse("2006-01-02", params.Expiry_date)


	if err!=nil{
		respondWithError(w, 500, fmt.Sprintf("Error parsing data %v", err))
	};

	task, err:= apiCfg.DB.CreateTask(r.Context(), database.CreateTaskParams{
		ID: uuid.New(),
		ExpiryDate: expiryDate,
		Type: params.Type,
		Status: "pending",		
	})

	if err!= nil{
		respondWithError(w, 500, fmt.Sprintf("Error creating task %v", err))
	};
	
	// perform the task instantly

	respondWithJson(w, 201, task)
}
