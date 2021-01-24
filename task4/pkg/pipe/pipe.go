package pipe

import "context"

type Pipe struct {
	ctx     context.Context
	in, out chan string
	fn      func(string) string
}

func NewPipe(ctx context.Context, in, out chan string, fn func(string) string) *Pipe {
	return &Pipe{
		ctx: ctx,
		in:  in,
		out: out,
		fn:  fn,
	}
}

func (p *Pipe) Run() {
	for {
		select {
		case <-p.ctx.Done():
			return
		case val, ok := <-p.in:
			if !ok {
				return
			}
			p.out <- p.fn(val)
		}
	}
}
