package main

import "testing"

func BenchmarkStringPlus100(b *testing.B) {
	p := initStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stringPlus(p)
	}
}

func BenchmarkStringFmt100(b *testing.B) {
	p := initStringi(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stringFmt(p)
	}
}

func BenchmarkStringJoin100(b *testing.B) {
	p := initStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stringJoin(p)
	}
}

func BenchmarkStringBuffer100(b *testing.B) {
	p := initStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stringBuffer(p)
	}
}

func BenchmarkStringBuilder100(b *testing.B) {
	p := initStrings(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stringBuilder(p)
	}
}
