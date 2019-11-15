package main

import "testing"

// sliceの容量を指定しない場合
func BenchmarkInitSliceVariable(b *testing.B) {
	var target []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target = append(target, i)
	}
}

// sliceでmakeで容量を確保しているがlength指定している場合(値が0で初期化されている場合)
func BenchmarkSliceCapacityNo(b *testing.B) {
	var target = make([]int, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target = append(target, i)
	}
}

// sliceのmake時にc容量を設定する場合
func BenchmarkSliceCapacityYes(b *testing.B) {
	var target = make([]int, 0, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target = append(target, i)
	}
}