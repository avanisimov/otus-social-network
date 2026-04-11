package user

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreateUser(user User) (string, error) {
	return "new-user-id", nil
}
