package services

import (
	"sesi6-gin/httpserver/controllers/params"
	"sesi6-gin/httpserver/controllers/views"
	"sesi6-gin/httpserver/repositories/models"

	"golang.org/x/crypto/bcrypt"
)

// static data users
var users = []models.User{
	{
		ID:       1,
		Username: "Deeva",
		Password: "password123",
	},
	{
		ID:       2,
		Username: "Rindu",
		Password: "password123",
	},
}

func CreateUser(req *params.UserCreateRequest) *views.Response {
	// step : (4) buat model
	var user models.User

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return views.BadRequestError(err)
	}

	user.ID = 1
	user.Password = string(hash)
	user.Username = req.Username

	// string(hash) users password
	for _, user := range users {
		user.Password = string(hash)
	}

	// step : (5) kirim ke repositories
	// err = repositories.CreateUser(&model)

	// step : (7) buat sebuah views
	v := views.SuccessCreateResponse(users, "created success!")

	// step : (8) kembalikan views ke controller
	return v
}

func GetUser() *views.Response {
	v := views.GetDataResponse(users)
	return v
}
