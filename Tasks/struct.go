package main

import (
	"fmt"
)

type User struct {
	Name string
}

func reset(u *User) {
	u.Name = "Oleg3"

	fmt.Println(u.Name)

	// u = &User{
	// 	Name: "Oleg4",
	// }
}

func updateUser(u User) {
	u.Name = "Oleg2"

	fmt.Println("2:", u.Name)

	reset(&u)
}

func main() {
	user := User{
		Name: "Oleg",
	}

	fmt.Println("1:", user.Name)

	updateUser(user)
}
