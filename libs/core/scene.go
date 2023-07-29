package core

type Scene interface {
	Start()
	Tick(input Input, timer Timer)
	Draw()
}
