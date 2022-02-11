package main

import "fmt"

func main() {
	fmt.Println(Ola("mundo"))
}

func Ola(nome string) string {
	return "Olá, " + nome
}
