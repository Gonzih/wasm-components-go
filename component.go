package main

type Component interface {
	Render() error
	RenderToString() (string, error)
}
