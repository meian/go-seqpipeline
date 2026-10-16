package seq_test

import (
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestFilterNilInput(t *testing.T) {
	iterator := seq.Filter(func(int) bool { t.Fatal("Filter callback called"); return true })(nil)
	for round := 0; round < 2; round++ {
		for range iterator {
			t.Fatal("nil input yielded a value")
		}
	}
}

func TestFilter2NilInput(t *testing.T) {
	iterator := seq.Filter2(func(int, string) bool { t.Fatal("Filter2 callback called"); return true })(nil)
	for round := 0; round < 2; round++ {
		for range iterator {
			t.Fatal("nil input yielded a value")
		}
	}
}
