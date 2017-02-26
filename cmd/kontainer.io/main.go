package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/ttdennis/kontainer.io/pkg/abstraction"
	"github.com/ttdennis/kontainer.io/pkg/user"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dbWrapper := abstraction.NewDB(db)

	userService, err := user.NewService(dbWrapper)
	if err != nil {
		panic(err)
	}
	_ = user.NewTransactionBasedService(userService)
}