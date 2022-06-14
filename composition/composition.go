package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	Id             int
	Name, Location string
}

//method
func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

type Player struct {
	User   //user will contain all the desired attributes
	GameId int
	Date   time.Time
}

type Job struct {
	Command string
	Logger  *log.Logger
}

func main() {
	p := Player{} //initializing
	p.Id = 42
	p.Name = "Matt"
	p.Location = "LA"
	p.GameId = 90404
	fmt.Printf("%+v\n", p)
	p.Id = 11
	fmt.Printf("%+v\n", p)

	//other way
	p2 := Player{User{Id: 43, Name: "Matt", Location: "LA"}, 90404, time.Now()}
	fmt.Printf("%+v\n", p2)
	fmt.Println(p2.Greetings())
	fmt.
	job := &Job{"demo", log.New(os.Stdout, "Job: ", log.Ldate)}
	job.Logger.Print("test")
}
