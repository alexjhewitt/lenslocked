package main

import (
	"html/template"
	"os"
	"time"
)

type User struct {
	Name string
	Bio  string
	Age  int
	Time time.Time
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	user := User{
		Name: "John Smith",
		Bio:  `<script>alert("Haha, you've been hacked!");</script>`,
		Age:  27,
		Time: time.Now(),
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
