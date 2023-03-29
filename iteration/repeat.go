package iteration

const repetitionsQuantity = 5

func Repeat(char string) (repetitions string) {
	for i := 0; i < repetitionsQuantity; i++ {
		repetitions += char
	}
	return
}
