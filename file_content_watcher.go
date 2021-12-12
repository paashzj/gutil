package gutil

import (
	"io/ioutil"
	"sync"
	"time"
)

type FileContentWatcher struct {
	filename string
	Content  []byte
	RwLock   sync.RWMutex
}

func (f *FileContentWatcher) init(filename string) {
	f.filename = filename
	go func() {
		for {
			bytes, err := ioutil.ReadFile(filename)
			if err == nil {
				f.RwLock.Lock()
				f.Content = bytes
				f.RwLock.Unlock()
			}
			time.Sleep(5 * time.Second)
		}
	}()
}
