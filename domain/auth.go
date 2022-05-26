package domain

type AuthRepo interface {
	FindBy(string, string) (*Login, error)
}
