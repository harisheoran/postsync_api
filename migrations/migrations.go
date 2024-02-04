package main

import (
	"github.com/harisheoran/postsync/inits"
	"github.com/harisheoran/postsync/models"
)

func init() {
	inits.LoadEnv()
	inits.DBinit()
}

// It automatically creates or updates the
// database tables to match the structure of the provided model.
// the AutoMigrate function checks the existing database schema
// against the structure of the Post model.
// if the Post table doesn't exist, it creates the table with the required columns.
// If the table already exists but has differences (e.g., missing or modified columns),
// it updates the table structure to match the Post model.
// This ensures that your database schema is in sync with your application's code, making it
// easy to evolve the database structure as your application evolves.

func main() {
	inits.DB.AutoMigrate(&models.Post{})
}
