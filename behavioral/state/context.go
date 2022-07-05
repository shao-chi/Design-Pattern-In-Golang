package main

type contextInterface interface {
	setState(state stateInterface)
	play()
	pause()
}

type context struct {
	state stateInterface
}

func (c *context) setState(state stateInterface) {
	(*c).state = state
}

func (c context) play() {
	c.state.play()
}

func (c context) pause() {
	c.state.pause()
}
