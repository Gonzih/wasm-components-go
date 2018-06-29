package main

import (
	"log"
)

var globalObserver Observer

type Observer interface {
	Register(string)
	SetContext(string)
	Notify(string)
}

type GlobalObserver struct {
	currentContext string
	dependencies   map[string][]string
}

func (g *GlobalObserver) Register(key string) {
	g.dependencies[key] = append(g.dependencies[key], g.currentContext)
}

func (g *GlobalObserver) SetContext(id string) {
	g.currentContext = id
}

func (g *GlobalObserver) Notify(key string) {
	log.Printf("Got notified that key %s changed\n", key)

	for k, ids := range g.dependencies {
		if k == key {
			for _, id := range ids {
				log.Printf("Component %s should be notified\n", id)
			}
		}
	}
}

func NewObserver() Observer {
	return &GlobalObserver{dependencies: make(map[string][]string, 0)}
}

func InitObserver() {
	if globalObserver == nil {
		globalObserver = NewObserver()
	}
}

func init() {
	InitObserver()
}
