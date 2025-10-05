package main

func main() {
	output(`func output(s string) {
	println("package main")
	println("")
	println("func main() {")
	println("\toutput(\x60" + s + "\x60)")
	println("}")
	println("")
	println(s)
}
	`)
}

func output(s string) {
	println("package main")
	println("")
	println("func main() {")
	println("\toutput(\x60" + s + "\x60)")
	println("}")
	println("")
	println(s)
}
