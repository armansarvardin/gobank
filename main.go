package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, name, lname, pw string) *Account {
	acc, err := NewAccount(name, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account => ", acc.Number)

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "arman", "sarvardin", "arman123")
}

func main() {

	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgressStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("Seeding the database")
		seedAccounts(store)
	}

	// seed stuff
	

	server := NewAPIServer(":3000", store)
	server.Run()
}