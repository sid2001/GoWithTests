package iteration

const counter = 5

func Repeat(character string) (str string) {
	for i := 0; i < counter; i++ {
		str = str + character
	}
	return
}
