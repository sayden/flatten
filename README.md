# flatten

Flatten is a package to flatten a Struct, a Map or a mix of them to a flattened *key-value* 2D matrix.
 
Useful to convert structs to 2D Matrices prior storing them in a *row column like* database like MySQL or Cassandra.

## Limitations

* Does not support arrays

## Example use

```go
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
			"email": "mariocaster@gmail.com",
		},
		ExtraInfo: Extra{
			FavoriteSport: "Road Cycling",
		},
		PersonalData: PersonalData{
			Name:    "Mario",
			Surname: "Castro",
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
	k, v, finished = mi.Next()
	for !finished {
		fmt.Printf("%s -> %v\n", k, v)
		k, v, finished = mi.Next()
	}
}
```

Outputs

```bash
City -> Madrid
PersonalData.Name -> Sayden
PersonalData.Surname -> A surname
PersonalData.Age -> 32
PersonalData.Weight -> 83.5
ContactMethods.email -> aaaa@aaaa.com
ContactMethods.Phone -> 1234
ExtraInfo.FavoriteSport -> Road Cycling
Hello -> World
```