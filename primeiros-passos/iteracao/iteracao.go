package iteracao

func Repetir(caractere string, repeticoes int) string {
	var repeticao string

	for i := 0; i < repeticoes; i++ {
		repeticao += caractere
	}

	return repeticao
}
