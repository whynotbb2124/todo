package user

var Users = make(map[string]User)

type User struct {
	Login    string
	Password string
	Name     string
}

func Create(login, password, name string) {
	var u User

	u.Login = login
	u.Name = name
	u.Password = password

	Users[u.Login] = u

}

func (u *User) Update() {
	Users[u.Login] = *u
}

func (u *User) Delete() {
	delete(Users, u.Login)
}

func Read(login string) User {
	user := Users[login]
	return user
}
