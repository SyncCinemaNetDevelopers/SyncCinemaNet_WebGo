package auth

import (
	"fmt"
)

type User struct{
	Username string
	Nickname string
	Email    string
	Password string
	ImgURI   string
}

type TokenAuth struct{
	User   User
	Device string
	Token  string
}

func fakeUser() {
	fakeUser = User("gregv2", "hamlet", "email", "kfodkoksdokfksodkfk", "uri link")
	
	return fakeUser
}

func main()  {
	fmt.print()
}