package robot

type Robot struct {
	Position
	Greeting string
}

func NewRobot() *Robot {
	robot := Robot{*NewPosition(), ""}
	return &robot
}

func (r *Robot) ProcessCommand(cmd Command) {
	cmd.Execute(r)
}
