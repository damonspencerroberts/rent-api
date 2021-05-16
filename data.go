package main

import (
	"strconv"
	"syreclabs.com/go/faker"
)

func handleUsers() []User {
	var users []User
	for i := 1; i <= 5; i++ {
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
			User: faker.Number().Between(1, 5),
		}
		cars = append(cars, newCar)
	}
	return cars
}

func handleParkings() []Parking {
	var parkings []Parking
	for i := 1; i <= 10; i++ {
		newParking := Parking{
			Id: strconv.Itoa(i),
			Address: faker.Address().String(),
			NumberOfCars: faker.Number().Between(1, 4),
			Images: []string{"syreclabs.com/go/faker", "syreclabs.com/go/faker", "syreclabs.com/go/faker"},
			Price: faker.Number().Between(20, 100),
			Desc: faker.Lorem().String(),
			Ammenities: faker.Lorem().Sentences(4),
			Latitude: faker.Number().Decimal(4, 2),
			Longitude: faker.Number().Decimal(4, 2),
			User: faker.Number().Between(1, 5),
		}
		parkings = append(parkings, newParking)
	}
	return parkings
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