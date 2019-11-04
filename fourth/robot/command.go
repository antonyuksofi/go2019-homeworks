package robot

type Command interface {
	Execute(r *Robot)
}
