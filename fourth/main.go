package main

import (
	"fmt"
	"homeworks/fourth/robot"
	"homeworks/fourth/robot/commands"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	robotPtr := robot.NewRobot()

	commandsSlice := make([]robot.Command, 0)

	commandsSlice = append(commandsSlice, commands.NewSayCommand())
	commandsSlice = append(commandsSlice, commands.NewUpdPositionCommand())
	commandsSlice = append(commandsSlice, commands.NewSayCommandWithVocabulary([]string{"Geralt", "Ciri", "Yennifer"}))
	commandsSlice = append(commandsSlice, commands.NewUpdPositionCommandWithMaxCoords(100, 100))

	for i := 0; i < 20; i++ {
		fmt.Print(i, ": ")

		commandNumber := rand.Intn(len(commandsSlice))
		robotPtr.ProcessCommand(commandsSlice[commandNumber])
	}
}
