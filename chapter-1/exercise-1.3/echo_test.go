package echo

import (
	"math/rand"
	"testing"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString() string {
	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func benchmarkEchoInefficiently(b *testing.B, argLength int) {
	args := []string{"./echo"}
	for i := 0; i < argLength; i++ {
		args = append(args, randString())
	}

	for i := 0; i < b.N; i++ {
		EchoInefficiently(args)
	}
}

func benchmarkEchoWithStringsJoin(b *testing.B, argLength int) {
	args := []string{"./echo"}
	for i := 0; i < argLength; i++ {
		args = append(args, randString())
	}

	for i := 0; i < b.N; i++ {
		EchoWithStringsJoin(args)
	}
}

func BenchmarkEchoInefficiently(b *testing.B) {
	benchmarkEchoInefficiently(b, 0)
}

func BenchmarkEchoInefficiently10(b *testing.B) {
	benchmarkEchoInefficiently(b, 10)
}

func BenchmarkEchoInefficiently100(b *testing.B) {
	benchmarkEchoInefficiently(b, 100)
}

func BenchmarkEchoWithStringsJoin(b *testing.B) {
	benchmarkEchoWithStringsJoin(b, 0)
}

func BenchmarkEchoWithStringsJoin10(b *testing.B) {
	benchmarkEchoWithStringsJoin(b, 10)
}

func BenchmarkEchoWithStringsJoin100(b *testing.B) {
	benchmarkEchoWithStringsJoin(b, 100)
}
