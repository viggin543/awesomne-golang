package main

import (
	work "github.com/viggin543/awesomne-golang/8_concurenncy_patterns/patterns/3_work"
	"log"
	"sync"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	p := work.New(2)
	defer p.Shutdown()
	var wg sync.WaitGroup
	wg.Add(100 * len(names))
	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{name: name} // deep copies name
			go func() {
				p.Run(&np) // what can happen if this line would be p.Run(&namePrinter{name: name}) ?
				wg.Done()
			}()
		}
	}
	wg.Wait()
}
