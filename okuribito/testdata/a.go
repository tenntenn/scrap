package main

var n *[100]int

func f() {
	n = new([100]int)
	n = nil
}

func main() {
	f()
	runtime.GC()
}
