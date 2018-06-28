package main

var globalObserver *Observer

type Observer interface {
}

func NewObserver() {
	return &Observer
}

func InitObserver() {
	if globalObserver == nil {
		globalObserver = NewObserver()
	}
}
