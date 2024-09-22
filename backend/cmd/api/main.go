package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/veise3/learning-record-app/config"
	"github.com/veise3/learning-record-app/internal/delivery/http/handler"
	"github.com/veise3/learning-record-app/internal/delivery/http/middleware"
	"github.com/veise3/learning-record-app/internal/infrastructure/database"
	"github.com/veise3/learning-record-app/internal/repository"
	"github.com/veise3/learning-record-app/internal/usecase"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.NewPostgresDB(&config)
	if err != nil {
		log.Fatalf("Failed to connet to database: %v", err)
	}

	repo := repository.NewLearningRecordRepository(db)
	useCase := usecase.NewLearningRecordUseCase(repo)
	handler := handler.NewLearningRecordHandler(useCase)

	r := gin.Default()

	// CORS設定
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	r.Use(middleware.ErrorMiddleware())

	api := r.Group("/api")
	{
		api.POST("/learning", handler.CreateLearningRecord)
		api.GET("/learning", handler.GetLearningRecords)
		api.DELETE("/learning/:id", handler.DeleteLearningRecord)
		api.PUT("/learning/:id", handler.UpdateLearningRecord)

	}

	// if err := r.Run(":8080"); err != nil {
	// 	log.Fatalf("Failed to start server: %v", err)
	// }

	if err := r.Run(":" + config.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
