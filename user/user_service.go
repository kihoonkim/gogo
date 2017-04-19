package user

type Service struct {
	DB     Repository      `inject:""`
	GormDB *GormRepository `inject:""`
}

func (s *Service) Save(user *User) *User {
	s.DB.Save(user)
	return user
}

func (s *Service) FindAll() (users []User) {
	users = s.DB.FindAll()
	return
}

func (s *Service) FindById(id string) (users []User) {
	users = s.DB.FindAll()
	return
}
