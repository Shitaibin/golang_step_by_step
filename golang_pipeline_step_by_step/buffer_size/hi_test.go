package main

import "testing"

func BenchmarkPipelineFan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PipelineFan(0)
	}
}

func BenchmarkPipelineSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PipelineSimple()
	}
}

func BenchmarkPipelineFanBuffered_0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PipelineFan(0)
	}
}

func BenchmarkPipelineFanBuffered_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PipelineFan(1)
	}
}

func BenchmarkPipelineFanBuffered_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PipelineFan(2)
	}
}

func BenchmarkPipelineFanBuffered_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PipelineFan(3)
	}
}

func BenchmarkPipelineFanBuffered_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PipelineFan(4)
	}
}

func BenchmarkPipelineFanBuffered_5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PipelineFan(5)
	}
}

func BenchmarkPipelineFanBuffered_6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PipelineFan(6)
	}
}

func BenchmarkPipelineFanBuffered_7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PipelineFan(7)
	}
}

func BenchmarkPipelineFanBuffered_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PipelineFan(8)
	}
}

func BenchmarkPipelineFanBuffered_9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PipelineFan(9)
	}
}

func BenchmarkPipelineFanBuffered_n10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PipelineFan(10)
	}
}
