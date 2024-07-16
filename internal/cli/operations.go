package cli

import (
	"bufio"
	"easysub/internal/io"
	"fmt"
	"os"
	"strconv"
)

func PerformSearch(path string) []string {
	// pega uma lista com os arquivos de áudio
	fileNames := io.SearchAllFileName(path)
	return fileNames
}

func AskForFileName(fileNames []string) string {
	fmt.Println("Please, choose one movie")

	// printa cada uma das opções para o usuário
	for index, fileName := range fileNames {
		fmt.Printf("[%d] - %s\n", index, fileName)
	}

	// aguarda a reposta
	reader := bufio.NewReader(os.Stdin)

	// realiza a leitura
	userInput, _ := reader.ReadString('\n')

	// remove o \n (todo não peguei muito o do readstring)
	userInput = userInput[:len(userInput)-1]

	// converte a opção selecionada pela o usuario
	userChoose, err := strconv.Atoi(userInput)

	if err != nil {
		fmt.Println(err)
	}

	// Pega a opção selecionada pelo usuario
	movieName := fileNames[userChoose]

	// retorna o nome do filme
	return movieName
}
