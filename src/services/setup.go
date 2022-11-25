package services

import service "github.com/ozan1338/go-user-service"

var (
	UserService service.Service
)

func Setup() {
	UserService = service.CreateService("http://users-ms:8000/api/")
}