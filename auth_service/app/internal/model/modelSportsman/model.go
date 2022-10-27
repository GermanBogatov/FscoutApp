package modelSportsman

import "time"

type SportsmanDTO struct {
	Sportsman_uuid string    `json:"sportsman_uuid"`
	Name           string    `json:"name" binding:"required"`
	Surname        string    `json:"surname" binding:"required"`
	Phone          string    `json:"phone" binding:"required"`
	Email          string    `json:"email" binding:"required"`
	Password       string    `json:"password" binding:"required"`
	Academy        string    `json:"academy"`
	Country_uuid   int       `json:"country_uuid"` // TODO не забыть изменить
	Address        string    `json:"address"`
	City           string    `json:"city"`
	State          string    `json:"state"`
	Index          int       `json:"index"`
	Birthday       time.Time `json:"birthday"`
	Height         string    `json:"height"`
	Weight         string    `json:"weight"`
	Strong_leg     string    `json:"strong_Leg"`
	Time_create    time.Time `json:"time_create"`
	Role_uuid      string    `json:"role_uuid"`
}

type SignInDTO struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
