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

const monitoramentos = 3
const delay = 7

func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}

		// if comando == 1 {
		// 	fmt.Println("Monitorando...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo Logs...")
		// } else if comando == 3 {
		// 	fmt.Println("Saindo fora")
		// } else {
		// 	fmt.Println("Não conheço esse comando")
		// }
	}
}

func exibeIntroducao() {
	nome := "João Pedro"
	versao := 1.1
	fmt.Println("Olá, sr. ", nome)
	fmt.Println("este programa esta na versao", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	//fmt.Scanf("%d", &comando)
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println("")
	return comandoLido
}

func iniciarMonitoramento() {
	// IMPLEMENTAÇÂO 1 - ARRAY
	// var sites [4]string
	// sites[0] = "https://random-status-code.herokuapp.com"
	// sites[1] = "https://www.alura.com.br"
	// sites[2] = "https://www.gremio.net"
	// sites[3] = "https://www.google.com.br"

	// IMPLEMENTAÇÂO 2 - SLICE
	// sites := []string{"https://random-status-code.herokuapp.com", "https://www.alura.com.br", "https://www.gremio.net", "https://www.google.com.br"}

	// IMPLEMENTAÇÂO 3 - SLICE BUSCANDO DE ARQUIVO
	sites := leSitesDoArquivo()

	// for i := 0; i < len(sites); i++ {
	// }
	for i := 0; i < monitoramentos; i++ {
		for indice, site := range sites {
			fmt.Println("Testando site", indice, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu o erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Código HTTP:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	//arquivo, err := ioutil.ReadFile("sites.txt") -> Le tudo de uma vez
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu o erro:", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu o erro:", err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
