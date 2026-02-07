package user

type Repository interface {
	FindAll() ([]User, error)
}
