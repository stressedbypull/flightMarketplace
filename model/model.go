package model

import (
	"log"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// User model entity
type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	Username string    `gorm:"size:100;unique" json:"username"`
	Email    string    `gorm:"size:255;unique" json:"email"`
	Password string    `gorm:"size:255" json:"password"`
	Role     string    `gorm:"size:50" json:"role"`
	// Other user attributes as needed
}

// Define the table name for User model
func (User) TableName() string {
	return "users"
}

// Order model entity
type Order struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	UserID        uuid.UUID `gorm:"type:uuid;index;column:user_id" json:"userId"`
	FlightID      uuid.UUID `gorm:"type:uuid;index;column:flight_id" json:"flightId"`
	BookingDate   time.Time `gorm:"column:booking_date" json:"bookingDate"`
	PaymentStatus string    `gorm:"size:50" json:"paymentStatus"`
	// Other order attributes as needed
}

// Define the table name for Order model
func (Order) TableName() string {
	return "orders"
}

// Flight model entity
type Flight struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	Origin            string    `gorm:"size:70" json:"origin"`
	Destination       string    `gorm:"size:70" json:"destination"`
	DepartureDateTime time.Time `gorm:"column:departure_date_time" json:"departureDateTime"`
	ArrivalDateTime   time.Time `gorm:"column:arrival_date_time" json:"arrivalDateTime"`
	Price             float64   `gorm:"type:numeric" json:"price"`
	AvailableSeats    int       `json:"availableSeats"`
	// Other flight attributes as needed
}

// Define the table name for Flight model
func (Flight) TableName() string {
	return "flights"
}

// Model entities for fake tokens
type Token struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// BeforeSave will set a UUID rather than numeric ID.
func (r *User) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID, err = uuid.NewV4()

	if err != nil {
		log.Println("Error during generation UUID for user before creation. Caused by " + err.Error())
		return err
	}
	return nil
}

// BeforeSave will set a UUID rather than numeric ID.
func (c *Flight) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID, err = uuid.NewV4()

	if err != nil {
		log.Println("Error during generation UUID for Flight before creation. Caused by " + err.Error())
		return err
	}
	return nil
}

// BeforeSave will set a UUID rather than numeric ID.
func (s *Order) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID, err = uuid.NewV4()

	if err != nil {
		log.Println("Error during generation UUID for Order before creation. Caused by " + err.Error())
		return err
	}
	return nil
}

// generate a function for each field of the struct
func (r *Order) AfterCreate(tx *gorm.DB) (err error) {
	//log.Println("AfterCreate")
	return nil
}
