// pacote principal (main) é o pacote que contém o arquivo executável
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	ColorBlack  = "\u001b[30m"
	ColorRed    = "\u001b[31m"
	ColorGreen  = "\u001b[32m"
	ColorYellow = "\u001b[33m"
	ColorBlue   = "\u001b[34m"
	ColorReset  = "\u001b[0m"
)

const seconds = 2
const version = "1.19"
const name = "Lucas"

func main() {
	printPersonalInfos()
	for {
		comando := verifyCommand()

		switch comando {
		case 1:
			InitMonitoramento()
		case 2:
			readLogs()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Número inválido!")
			os.Exit(-1)
		}
	}
}

func printPersonalInfos() {
	fmt.Println("Olá,", name)
	fmt.Println("A versão do seu programa é ", version)
}

func verifyCommand() int {
	fmt.Println("==================================")
	colorize(ColorYellow, "O que você deseja fazer?")
	fmt.Println("==================================")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir os Logs")
	fmt.Println("0- Sair do Programa")
	fmt.Println("==================================")

	var comando int
	fmt.Scan(&comando)

	return comando
}

func InitMonitoramento() {

	sites := readSites()
	for _, site := range sites {

		resp, error := http.Get(site)

		if error != nil {
			fmt.Println("Aconteceu um erro", error)
		}

		var printer = "Monitorando " + site
		colorize(ColorBlue, printer)
		time.Sleep(seconds * time.Second)
		fmt.Print(" ")
		if resp.StatusCode == 200 {
			var printer = "Site: " + site + " foi carregado com sucesso!"
			colorize(ColorGreen, printer)
			logRegistration(site, resp.StatusCode)
		} else {
			var printer = "Site: " + site + " está com problemas. Status Code:" + fmt.Sprint(resp.StatusCode)
			colorize(ColorRed, printer)
			logRegistration(site, resp.StatusCode)
		}
		fmt.Println(" ")
	}
}

func colorize(color string, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func readSites() []string {
	var sites []string
	file, error := os.Open("sites.txt")
	if error != nil {
		fmt.Println("Ocorreu um erro:", error)
	}

	leitor := bufio.NewReader(file)
	for {
		linha, error := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)
		sites = append(sites, linha)
		if error == io.EOF {
			break
		} else if error != nil {
			fmt.Println("Ocorreu um erro:", error)
		}
	}
	file.Close()
	return sites
}

func logRegistration(site string, status int) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	log := fmt.Sprint(time.Now().Local().Format("02/01/2006 15:04:05"))+ " " + site + " " + fmt.Sprint(status) + "\n"
	file.WriteString(log)

	file.Close()
}

func readLogs() {
	file, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	binario := bufio.NewReader(file)
	for {
		line, err := binario.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Ocorreu um erro", err)
		}
	}
}
