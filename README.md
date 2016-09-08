# EventLib

A lightweight event library for Go, written in Go.

## How to use

```

import (
    
    // Core
    "log"
    "time"

    // Third party
    "github.com/EwanValentine/eventlib"
)

func main() {
    event := eventlib.NewEventLib()

    event.Subscribe("my.event", func(arg []byte) {
        log.Println(string(arg))    
    })

    time.Sleep(2 * time.Second)

    event.Publish("my.event", []byte("HELLOOOOO!"))

    time.Sleep(2 * time.Second)
}
```
