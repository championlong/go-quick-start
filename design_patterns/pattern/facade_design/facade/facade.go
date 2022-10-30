package facade

type IUser interface {
	Login(phone int, code int) (*User, error)
	Register(phone int, code int) (*User, error)
}

type IUserFacade interface {
	LoginOrRegister(phone int, code int) error
}

type User struct {
	Name string
}

type UserService struct {}

func (u UserService) Login(phone int, code int) (*User, error) {
	return &User{Name: "test login"}, nil
}

func (u UserService) Register(phone int, code int) (*User, error) {
	return &User{Name: "test register"}, nil
}

func (u UserService)LoginOrRegister(phone int, code int) (*User, error) {
	user, err := u.Login(phone, code)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	return u.Register(phone, code)
}