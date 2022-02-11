package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	repeticao := Repetir("a", 5)
	esperado := "aaaaa"

	if repeticao != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticao)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 5)
	}
}

func ExampleRepetir() {
	repeticao := Repetir("a", 5)
	fmt.Println(repeticao)
	// Output: aaaaa
}
