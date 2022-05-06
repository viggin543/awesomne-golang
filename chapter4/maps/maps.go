package maps

import "fmt"

func DeclaringAMap() (map[string]string, map[int][]string) {
	dict := make(map[string]string)                                 // no reason to use `new` since map is a reference type
	dict = map[string]string{"Red": "#da1337", "Orange": "#e95a22"} // map literal
	dict2 := map[int][]string{}                                     // is dict2 nil ?
	dict2[0] = []string{"tak"}
	return dict, dict2
}

func NilMap() {
	var colors map[string]string
	colors["Red"] = "#da1337" // crush ~!@$%!@#R
}

func PassingMapsToFunctions() {
	m := map[string]int{"foo": 1}
	bar(m)
	fmt.Println(m["foo"])

}

func bar(m map[string]int) { // a map is also a reference type, zero value is nil and it's always passed by reference to methods
	m["foo"] = 2
}

func isExisting() {
	_map := map[string]func(){"jump": func() {
		print("hop") // functions are first class citizens
	}}
	value, exists := _map["jump"] // common pattern
	if exists {
		value()
	}

	f := _map["boom"]
	println(f) // f is nil, why ? what would it be if _map was map[string]string
}

func iteratingOverAMap() {
	for key, value := range map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	} {
		fmt.Printf("Key: %s  Value: %s\n", key, value)
	}
}

func deletingFromAMap() {
	m := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	delete(m, "DarkGray") // new and delete are juts functions, not fancy operators ( like js )
}
