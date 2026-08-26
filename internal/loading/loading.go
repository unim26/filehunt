package loading

import (
	"fmt"
	"time"
)

type Loading struct {
	hide chan bool
}

// show loading
func Show(msg string) *Loading {
	l := &Loading{
		hide: make(chan bool),
	}

	go func() {
		chars := []string{"|", "/", "-", "\\"}
		i := 0
		for {
			select {
			case <-l.hide:
				return
			default:
				fmt.Printf("\r%s %s", msg, chars[i])
				i = (i + 1) % len(chars)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	return l
}

// hide
func (L *Loading) Hide() {
	L.hide <- true
	fmt.Printf("\r\033[K")
}
