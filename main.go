package main

import (
	"flag"
	"log"
)

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "dave", "sarasa", "huner88888")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()
	store, err := NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		// seed stuff
		seedAccounts(store)
	}

	server := NewAPIServer(":3030", store)
	server.Run()
}
