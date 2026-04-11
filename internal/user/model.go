package user

type User struct {
	ID           string `db:"id"`
	FirstName    string `db:"first_name"`
	SecondName   string `db:"second_name"`
	Birthdate    string `db:"birthdate"`
	Biography    string `db:"biography"`
	City         string `db:"city"`
	PasswordHash string `db:"password_hash"`
}
