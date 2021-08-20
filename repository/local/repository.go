package local

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/praveen4g0/url-shortner/shortner"
)

type LocalRepository struct {
	database map[shortner.Redirect]int
	timeout  time.Duration
}

func NewLocalRepository() (shortner.RedirectRepository, error) {
	repo := &LocalRepository{
		database: make(map[shortner.Redirect]int),
		//timeout:  2 * time.Second, // As of now providing no timeout
	}
	return repo, nil
}

func (r *LocalRepository) Find(code string) (*shortner.Redirect, error) {
	_, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	for redirect, _ := range r.database {
		if redirect.Code == code {
			return &redirect, nil
		}
	}
	return nil, errors.New("repository.Redirect.Find")
}

func (r *LocalRepository) Store(redirect *shortner.Redirect) error {
	_, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	r.database[*redirect] = 1
	return nil
}
