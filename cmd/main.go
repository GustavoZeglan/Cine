package main

import (
	"fmt"

	"github.com/GustavoZeglan/Cine/internal/domain/services"
	psql "github.com/GustavoZeglan/Cine/internal/infra/db/postgres"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Say hello to my little friend!")
	fmt.Println("The eyes Chico, the eyes never lie!")

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env")
	}

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
	_ = services.NewMovieService(movieRepository)
	_ = services.NewRoomService(roomRepository)
	_ = services.NewSeatService(seatRepository)
	_ = services.NewSessionService(sessionRepository)
	_ = services.NewReservationService(reservationRepository)

}
