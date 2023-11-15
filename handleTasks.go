package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"context"

	"dev.azure.com/wole0010243/scheduler/_git/scheduler/internal/database"
	pb "dev.azure.com/wole0010243/scheduler/_git/scheduler/proto"
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


func (s * helloServer) CreateSchedule (ctx context.Context, req *pb.CreateScheduleRequest)(*pb.CreateScheduleResponse, error){

	actionType := req.GetType()
	expiryDate, err := time.Parse("2006-01-02",req.GetExpiryDate())

		if err !=nil{
			return nil, fmt.Errorf("unable to parse expiry date it should be in the format YYYYMMDD: %v", err)
		}
		
		task, err :=s.grpcDB.DB.CreateTask(ctx, database.CreateTaskParams{
			ID: uuid.New(),
			ExpiryDate: expiryDate,
			Type: actionType,
			Status: "ongoing",	
		});

		if err != nil{
			return nil, err
		}

		return &pb.CreateScheduleResponse{
			Id: task.ID.String(),
			Status: task.Status,
		}, nil
}


func (s*helloServer) CreateTransactionJob (ctx context.Context, req *pb.CreateTransactionJobRequest)(*pb.CreateTransactionJobResponse, error){
	Amount:= req.Amount
	BeneficiaryAccountName:=req.BeneficiaryAccountName
	BeneficiaryAccountNumber :=  req.BeneficiaryAccountNumber
	Narration :=req.Naration
	DestinationInstitutionName := req.DestinationInstitutionName
	DestinationInstitutionCode := req.DestinationInstitutionCode
	OriginatorAccountNumber := req.OriginatorAccountNumber
	CustomerAcccountName := req.CustomerAcccountName



	EndDate,err := time.Parse("2006-01-02", req.EndDate)
	
	if err !=nil{
		return nil, fmt.Errorf("unable to parse expiry date it should be in the format YYYYMMDD: %v", err)
	}

	data, err := s.grpcDB.DB.CreateTransactionJob(ctx, database.CreateTransactionJobParams{
		ID: uuid.New(),
		Amount:Amount,
		BeneficiaryAccountName: BeneficiaryAccountName,
		BeneficiaryAccountNumber: BeneficiaryAccountNumber,
		Narration: Narration,
		DestinationInstitutionName: DestinationInstitutionName,
		DestinationInstitutionCode: DestinationInstitutionCode,
		CustomerAccountName: CustomerAcccountName,
		EndDate: EndDate,
		Status: "ongoing",
		OriginatorAccountNumber: OriginatorAccountNumber,
		CreatedAt: time.Now(),
	})

	if err != nil{
		log.Printf("error creating job because %v", err)
	}
	
	log.Printf("job successfully created %v", data)
	
	return &pb.CreateTransactionJobResponse{
		Success: true,
	}, nil

}

