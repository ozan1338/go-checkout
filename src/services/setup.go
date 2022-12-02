package services

import (
	"fmt"
	"os"

	service "github.com/ozan1338/go-user-service"
)

var (
	UserService service.Service
)

func Setup() {
	UserService = service.CreateService(fmt.Sprintf("%s/api/",os.Getenv("USERS_MS")))
}