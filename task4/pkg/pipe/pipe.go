package pipe

func NewPipe(ch1, ch2 chan string, fn func(string) string) {
	for val := range ch1 {
		ch2 <- fn(val)
	}
}
