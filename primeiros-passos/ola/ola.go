package main

import "fmt"

const prefixoOlaPortugues = "Olá, "

func main() {
	fmt.Println(Ola("mundo"))
}

func Ola(nome string) string {
	if nome == "" {
		nome = "mundo"
	}

	return prefixoOlaPortugues + nome
}
