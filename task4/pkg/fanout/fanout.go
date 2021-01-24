package fanout

import (
	"context"
)

type FanOut struct {
	ctx context.Context
	in  chan string
	out []chan string
}

func NewFanOut(ctx context.Context, in chan string, channelsOut []chan string) *FanOut {
	return &FanOut{
		ctx: ctx,
		in:  in,
		out: channelsOut,
	}
}

func (f *FanOut) Attach(ch ...chan string) {
	f.out = append(f.out, ch...)
}

func (f *FanOut) Detach(ch chan string) {
	for i, v := range f.out {
		if v == ch {
			f.out = append(f.out[:i], f.out[i+1:]...)
			break
		}
	}
}

func (f *FanOut) Run() {
	for {
		if len(f.out) == 0 {
			return
		}
		select {
		case <-f.ctx.Done():
			return
		case val, ok := <-f.in:
			if !ok {
				return
			}
			for _, c := range f.out {
				c <- val
			}
		}
	}
}
