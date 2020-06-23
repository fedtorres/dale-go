package app

import (
	"github.com/fedtorres/dale-go/src/api/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
