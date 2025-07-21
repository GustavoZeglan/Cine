package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GustavoZeglan/Cine/internal/config"
	psql "github.com/GustavoZeglan/Cine/internal/infra/db/postgres"
	"github.com/GustavoZeglan/Cine/internal/services"
	adapter "github.com/GustavoZeglan/Cine/pkg/adapter/handler"
)

func main() {
	fmt.Println("Say hello to my little friend!")
	fmt.Println("The eyes Chico, the eyes never lie!")

	config.LoadConfig()

	DB := psql.Connect()
	// DB.AutoMigrate(
	// 	&entities.Movie{},
	// 	&entities.Room{},
	// 	&entities.Seat{},
	// 	&entities.Session{},
	// 	&entities.Reservation{},
	// )

	// ctx := context.Background()

	// Repositories
	movieRepository := psql.NewMovieRepository(DB)
	roomRepository := psql.NewRoomRepository(DB)
	seatRepository := psql.NewSeatRepository(DB)
	sessionRepository := psql.NewSessionRepository(DB)
	reservationRepository := psql.NewReservationRepository(DB)

	// Services
	movieService := services.NewMovieService(movieRepository)
	_ = services.NewRoomService(roomRepository)
	_ = services.NewSeatService(seatRepository)
	_ = services.NewSessionService(sessionRepository)
	_ = services.NewReservationService(reservationRepository)

	// Start the server
	router := adapter.NewMuxAdapter()
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
