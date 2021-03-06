package main

import "fmt"

func showMap(m map[string]int) {
	fmt.Printf("%#v, len=%d\n", m, len(m))
}

func main() {
	// start block OMIT
	a := map[string]int{
		"hoge": 5,
		"fuga": 7,
	}
	showMap(a)

	val, ok := a["fuga"]
	fmt.Println(val, ok)

	delete(a, "fuga")
	showMap(a)

	val, ok = a["fuga"]
	fmt.Println(val, ok)
	// end block OMIT
}
