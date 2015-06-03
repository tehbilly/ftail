// Package ftail is an opinionated file tailer similar to `tail -F somefile`
package ftail

import (
	"os"
	"sync"
	"time"

	"github.com/ActiveState/tail"
)

var (
	wg sync.WaitGroup
)

type FileLine struct {
	Time time.Time
	Line string
}

// Returns an unbuffered channel with lines from the file
// Begins at the end of the file (using os.SEEK_END)
func TailFile(filename string) (<-chan *FileLine, error) {
	t, err := tail.TailFile(filename, tail.Config{
		Follow:    true,
		ReOpen:    true,
		MustExist: false,
		Logger:    tail.DiscardingLogger,
		Location:  &tail.SeekInfo{Whence: os.SEEK_END},
	})

	if err != nil {
		return nil, err
	}

	lc := make(chan *FileLine)

	go func() {
		for l := range t.Lines {
			line := &FileLine{
				Time: time.Now(),
				Line: l.Text,
			}
			lc <- line
		}
	}()

	return lc, nil
}

func Cleanup() {
	tail.Cleanup()
}
