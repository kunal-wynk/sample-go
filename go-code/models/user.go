package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	Lastname  string
}

// memory alignment -->

// ROAD MAP

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("User with id not allowed")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with Id `%v` not found", u.ID)
}

func DeleteUser(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with Id `%v` not found", id)
}

func GetUserById(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}
	return User{}, fmt.Errorf("User with Id `%v` not found", id)
}
