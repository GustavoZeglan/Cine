package main

import (
	"github.com/GustavoZeglan/Cine/internal/config"
	"github.com/GustavoZeglan/Cine/internal/domain/services"
	oracle "github.com/GustavoZeglan/Cine/internal/infrastructure/db/oracle"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
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
	_ = services.NewMovieService(movieRepository)
	_ = services.NewRoomService(roomRepository)
	_ = services.NewSeatService(seatRepository)
	_ = services.NewSessionService(sessionRepository)
	_ = services.NewReservationService(reservationRepository)

	// movieService.CreateMovie(ctx, &entities.Movie{Title: "Scarface"})

}
