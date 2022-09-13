package iteration

func Repeat(character string, times uint) (repeated string) {
	var i uint
	for ; i < times; i++ {
		repeated += character
	}
	return
}
