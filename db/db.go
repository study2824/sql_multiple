package db

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/boil"
	_ "github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql/driver"
	"sql_multiple1/config"
	"sql_multiple1/models"
	"strconv"
)

var db *sql.DB

//init
func Init() {
	//databaseURL := os.Getenv("DATABASE_URL")
	tmp, err := sql.Open("postgres", config.ConnectingDB())
	if err != nil {
		panic(err)
	}
	db = tmp
}

// get all users
func GetAllUsers() (models.UserSlice, error) {
	return models.Users().All(context.Background(), db)
}

// get one user
func GetOneUser(n string) (*models.User, error) {
	id, err := strconv.Atoi(n)
	if err != nil {
		return nil, err
	}

	return models.FindUser(context.Background(), db, id)
}

// user Insert
func InsertUser(name string) error {
	user := models.User{
		Name: name,
	}
	return user.Insert(context.Background(), db, boil.Infer())
}

// user delete
func DeleteUser(n string) error {
	id, err := strconv.Atoi(n)
	if err != nil {
		return err
	}

	user := models.User{ID: id}
	_, err = user.Delete(context.Background(), db)
	return err
}

// get all tags
func GetAllTags() (models.TagSlice, error) {
	return models.Tags().All(context.Background(), db)
}

// get one tag
func GetOneTag(n string) (*models.Tag, error) {
	id, err := strconv.Atoi(n)
	if err != nil {
		return nil, err
	}

	return models.FindTag(context.Background(), db, id)
}

//tag insert
func InsertTag(name string) error {
	tag := models.Tag{
		Name: name,
	}
	return tag.Insert(context.Background(), db, boil.Infer())
}

func DeleteTag(n string) error {
	id, err := strconv.Atoi(n)
	if err != nil {
		return err
	}

	tag := models.Tag{ID: id}
	_, err = tag.Delete(context.Background(), db)
	return err
}
