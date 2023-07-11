package main

func main() {
	const num = iota + 1 // iota is a special constant, it starts at 0 and increments by 1 for each item in the list
	println(num)
}
