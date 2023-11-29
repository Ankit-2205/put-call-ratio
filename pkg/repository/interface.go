package repository

type Repository interface {
	GetOptionChain(string) (OptionChain, error)
}
