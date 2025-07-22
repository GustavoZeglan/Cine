package main

import (
	"fmt"
	"net/http"

	"github.com/GustavoZeglan/Cine/config"
	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/handlers"
	"github.com/GustavoZeglan/Cine/internal/infrastructure/db/postgres"
	"github.com/GustavoZeglan/Cine/internal/usecase"
	adapter "github.com/GustavoZeglan/Cine/pkg/adapter/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig()
	db := postgres.Connect()

	db.AutoMigrate(
		&entities.Movie{},
		&entities.Room{},
		&entities.Seat{},
		&entities.Session{},
		&entities.Reservation{},
	)

	// Repositories
	movieRepository := postgres.NewMovieRepository(db)
	_ = postgres.NewRoomRepository(db)
	_ = postgres.NewSeatRepository(db)
	_ = postgres.NewSessionRepository(db)
	_ = postgres.NewReservationRepository(db)

	// UseCase
	getMovies := usecase.NewGetMovies(movieRepository)
	createMovie := usecase.NewCreateMovie(movieRepository)

	// Handler
	movieHandler := handlers.NewMovieHandler(getMovies, createMovie)

	// Start the server
	router := gin.Default()
	router.POST("/movies", adapter.Handler(movieHandler.Create))
	router.GET("/movies", adapter.Handler(movieHandler.GetAll))
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
