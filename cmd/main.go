package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GustavoZeglan/Cine/internal/config"
	"github.com/GustavoZeglan/Cine/internal/domain/entities"
	"github.com/GustavoZeglan/Cine/internal/infrastructure/db/postgres"
	"github.com/GustavoZeglan/Cine/internal/services"
	adapter "github.com/GustavoZeglan/Cine/pkg/adapter/handler"
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
	roomRepository := postgres.NewRoomRepository(db)
	seatRepository := postgres.NewSeatRepository(db)
	sessionRepository := postgres.NewSessionRepository(db)
	reservationRepository := postgres.NewReservationRepository(db)

	// Services
	movieService := services.NewMovieService(movieRepository)
	_ = services.NewRoomService(roomRepository)
	_ = services.NewSeatService(seatRepository)
	_ = services.NewSessionService(sessionRepository)
	_ = services.NewReservationService(reservationRepository)

	// Start the server
	router := adapter.NewGinAdapter()
	router.Get("/movies", func(w http.ResponseWriter, r *http.Request) {
		movies, err := movieService.GetAllMovies(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(movies)
	})
	err := http.ListenAndServe(":8080", router.Router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
