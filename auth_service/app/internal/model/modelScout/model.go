package modelScout

import "time"

type ScoutDTO struct {
	Scout_uuid   string    `json:"scout_uuid"`
	Name         string    `json:"name" binding:"required"`
	Surname      string    `json:"surname" binding:"required"`
	Phone        string    `json:"phone" binding:"required"`
	Email        string    `json:"email" binding:"required"`
	Password     string    `json:"password" binding:"required"`
	Company      string    `json:"company"`
	Country_uuid int       `json:"country_uuid"` // TODO не забыть изменить
	Address      string    `json:"address"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	Zipcode      int       `json:"zipcode"`
	Gender       string    `json:"gender"`
	Birthday     time.Time `json:"birthday"`
	Vat_number   string    `json:"vat_number"`
	Passport     string    `json:"passport"`
	Time_create  time.Time `json:"time_create"`
	Role_uuid    string    `json:"role_uuid"`
	Confirmed    bool      `json:"confirmed"`
}

type SignInDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
