package main

import (
	"log"
)

var globalObserver Observer

type Observer interface {
	Register(string)
	SetContext(func())
	Notify(string)
	StartRecording()
	StopRecording()
}

type GlobalObserver struct {
	currentContext func()
	dependencies   map[string][]func()
	recordingStage bool
}

func (g *GlobalObserver) StartRecording() {
	g.recordingStage = true
}

func (g *GlobalObserver) StopRecording() {
	g.recordingStage = false
}

func (g *GlobalObserver) Register(key string) {
	if g.recordingStage {
		g.dependencies[key] = append(g.dependencies[key], g.currentContext)
	}
}

func (g *GlobalObserver) SetContext(notifyFn func()) {
	if g.recordingStage {
		g.currentContext = notifyFn
	}
}

func (g *GlobalObserver) Notify(key string) {
	log.Printf("Got notified that key %s changed\n", key)

	for k, fns := range g.dependencies {
		if k == key {
			log.Printf("Notifying that key %s has changed", key)
			for _, fn := range fns {
				fn()
			}
		}
	}
}

func NewObserver() Observer {
	return &GlobalObserver{dependencies: make(map[string][]func(), 0)}
}

func InitObserver() {
	if globalObserver == nil {
		globalObserver = NewObserver()
		globalObserver.StartRecording()
	}
}

func init() {
	InitObserver()
}
