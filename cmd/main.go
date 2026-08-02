package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Email string
}

func (u *User) getName() {
	fmt.Println(u.Name)
}

func (u *User) getEmail() {
	fmt.Println(u.Email)
}

func (u *User) updateAge(age int) {
	u.Age = age
}
func main() {
	user := User{
		Name:  "Abhi",
		Age:   24,
		Email: "abhi@gmail.com",
	}

	user.getName()
	user.getEmail()

	fmt.Println("age is", user.Age)
	user.updateAge(45)

	fmt.Println("after update age is", user.Age)
}
