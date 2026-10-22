package seq_test

import (
	"context"
	"testing"
	"testing/synctest"

	"github.com/meian/go-seqpipeline/op"
	"github.com/meian/go-seqpipeline/seq"
)

// 同じ context のキャンセルで上流の受信待ちを解除し、パイプラインの出力を閉じる。
func TestChannelPipelineCancelReceive(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		input := make(chan int)
		pipeline := seq.FromChannel(ctx, input).
			Pipe(op.Map(func(v int) int { return v * 10 })).
			Pipe(op.Take(2).Operator)
		output := seq.ToChannel(ctx, pipeline)
		input <- 3
		assertEqual(t, <-output, 30)
		synctest.Wait() // 上流の受信待ちも同じ context で解除する。
		cancel()
		if _, ok := <-output; ok {
			t.Fatal("output must close after cancellation")
		}
	})
}
