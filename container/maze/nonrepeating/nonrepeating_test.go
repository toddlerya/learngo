package main

import (
	"testing"
)

func TestSubStr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbb", 1},
		{"abcabcabcd", 4},

		// Chinese Support
		{"这里是慕课网", 6},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		acutal := lengthOfNonRepeatingSubStr(tt.s)
		if acutal != tt.ans {
			t.Errorf("got %d for input %s; expected %d", acutal, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubStr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		s = s + s
	}
	b.Logf("len(s) = %d", len(s))
	ans := 8

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		acutal := lengthOfNonRepeatingSubStr(s)
		if acutal != ans {
			b.Errorf("go %d for input %s; expected %d", acutal, s, ans)
		}
	}
}
