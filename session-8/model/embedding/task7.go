package embedding

import "fmt"

type Mover interface {
	Move(m string)
}
type Sayer interface {
	Say(m string)
}
type Robot interface {
	Mover
	Sayer
}
type AutoBot struct {
	message string
}

func (a AutoBot) Move(m string) {
	fmt.Println(m)
}

func (a AutoBot) Say(m string) {
	fmt.Println(m)
}

func Task7() {
	fmt.Println("Task 7  ****************")
	bot := AutoBot{}
	var r Robot = bot
	r.Move("Moving forward")
	r.Say("I am a robot")
}
