package handlers

import (
	"github.com/brianvoe/gofakeit/v7"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	name string
	email string
	phone string
	creditcard string
}

func GenerateRandomUsers(count int) []User {
	userList := make([]User, 0, count)
	for i := 0; i < count; i++ {
		newUser := User{gofakeit.Name(), gofakeit.Email(), gofakeit.Phone(), gofakeit.CreditCardNumber(&gofakeit.CreditCardOptions{})}
		userList = append(userList, newUser)
	}
	return userList
}

func GenerateRandomUsersAsBSON(count int) []interface{} {
	userList := make([]interface{}, 0, count)
	for i := 0; i < count; i++ {
		newUser := bson.D{{"name", gofakeit.Name()}, {"email", gofakeit.Email()}, {"phone", gofakeit.Phone()}, {"creditcard", gofakeit.CreditCardNumber(&gofakeit.CreditCardOptions{})}}
		userList = append(userList, newUser)
	}
	return userList
}

func GeneratePhoneNumber() string {
	return gofakeit.Phone()
}