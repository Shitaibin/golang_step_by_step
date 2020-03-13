package strconv_test

import (
	"strconv"
	"testing"
)

func BenchmarkItoaString(b *testing.B) {
	// 只有待转换数字在0~10才是安全的
	for i := 0; i < b.N; i++ {
		_ = string(1 + '0')
	}
}

func BenchmarkItoaStrConv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(1)
	}
}

func TestItoa(t *testing.T) {
	// 只有待转换数字在0~10才是安全的
	for i := 0; i < 1000; i++ {
		x := string(i + '0')
		y := strconv.Itoa(i)
		if x != y {
			t.Fatalf("x: %v, y: %v\n", x, y)
		}
	}
}
