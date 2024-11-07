package repository

import (
	"context"
	"crud_golang_mysql/entity"
	"database/sql"
	"errors"
	"strconv"
)

type profileRepositoryImpl struct {
	DB *sql.DB
}

func NewProfileRepositoryImpl(db *sql.DB) ProfileRepository {
	return &profileRepositoryImpl{DB: db}
}
func (repository *profileRepositoryImpl) Insert(ctx context.Context, profile entity.Profile) (entity.Profile, error) {
	script := "INSERT INTO profile (nama, umur, alamat) values (?,?,?)"
	result, err := repository.DB.ExecContext(ctx, script, profile.Name, profile.Umur, profile.Alamat)
	if err != nil {
		return profile, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return profile, err
	}
	profile.Id = int(id)
	return profile, nil
}

func (repository *profileRepositoryImpl) Update(ctx context.Context, profile entity.Profile, id int) (entity.Profile, error) {
	script := "UPDATE profile SET nama = ?, umur = ?, alamat = ? WHERE id = ?"
	result, err := repository.DB.ExecContext(ctx, script, profile.Name, profile.Umur, profile.Alamat, id)
	if err != nil {
		return profile, err
	}
	number, err := result.LastInsertId()
	if err != nil {
		return profile, err
	}
	profile.Id = int(number)
	return profile, nil
}

func (repository *profileRepositoryImpl) FindById(ctx context.Context, id int) (entity.Profile, error) {
	script := "SELECT nama, umur, alamat FROM profile WHERE id = ? limit 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	profile := entity.Profile{}
	if err != nil {
		return profile, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&profile.Name, &profile.Umur, &profile.Alamat)
		return profile, nil
	} else {
		// tidak ada
		return profile, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repository *profileRepositoryImpl) DeleteById(ctx context.Context, id int) (entity.Profile, error) {
	script := "DELETE from profile WHERE id = ?"
	result, err := repository.DB.ExecContext(ctx, script, id)
	profile := entity.Profile{}
	if err != nil {
		return profile, err
	}
	number, err := result.LastInsertId()
	if err != nil {
		return profile, err
	}
	profile.Id = int(number)
	return profile, nil
}

func (repository *profileRepositoryImpl) FindAll(ctx context.Context) ([]entity.Profile, error) {
	script := "SELECT nama, umur, alamat FROM profile"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var profiles []entity.Profile
	for rows.Next() {
		profile := entity.Profile{}
		rows.Scan(&profile.Name, &profile.Umur, &profile.Alamat)
		profiles = append(profiles, profile)
	}
	return profiles, nil
}
