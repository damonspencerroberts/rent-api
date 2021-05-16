package main

import (
	"strconv"
	"syreclabs.com/go/faker"
)

func createUsers() []User {
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

func createCars() []Car {
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

func createParkings() []Parking {
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

func createBookings() []Booking {
	var bookings []Booking
	for i := 1; i <= 15; i++ {
		newBooking := Booking{
			Id: strconv.Itoa(i),
			StartDate: faker.Date().Forward(100000),
			EndDate: faker.Date().Forward(100000),
			StartTime: faker.Time().Forward(100000),
			EndTime: faker.Time().Forward(100000),
			Car: faker.Number().Between(1, 5),
			Parking: faker.Number().Between(1, 10),
		}
		bookings = append(bookings, newBooking)
	}
	return bookings
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