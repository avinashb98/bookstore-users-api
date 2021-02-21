package app

import (
	"github.com/avinashb98/bookstore-users-api/controller/ping"
	"github.com/avinashb98/bookstore-users-api/controller/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/search", user.SearchUser)
	router.GET("/user/:user_id", user.GetUser)
	router.POST("/user", user.CreateUser)
}