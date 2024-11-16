package util

func Presentation() {
	println("Hello world...")
	println("This is a new Module")
}

func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func PrintNames(names ...string) {
	for _, name := range names {
		println(name)
	}
}
