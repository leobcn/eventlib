package main

import (
	"log"
	"time"

	"github.com/EwanValentine/eventlib"
)

func main() {
	event := eventlib.NewEventLib()

	event.Subscribe("test", func(arg []byte) {
		log.Println(string(arg))
	})

	event.Publish("test", []byte("BIG TEST LOLOLOL"))

	event.Publish("another.test", []byte("FUCKING SHIT"))

	time.Sleep(time.Second * 5)

	event.Publish("test", []byte("FUCKKKKK"))

	event.Subscribe("another.test", func(arg []byte) {
		log.Println(string(arg))
	})

	time.Sleep(time.Second * 5)
}
