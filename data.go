package main

import (
	"strconv"
	"syreclabs.com/go/faker"
)

func handleUsers() []User {
	var users []User
	for i := 0; i <= 5; i++ {
		newUser := User{
			Id: strconv.Itoa(i),
			FirstName: faker.Name().FirstName(),
			LastName: faker.Name().LastName(),
			Email: faker.Internet().Email(),
			PhoneNumber: faker.PhoneNumber().CellPhone(),
		}
		users = append(users, newUser)
	}
	return users
}

func handleCars() []Car {
	var cars []Car
	for i := 1; i <= 5; i++ {

		newCar := Car{
			Id: strconv.Itoa(i),
			Model: faker.Lorem().Word(),
			Color: faker.Commerce().Color(),
			Plate: faker.Company().Ein(),
			User: faker.Number().Positive(5),
		}
		cars = append(cars, newCar)
	}
	return cars
}

// func findUser(id int) User {
// 	newId := strconv.Itoa(id)
// 	var foundUser User
// 	for _, user := range Users {
// 		if user.Id == newId {
// 			foundUser = user
// 		}
// 	}
// 	return foundUser
// }