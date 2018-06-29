package main

import (
	"fmt"
	"log"
	"time"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error: %s\r\n", err)
		panic(err)
	}
}

func main() {
	c := make(chan struct{}, 0)

	store := NewStore()

	store.Set("label", "hello from the store")
	store.Set("count", 0)

	go func() {
		for {
			store.Update("count", func(i interface{}) interface{} {
				return i.(int) + 1
			})
			time.Sleep(time.Second * 5)
		}
	}()

	cmp, err := NewComponent("helloTemplate", "root", func(cmp *GenericComponent) error {
		cmp.props = struct {
			ID    string
			Label string
			Count int
		}{
			ID:    cmp.componentID,
			Label: store.GetString("label"),
			Count: store.GetInt("count"),
		}

		return nil
	})

	checkErr(err)
	checkErr(cmp.Render())

	globalObserver.StopRecording()

	go func() {
		for {
			checkErr(cmp.Render())
			time.Sleep(time.Millisecond * 100)
		}
	}()

	log.Println(globalObserver)

	<-c
}
