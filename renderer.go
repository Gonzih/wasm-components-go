package main

type Renderer interface {
	Render(*El) error
}
