package fanin

import (
	"context"
)

type FanIn struct {
	ctx context.Context
	in  []chan string
	out chan string
}

func NewFanIn(ctx context.Context, channelsIn []chan string, out chan string) *FanIn {
	return &FanIn{
		ctx: ctx,
		in:  channelsIn,
		out: out,
	}
}

func (f *FanIn) Attach(ch ...chan string) {
	f.in = append(f.in, ch...)
}

func (f *FanIn) Detach(ch chan string) {
	for i, v := range f.in {
		if v == ch {
			f.in = append(f.in[:i], f.in[i+1:]...)
			break
		}
	}
}

func (f *FanIn) Run() {
	for {
		if len(f.in) == 0 {
			return
		}
		for _, c := range f.in {
			select {
			case <-f.ctx.Done():
				return
			case val, ok := <-c:
				if !ok {
					f.Detach(c)
					continue
				}
				f.out <- val
			default:
			}
		}
	}
}
