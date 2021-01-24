package bufferedchan

import "context"

type BufferedChan struct {
	ctx  context.Context
	in   chan string
	out  chan string
	size int
}

func NewChan(ctx context.Context, in chan string, bufferSize int) *BufferedChan {
	return &BufferedChan{
		ctx:  ctx,
		in:   in,
		out:  make(chan string, bufferSize),
		size: bufferSize,
	}
}

func (b *BufferedChan) Run() {
	defer close(b.out)
	for val := range b.in {
		select {
		case <-b.ctx.Done():
			return
		case b.out <- val:
		default: // scip values
		}
	}
}

func (b *BufferedChan) C() chan string {
	return b.out
}
