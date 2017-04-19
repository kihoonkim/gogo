package user

import "fmt"

type Repository interface {
	Save(user *User) User
	FindAll() []User
	FindById(id string) User
}

type GormRepository struct{}

func (repo *GormRepository) Save(user *User) User {
	fmt.Println("SAVE")

	return User{}
}

func (repo *GormRepository) FindAll() []User {
	fmt.Println("FindAll")
	var users []User
	return users
}

func (repo *GormRepository) FindById(id string) User {
	fmt.Println("FindById")
	return User{}
}
