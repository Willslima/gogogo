package main

import (
	"fmt"
	"os"
)

// if comando == 1 {
// 	fmt.Println("Monitorando...")
// } else if comando == 2 {
// 	fmt.Println("Exibindo logs...")
// } else if comando == 0 {
// 	fmt.Println("Saindo do programa")
// } else {
// 	fmt.Println("A opção escolhida não existe")
// }

func main() {
	exibeIntroducao()
	exibirMenu()
	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("A opção escolhida não existe")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "Willian"
	versao := 1.1

	fmt.Println("Olá", nome)
	fmt.Println("Este programa está na versão: ", versao)
}

func exibirMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir os logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scanf("%d", &comandoLido)
	fmt.Println("A opção escolhida foi", comandoLido)
	return comandoLido
}
