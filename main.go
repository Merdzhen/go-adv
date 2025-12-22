package main

func main() {
	age := getAge()
	canDrink(age)

	// go run -gcflags '-m -l' main.go
}

func canDrink(age *int) bool {
	return *age >= 18
}

func getAge() *int {
	age := 18
	return &age
}
