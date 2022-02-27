package main

import (
	"fmt"
	"github.com/zzlalani/go-practice/internal/repository"
)

const ConnectionString = "host=localhost user=postgres password=testtest dbname=go-practice port=5432 sslmode=disable"

func main() {
	db, err := repository.Connect(ConnectionString)
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepo(db)

	// createUsers(userRepo)

	// read(userRepo)

	delete(userRepo)
}


func createUsers(userRepo repository.UserRepository) {
	user := repository.User{
		Username: "Ravi",
		Password: "123456",
	}

	user2 := repository.User{
		Username: "Shadood",
		Password: "123456",
	}

	id1, err := userRepo.Insert(&user)
	if err != nil {
		panic("cannot create ravi")
	}

	fmt.Println("id1", id1)

	id2, err := userRepo.Insert(&user2);
	if err != nil {
		panic("cannot create shadood")
	}

	fmt.Println("id2", id2)
}

func read(userRepo repository.UserRepository) {
	user2, err := userRepo.Read(5)
	if err != nil {
		panic(err)
	}
	user2.Password = "1234567"
	if err := userRepo.Update(user2, 5); err != nil {
		panic(err)
	}
}

func delete(userRepo repository.UserRepository) {
	user2, err := userRepo.Read(2)
	if err != nil {
		panic(err)
	}
	if err := userRepo.Delete(user2); err != nil {
		panic(err)
	}
}
