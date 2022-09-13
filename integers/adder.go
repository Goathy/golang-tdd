package integers

// Add takes multiple integers and returns the sum of them.
func Add(numbers ...int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}
