package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Criando um arquivo
	f, err := os.Create("arquivo.txt")

	// Se der erro envio um panic()
	if err != nil {
		panic(err)
	}

	// Grava um texto no arquivo e retorna o seu tamanho
	// tamanho, err := f.WriteString("Hello World")

	// Gravando bytes no meu arquivo, utilizo essa forma quando
	// eu não sei exatamente quais dados eu vou gravar no meu arquivo
	tamanho, err := f.Write([]byte("Gravando bytes no arquivo"))

	// Se a gravação der erro ele gera um erro
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes \n", tamanho)

	// Fechando a manipulação do arquivo
	f.Close()

	// Lendo o nosso arquivo
	arquivo2, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	// Eu converto para String, pois dentro dos arquivos são os
	// slices de bytes que ficam gravados
	fmt.Println(string(arquivo2))

	// Leitura de pouco em pocuo abrindo o arquivo
	arquivo, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// Utilizando o pacote bufio
	// Criando um reader (lê o conteúdo aos poucos)
	reader := bufio.NewReader(arquivo)
	// O buffer vai ler o arquivo de 10 em 10 bytes
	buffer := make([]byte, 10)
	for {
		// Faz a leitura do arquivo baseado no buffer
		n, err := reader.Read(buffer)
		if err != nil {
			// Se der algum erro, ele para de ler o arquivo
			break
		}
		// Converte o buffer para string
		// :n é a posição onde ele está fazendo a leitura
		fmt.Println(string(buffer[:n]))
	}
}
