package main

import (
	"context"
	"log"
	"math"
	"sync"

	"dev.azure.com/wole0010243/scheduler/_git/scheduler/internal/database"
	"github.com/robfig/cron"
)



func (apiCfg * apiConfig) startJob (){
	c := cron.New()


	err := c.AddFunc("0 0 * * *", func(){
		procesValidJobs(apiCfg)
	});

	if err!= nil{
		log.Fatal("Error scheduling job:", err)
	}

	c.Start()
}


func procesValidJobs(apiCfg *apiConfig){
	data, err:= apiCfg.DB.FindValidJob(context.Background())

	if err != nil{
		log.Printf("Error getting datas, %v", err)
	}

	totalJobs := len(data);
	bactchCount:= int(math.Ceil(float64(totalJobs)/200))

	wg := &sync.WaitGroup{};


	for i :=0; i<=bactchCount; i++{
		wg.Add(1)
		startIndex := i*200
		endIndex := (i+1)*200
		
		go processBatch(data[startIndex:endIndex], wg)

	}

	wg.Wait()

}


func processBatch(batch []database.Task, wg *sync.WaitGroup){
	defer wg.Done()
	for _, job := range batch{
		//start sending request to the respective services

		if job.Type =="data"{
			// send request to data service
		} else if job.Type=="transfer" {
			// send to transfer service
		} else if job.Type =="bill"{
			// send to bill service
		}else if job.Type =="airtime"{
			//send request to airtime service
		}		
	}	
}
