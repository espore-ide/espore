package main

import (
	"espore/builder"
	"espore/config"
	"fmt"
	"log"
	"time"

	"github.com/radovskyb/watcher"
)

func watch(config *config.EsporeConfig) {
	w := watcher.New()
	w.SetMaxEvents(1)

	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Println(event) // Print the event's info.
				builder.Build(&config.Build)
				fmt.Println("done")
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch test_folder recursively for changes.
	if err := w.AddRecursive("firmware"); err != nil {
		log.Fatalln(err)
	}
	if err := w.AddRecursive("site"); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Watching for events...")
	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
