package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"superapps/controllers"
	helper "superapps/helpers"
	middleware "superapps/middlewares"
	"superapps/services"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	services.InitDBs()

	if err != nil {
		helper.Logger("error", "Error getting env")
	}

	router := mux.NewRouter()

	router.Use(middleware.CorsMiddleware)
	router.Use(middleware.JwtAuthentication)

	errMkidr := os.MkdirAll("public", os.ModePerm)
	if errMkidr != nil {
		log.Fatalf("Failed to create or access directory: %v", err)
	}

	dir, err := os.Open("public")
	if err != nil {
		log.Fatalf("Failed to open public directory: %v", err)
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		log.Fatalf("Failed to read directory contents: %v", err)
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			staticPath := "/" + fileInfo.Name() + "/"
			publicPath := "./public/" + fileInfo.Name() + "/"

			log.Printf("Serving static files from %s at %s", publicPath, staticPath)

			router.PathPrefix(staticPath).Handler(http.StripPrefix(staticPath, http.FileServer(http.Dir(publicPath))))
		}
	}

	// Admin
	router.HandleFunc("/api/v1/admin/list/project", controllers.AdminListProject).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/admin/list/user", controllers.AdminListUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/admin/verify/user", controllers.UpdateAdminVerifyUser).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/v1/admin/verify/project", controllers.UpdateAdminVerifyProject).Methods("PUT", "OPTIONS")

	// Auth
	router.HandleFunc("/api/v1/auth/login", controllers.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/auth/register", controllers.Register).Methods("POST", "OPTIONS")

	// Profile
	router.HandleFunc("/api/v1/profile", controllers.GetProfile).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/profile/update", controllers.UpdateProfile).Methods("PUT", "OPTIONS")

	// Account
	router.HandleFunc("/api/v1/account/update", controllers.UpdateAccount).Methods("PUT", "OPTIONS")

	// Project
	router.HandleFunc("/api/v1/project/list", controllers.ProjectList).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/project/detail/{id}", controllers.ProjectDetail).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/project/store", controllers.ProjectStore).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/project/update", controllers.ProjectUpdate).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/v1/project/delete", controllers.ProjectDelete).Methods("DELETE", "OPTIONS")

	portEnv := os.Getenv("PORT")
	port := ":" + portEnv

	fmt.Println("Starting server at", port)

	server := &http.Server{
		Addr:              port,
		Handler:           router,
		ReadHeaderTimeout: 3 * time.Second,
	}

	errListenAndServe := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", errListenAndServe)
	}
}
