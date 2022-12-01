package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()
	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()

		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Comando invalido digite novamente.")
			os.Exit(-1)
		}

	}

}

func exibeIntroducao() {
	nome := "Rafael"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versao:", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi:", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://www.globo.com"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("site Ok")
	} else {
		fmt.Println("error", resp.StatusCode)
	}
}
