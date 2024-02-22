package main

import (
	"context"

	concurrencypractice "github.com/leonardodelira/100-go-mistakes/8-concurrency-practice"
)

func main() {
	points := []int{1, 2, 4, 5, 3}
	ctx := context.Background()
	concurrencypractice.Mistake73_UseErrGroup(ctx, points)
}
