package eventlib

import (
	"testing"

	"github.com/EwanValentine/eventlib"
)

func TestPublish(t *testing.T) {

	payload := []byte("This is a test")

	event := eventlib.NewEventLib()
	event.Subscribe("test", func(arg []byte) {
		if string(arg) != string(payload) {
			t.Fail()
		}
	})

	event.Publish("test", payload)
}
