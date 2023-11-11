package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"dev.azure.com/wole0010243/scheduler/_git/scheduler/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct{
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	port:= os.Getenv("PORT")
	dbUrl:= os.Getenv("DB_URL")

	if port==""{
		log.Fatal("Port is not found in the env")
	}
	
	if dbUrl==""{
		log.Fatal("DB URL is not found")
	};

	// connecting to DB
	conn, err := sql.Open("postgres", dbUrl)

	if err != nil{
		log.Fatal("can't connect to the database because:", err)
	}

	// DataBase
	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	};

	
	router := chi.NewRouter()
	
	// Allowing cross-origin connection
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods:[]string{"POST", "GET", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge:  300,
	}));

	// api version 1 router, mounting it on the ain router
	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", HandleReadiness)
	v1Router.Post("/create-task", apiCfg.CreateTask)

	router.Mount("/api/v1",v1Router)

	srv := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}


	log.Printf("Sever starting on port %v", port)
	err = srv.ListenAndServe();

	if err != nil{
		log.Fatal(err)
	}

	
}