package bufferedchan

func NewChan(in chan string, bufferSize int) chan string {
	out := make(chan string, bufferSize)

	go func() {
		defer close(out)
		for val := range in {
			select {
			case out <- val:
			default: // scip values
			}
		}
	}()

	return out
}
