package main

import (
	"net/http"
	"project/controllers"
)

func main() {

	//u := models.User{FirstName: "kunal", ID: 2, Lastname: "sharma"}
	//fmt.Println(u)

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
	// http://localhost:3000/users
}
