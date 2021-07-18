package service

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Title     string `json:"title"`
}

type Users []User // Alias type

type UserService struct {
}

func (us *UserService) GetUsers() Users {
	u := Users{
		User{
			Firstname: "f1",
			Lastname:  "l1",
			Title:     "Mr.",
		},
		User{
			Firstname: "f2",
			Lastname:  "l2",
			Title:     "Miss.",
		},
	}
	return u
}