package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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
	for {
		exibirMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLog()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("A opção escolhida não existe")
			os.Exit(-1)
		}
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

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()
	for i := 0; i < 5; i++ {
		time.Sleep(5 * time.Second)
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
	}

}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site", site, "está com problema. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
	fmt.Println("")
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}
	// fmt.Println(string(arquivo))
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		// fmt.Println(linha)

		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	fmt.Println(sites)
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	fmt.Println(arquivo)
	arquivo.Close()
}

func imprimeLog() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))

}
