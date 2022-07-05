package main

import (
	"fmt"
)

type stateInterface interface {
	play()
	pause()
}

type readyState struct {
	context *context
}

type playingState struct {
	context *context
}

type pausedState struct {
	context *context
}

type finishedState struct {
	context *context
}

func (state readyState) play() {
	fmt.Println("Start playing.")
	state.context.setState(playingState{context: state.context})
}

func (state readyState) pause() {
	fmt.Println("It's ready state. You can't pause.")
}

func (state playingState) play() {
	fmt.Println("It's already playing.")
}

func (state playingState) pause() {
	fmt.Println("Paused.")
	state.context.setState(pausedState{context: state.context})
}

func (state pausedState) play() {
	fmt.Println("Continue playing.")
	state.context.setState(playingState{context: state.context})
}

func (state pausedState) pause() {
	fmt.Println("It's already paused.")
}

func (state finishedState) play() {
	fmt.Println("Replay.")
	state.context.setState(playingState{context: state.context})
}

func (state finishedState) pause() {
	fmt.Println("It's finished. You can't pause.")
}
