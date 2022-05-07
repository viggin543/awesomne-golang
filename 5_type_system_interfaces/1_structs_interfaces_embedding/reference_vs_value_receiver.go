package __structs_interfaces_embedding

import (
	"net"
	"os"
	"time"
)

func IpExample() {
	ip := net.ParseIP("192.0.2.1")
	// [idea tip] peek implementation
	// [idea tip] find usages of IP figure how to use it
	_, _ = ip.MarshalText() // <- why the MarshalText has a Value Receiver
	// when to use value receivers and when reference receivers ??
	// - value --> no mutation happens
	// - reference -> mutating type internal state
	// avoid value receivers on types / structs that occupy a large space since they will be copied on every method call !
}

func TimeExample() time.Time { // time.Time is a struct containing primitive data types
	now := time.Now()
	//When you think about time, you realize that any given point in time is not something that can change.
	//This is exactly how the standard library implements the Time type.
	future := now.Add(time.Hour) // peek Add implementation, why a value receiver ?
	return future
	// time is special, most structs should not be treated like a primitive and have value receivers
	// and instead hold some state that can be mutated and have methods with reference recievers
}

func FileExample() {
	file, _ := os.Open("/some/path") // Open factory method should return a Pointer ? or a Value ?
	file.Name()                      // what receiver should Name method have ?
	//file is non-primitive. Values of this type are unsafe to be copied
	// you can close/delete the wrong file.
	_ = file.Chdir() // what receiver should Chdir have ? even though it does not mutate anything
	// all "file" methods share the same reference.
	// in a sense a non-primitive struct that have methods that share the same ref, define a behavior
	// can be called an "Object"
}
