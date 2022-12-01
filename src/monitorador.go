package main

import "fmt"

func main() {
	nome := "Rafael"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versao:", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi:", comando)

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo logs...")
	} else if comando == 3 {
		fmt.Println("Saindo...")
	} else {
		fmt.Println("Comando invalido digite novamente.")
	}
}
