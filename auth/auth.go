package auth

import (
	"fmt"
)

type User struct{
	UserStatus int8   //0 - simple user, 1 - moderator, 2 - admin
	Username   string
	Nickname   string
	Email      string
	Password   string
	ImgURI     string
}

type TokenAuth struct{
	User   *User
	Device string
	Token  string
}

func fakeUser() User{
	fakeUser := User{2, "gregv2", "hamlet", "email", "kfodkoksdokfksodkfk", "uri link"}
	
	return fakeUser
}

func main()  {
	fake_User := fakeUser()
	fmt.Println(fake_User)
}