package repository

import (
	"context"
	"crud_golang_mysql/entity"
)

type ProfileRepository interface {
	Insert(ctx context.Context, profile entity.Profile) (entity.Profile, error)
	Update(ctx context.Context, profile entity.Profile, id int) (entity.Profile, error)
	FindById(ctx context.Context, id int) (entity.Profile, error)
	DeleteById(ctx context.Context, id int) (entity.Profile, error)
	FindAll(ctx context.Context) ([]entity.Profile, error)
}
