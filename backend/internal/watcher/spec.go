package watcher

import (
	"log"

	"github.com/fsnotify/fsnotify"

	"github.com/agnivo988/MockForge/backend/internal/parser"
	"github.com/agnivo988/MockForge/backend/internal/store"
)

func Watch(env, path string) {
	watcher, _ := fsnotify.NewWatcher()
	_ = watcher.Add(path)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					spec, err := parser.LoadSpec(path)
					if err == nil {
						store.Manager.Set(env, spec)
						log.Println("Reloaded spec:", env)
					} else {
						log.Println("Invalid spec reload:", err)
					}
				}
			}
		}
	}()
}
