package commands

import (
	"fmt"
	"homeworks/fourth/robot"
	"math/rand"
)

type UpdPositionCommand struct {
	maxPositionWidth  int
	maxPositionHeight int
}

func NewUpdPositionCommand() *UpdPositionCommand {
	return &UpdPositionCommand{
		robot.MaxPositionWidth,
		robot.MaxPositionHeight,
	}
}

func NewUpdPositionCommandWithMaxCoords(maxPositionWidth int, maxPositionHeight int) *UpdPositionCommand {
	return &UpdPositionCommand{
		maxPositionWidth,
		maxPositionHeight,
	}
}

func (cmd *UpdPositionCommand) Execute(r *robot.Robot) {
	r.Position.X = rand.Intn(cmd.maxPositionWidth)
	r.Position.Y = rand.Intn(cmd.maxPositionHeight)

	fmt.Println("New robot position: (", r.X, ";", r.Y, ")")
}
