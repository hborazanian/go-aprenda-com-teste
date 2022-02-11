package iteracao

import "testing"

func TestRepetir(t *testing.T) {
	repeticao := Repetir("a")
	esperado := "aaaaa"

	if repeticao != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticao)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a")
	}
}
