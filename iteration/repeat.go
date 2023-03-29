package iteration

func Repeat(char string, repeat int) (repetitions string) {
	for i := 0; i < repeat; i++ {
		repetitions += char
	}
	return
}
