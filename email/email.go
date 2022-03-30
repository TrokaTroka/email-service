package email

type from struct {
	Email    string
	Password string
}

type to struct {
	Email string
}

type Email struct {
	ID   string
	From from
	To   to
}

func New() (e *Email) {
	e = &Email{}
	return
}