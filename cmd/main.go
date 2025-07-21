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

	config.LoadConfig()
	oracle.Connect()

	// sql, _ := db.DB()
	// driver := sql.Driver()

	// m, _ := migrate.NewWithDatabaseInstance(
	// 	"file://internal/infra/db/migrations",
	// 	"godror", &database.Driver,
	// )

	// fmt.Println(m)

	// if err := m.Down(); err != nil && err != migrate.ErrNoChange {
	// 	log.Fatal(err)
	// }

	// log.Println("Migrations ran successfully.")

	// db.AutoMigrate(
	// 	&entities.Movie{},
	// 	&entities.Room{},
	// 	&entities.Seat{},
	// 	&entities.Session{},
	// 	&entities.Reservation{},
	// )

	// ctx := context.Background()

	// Repositories
	movieRepository := oracle.NewMovieRepository()
	roomRepository := oracle.NewRoomRepository()
	seatRepository := oracle.NewSeatRepository()
	sessionRepository := oracle.NewSessionRepository()
	reservationRepository := oracle.NewReservationRepository()

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
