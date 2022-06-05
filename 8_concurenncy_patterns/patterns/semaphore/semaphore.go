// This sample program demonstrates how to implement a semaphore using
// channels that can allow multiple reads but a single write.
//
// It uses the generator pattern to create channels and goroutines.
//
// Multiple reader/writers can be created and run concurrently. Then after
// a timeout period, the program shutdowns cleanly.
//
// http://www.golangpatterns.info/concurrency/semaphores
package main

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type (
	semaphore chan struct{}
)

type (
	readerWriter struct {
		name string

		write sync.WaitGroup

		readerControl semaphore

		shutdown chan struct{}

		reportShutdown sync.WaitGroup

		// maxReads defined the maximum number of reads that can occur at a time.
		maxReads int

		// maxReaders defines the number of goroutines launched to perform read operations.
		maxReaders int

		// currentReads keeps a safe count of the current number of reads occurring
		// at any give time.
		currentReads int32
	}
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	log.Println("Starting Process")
	first := start("First", 3, 6)
	second := start("Second", 2, 2)
	time.Sleep(2 * time.Second)
	shutdown(first, second)
	log.Println("Process Ended")
	return
}

// start uses the generator pattern to create the readerWriter value. It launches
// goroutines to process the work, returning the created ReaderWriter value.
func start(name string, maxReads int, maxReaders int) *readerWriter {
	rw := readerWriter{
		name:          name,
		shutdown:      make(chan struct{}),
		maxReads:      maxReads,
		maxReaders:    maxReaders,
		readerControl: make(semaphore, maxReads),
	}

	rw.reportShutdown.Add(maxReaders)
	for goroutine := 0; goroutine < maxReaders; goroutine++ {
		go rw.reader(goroutine)
	}

	rw.reportShutdown.Add(1)
	go rw.writer()

	return &rw
}

func shutdown(readerWriters ...*readerWriter) {
	var waitShutdown sync.WaitGroup
	waitShutdown.Add(len(readerWriters))

	for _, readerWriter := range readerWriters {
		go readerWriter.stop(&waitShutdown)
	}
	waitShutdown.Wait()
}

func (rw *readerWriter) stop(waitShutdown *sync.WaitGroup) {
	defer waitShutdown.Done()

	log.Printf("%s\t: #####> Stop", rw.name)

	// Close the channel which will causes all the goroutines waiting on
	// this channel to receive the notification to shutdown.
	close(rw.shutdown)

	// Wait for all the goroutine to call Done on the waitgroup we
	// are waiting on.
	rw.reportShutdown.Wait()

	log.Printf("%s\t: #####> Stopped", rw.name)
}

func (rw *readerWriter) reader(reader int) {
	defer rw.reportShutdown.Done()

	for {
		select {
		case <-rw.shutdown:
			log.Printf("%s\t: #> Reader Shutdown", rw.name)
			return

		default:
			rw.performRead(reader)
		}
	}
}

func (rw *readerWriter) performRead(reader int) {
	rw.ReadLock(reader)

	// Safely increment the current reads counter
	count := atomic.AddInt32(&rw.currentReads, 1)

	// Simulate some reading work
	log.Printf("%s\t: [%d] Start\t- [%d] Reads\n", rw.name, reader, count)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	// Safely decrement the current reads counter
	count = atomic.AddInt32(&rw.currentReads, -1)
	log.Printf("%s\t: [%d] Finish\t- [%d] Reads\n", rw.name, reader, count)

	// Release the read lock for this critical section.
	rw.ReadUnlock(reader)
}

// writer is a goroutine that listens on the shutdown channel and
// performs writes until the channel is signaled.
func (rw *readerWriter) writer() {
	defer rw.reportShutdown.Done()

	for {
		select {
		case <-rw.shutdown:
			log.Printf("%s\t: #> Writer Shutdown", rw.name)
			return
		default:
			rw.performWrite()
		}
	}
}

// performWrite performs the actual write work.
func (rw *readerWriter) performWrite() {
	// Wait a random number of milliseconds before we write again.
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	log.Printf("%s\t: *****> Writing Pending\n", rw.name)

	// Get a write lock for this critical section.
	rw.WriteLock()

	// Simulate some writing work.
	log.Printf("%s\t: *****> Writing Start", rw.name)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("%s\t: *****> Writing Finish", rw.name)

	// Release the write lock for this critical section.
	rw.WriteUnlock()
}

// ReadLock guarantees only the maximum number of goroutines can read at a time.
func (rw *readerWriter) ReadLock(reader int) {

	rw.write.Wait()

	rw.readerControl.Acquire(1)
}

func (rw *readerWriter) ReadUnlock(reader int) {
	rw.readerControl.Release(1)
}

func (rw *readerWriter) WriteLock() {
	rw.write.Add(1)
	rw.readerControl.Acquire(rw.maxReads)
}

func (rw *readerWriter) WriteUnlock() {
	rw.readerControl.Release(rw.maxReads)
	rw.write.Done()
}

func (s semaphore) Acquire(buffers int) {
	var e struct{}
	for buffer := 0; buffer < buffers; buffer++ {
		s <- e
	}
}

func (s semaphore) Release(buffers int) {
	for buffer := 0; buffer < buffers; buffer++ {
		<-s
	}
}
