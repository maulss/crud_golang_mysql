package repository

import (
	"context"
	"crud_golang_mysql"
	"crud_golang_mysql/entity"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestInsertProfile(t *testing.T) {
	profileRepository := NewProfileRepositoryImpl(crud_golang_mysql.GetConnection())
	ctx := context.Background()
	profile := entity.Profile{
		Name:   "Maulana",
		Umur:   21,
		Alamat: "Jln Test Insert 3",
	}
	profile, err := profileRepository.Insert(ctx, profile)
	if err != nil {
		panic(err)
	}
	fmt.Println(profile)
}

func TestUpdateProfile(t *testing.T) {
	profileRepository := NewProfileRepositoryImpl(crud_golang_mysql.GetConnection())
	ctx := context.Background()
	profile := entity.Profile{
		Name:   "Givari",
		Umur:   10,
		Alamat: "Jln Test Update By Id",
	}

	profilr, err := profileRepository.Update(ctx, profile, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(profilr)
}

func TestDeleteProfile(t *testing.T) {
	profileRepository := NewProfileRepositoryImpl(crud_golang_mysql.GetConnection())
	ctx := context.Background()

	profile, err := profileRepository.DeleteById(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(profile)
}

func TestFindProfile(t *testing.T) {
	profileRepository := NewProfileRepositoryImpl(crud_golang_mysql.GetConnection())
	ctx := context.Background()
	profile, err := profileRepository.FindById(ctx, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(profile)
}

func TestFindAllProfiles(t *testing.T) {
	profileRepository := NewProfileRepositoryImpl(crud_golang_mysql.GetConnection())
	ctx := context.Background()
	allProfiles, err := profileRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(allProfiles)
}
