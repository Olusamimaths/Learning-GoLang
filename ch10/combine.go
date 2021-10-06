package main

import (
	"context"
	"time"
)

/**
We have a function
that calls three web services. We send data to two of those services, and then take the
results of those two calls and send them to the third, returning the result. The entire
process must take less than 50 milliseconds, or an error is returned
**/

type Input struct {
	CIn
	C int
}

type AOut int
type BOut int
type COut struct {}
type CIn struct {
	A AOut
	B BOut
}

type processor struct {
	outA chan AOut
	outB chan BOut
	outC chan COut
	inC chan CIn
	errs chan error
}

func GatherAndProcess(ctx context.Context, data Input) (COut, error) {
	ctx, cancel := context.WithTimeout(ctx, 50 * time.Millisecond)
	defer cancel()

	p := processor{
		outA: make(chan AOut, 1),
		outB: make(chan BOut, 1),
		inC: make(chan CIn, 1),
		outC: make(chan COut, 1),
		errs: make(chan error, 2), // two errors can be potentially written - from A and B
	}
	p.launch(ctx, data)
	inputC, err := p.waitForAB(ctx)
	if err != nil {
		return COut{}, err
	}
	p.inC <- inputC
	out, err := p.waitForC(ctx)
	return out, err
}

func (p *processor) launch(ctx context.Context, data Input) {
	go func() {
		aOut, err := getResultA(ctx, data.A)
		if err != nil {
			p.errs <- err
			return
		}
		p.outA <- aOut
	}()

	go func() {
		bOut, err := getResultB(ctx, data.B)
		if err != nil {
			p.errs <- err
			return
		}
		p.outB <- bOut
	}()

	go func() {
		select {
			// if the context is cancelled first, do nothing
		case <- ctx.Done():
			return
			// if inputC is read first, then getResultC
		case inputC := <-p.inC:
			cOut, err := getResultC(ctx, inputC)
			if err != nil {
				p.errs <- err
				return
			}
			p.outC <- cOut
		}
	}()
}

func (p *processor) waitForAB(ctx context.Context) (CIn, error) {
	var inputC CIn
	count := 0

	for count < 2 {
		select {
		case a := <-p.outA:
			inputC.A = a
			count++
		case b := <-p.outB:
			inputC.B = b
			count++
		case err := <-p.errs:
			return CIn{}, err
		case <- ctx.Done():
			return CIn{}, ctx.Err()
		}
	}
	return inputC, nil
}

func (p *processor) waitForC(ctx context.Context) (COut, error) {
	select {
	case out := <-p.outC:
		return out, nil
	case err := <-p.errs:
		return COut{}, err
	case <- ctx.Done():
		return COut{}, ctx.Err()
	}
}

func getResultA(ctx context.Context, input AOut) (AOut,error) {
	return input, nil
}
func getResultB(ctx context.Context, input BOut) (BOut,error) {
	return input, nil
}

func getResultC(ctx context.Context, input CIn) (COut,error) {
	return COut{}, nil
}