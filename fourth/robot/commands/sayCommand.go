package commands

import (
	"fmt"
	"homeworks/fourth/robot"
	"math/rand"
)

type SayCommand struct {
	vocabulary []string
}

func NewSayCommand() *SayCommand {
	return NewSayCommandWithVocabulary([]string{"Acila", "Tar", "Zoola"})
}

func NewSayCommandWithVocabulary(vocabulary []string) *SayCommand {
	return &SayCommand{vocabulary}
}

func (cmd *SayCommand) Execute(r *robot.Robot) {
	r.Greeting = "Hello, " + cmd.vocabulary[rand.Intn(len(cmd.vocabulary))]
	fmt.Println(r.Greeting)
}
