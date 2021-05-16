package main

func handleUsers() []User {
	return []User{
		User{Id: "1", FirstName: "John", LastName: "Doe", Email: "john@gmail.com", PhoneNumber: "07644655293"},
		User{Id: "2", FirstName: "Melanie", LastName: "Smith", Email: "ms10101@aol.com", PhoneNumber: "1474558011"},
		User{Id: "3", FirstName: "Leonardo", LastName: "DiCaprio", Email: "leo@dicap.com", PhoneNumber: "2123459862"},
	}
}

func handleCars() []Car {
	return []Car{
		Car{Id: "1", Model: "Mazda N5", Color: "Red", Plate: "4C4-4567", User: "1"},
		Car{Id: "2", Model: "Alfa Romeo Giulia", Color: "Black", Plate: "548-7781", User: "2"},
		Car{Id: "3", Model: "Toyota Corolla", Color: "Grey", Plate: "657-2HHM", User: "3"},
	}
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