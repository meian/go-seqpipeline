package seq_test

import (
	"testing"

	"github.com/meian/go-seqpipeline/seq"
)

func TestMapNilInput(t *testing.T) {
	iterator := seq.Map(func(int) int { t.Fatal("Map callback called"); return 0 })(nil)
	for round := 0; round < 2; round++ {
		for range iterator {
			t.Fatal("nil input yielded a value")
		}
	}
}

func TestMap2NilInput(t *testing.T) {
	iterator := seq.Map2(func(int, string) (int, string) { t.Fatal("Map2 callback called"); return 0, "" })(nil)
	for round := 0; round < 2; round++ {
		for range iterator {
			t.Fatal("nil input yielded a value")
		}
	}
}
