package repositories

import (
	"go-learn/repositories/auth_repo"
)

type Repo struct {
	AuthRepo auth_repo.AuthContract
}

func NewRepo() *Repo {
	return &Repo{
		AuthRepo: auth_repo.NewAuthRepositories(),
	}
}
