package main

import (
	"log"
)

var globalObserver Observer

type Observer interface {
	Register(string)
	SetContext(string)
	Notify(string)
	StartRecording()
	StopRecording()
}

type GlobalObserver struct {
	currentContext string
	dependencies   map[string][]string
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

func (g *GlobalObserver) SetContext(id string) {
	if g.recordingStage {
		g.currentContext = id
	}
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
		globalObserver.StartRecording()
	}
}

func init() {
	InitObserver()
}
