package main

import (
	"fmt"

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
	_ = psql.NewMovieRepository(DB)
	_ = psql.NewRoomRepository(DB)
	_ = psql.NewSeatRepository(DB)
	_ = psql.NewSessionRepository(DB)
	_ = psql.NewReservationRepository(DB)

}
