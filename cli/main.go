package main

import (
	"fmt"

	"github.com/sayden/flatten"
)

type User struct {
	City string
	PersonalData
	ContactMethods map[string]string
	ExtraInfo      Extra
}

type Extra struct {
	FavoriteSport string
}

type PersonalData struct {
	Name    string
	Surname string
	Age     int
	Weight  float64
}

func main() {
	user := User{
		City: "Madrid",
		ContactMethods: map[string]string{
			"Phone": "1234",
			"email": "aaaa@aaaa.com",
		},
		ExtraInfo: Extra{
			FavoriteSport: "Road Cycling",
		},
		PersonalData: PersonalData{
			Name:    "sayden",
			Surname: "A surname",
			Age:     32,
			Weight:  83.5,
		},
	}

	si := flatten.NewIterator(user)
	k, v, finished := si.Next()
	for !finished {
		fmt.Printf("%s -> %v\n", k, v)
		k, v, finished = si.Next()
	}

	m := map[string]string{
		"Hello": "World",
	}

	mi := flatten.NewIterator(m)
	k, b, finished := mi.NextBytes()
	for !finished {
		fmt.Printf("%s -> %s\n", k, string(b))
		k, b, finished = mi.NextBytes()
	}
}
