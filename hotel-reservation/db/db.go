package db

const (
	DBNAME     = "hotel-reservation"
	TESTDBNAME = "hotel-reservation-test"
	DBURI      = "mongodb://localhost:27017"
)

type Store struct {
	Users  UserStore
	Hotels HotelStore
	Rooms  RoomStore
}
