package iteracao

const quantidadeRepeticoes = 5

func Repetir(caractere string) string {
	var repeticao string

	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticao += caractere
	}

	return repeticao
}
