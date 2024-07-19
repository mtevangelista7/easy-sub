package cli

import (
	"bufio"
	"easysub/internal/api"
	"easysub/internal/io"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func PerformSearch(path string) ([]string, error) {
	// pega uma lista com os arquivos de áudio
	fileNames, err := io.SearchAllFileName(path)

	if err != nil {
		return nil, err
	}

	return fileNames, nil
}

func AskForFileName(fileNames []string) (string, error) {
	// TODO: implementar a possibilidade de escolher vários de uma vez
	fmt.Println("Por favor, selecione um filme")

	// printa cada uma das opções para o usuário
	for index, fileName := range fileNames {
		fmt.Printf("[%d] - %s\n", index, fileName)
	}

	// bufio cria um novo Reader
	// Reader é uma abstração que pode ler dados de uma fonte
	// os.Stdin representa a entrada padrão, que aqui é usada para ler a entrada do usuário pelo teclado
	reader := bufio.NewReader(os.Stdin)

	// ReadString realiza a leitura da entrada de usuário até encontrar o delimitador no caso uma quebra de linha \n
	userInput, err := reader.ReadString('\n')

	if err != nil {
		log.Print("Não foi possível realizar a leitura da entrada do usuário")
		return "", err
	}

	// slice[:n] -> retorna um substring do começo da string até n
	// aqui pegamos a entrada do usuário até o último penúltimo elemento (antes da quebra de linha)
	userInput = userInput[:len(userInput)-1]

	// Atoi serve para converter a string para int
	userChoose, err := strconv.Atoi(userInput)

	if err != nil {
		log.Printf("Não foi possível realizar a conversão: %s", userInput)
		return "", err
	}

	// Pega a opção selecionada pelo usuario
	movieName := fileNames[userChoose]

	// retorna o nome do filme
	return movieName, nil
}

func GetMovieIdByName(fileName string) (string, error) {
	if fileName == "" {
		log.Print("Não foi possível buscar o ID pois, o nome do filme está em branco")
		return "", errors.New("o nome do filme não pode ficar em branco")
	}

	movie, err := api.GetMovieInfo(fileName)

	if err != nil {
		return "", err
	}

	if movie == nil {
		log.Fatal("Não foi possível prosseguir pois o filme não foi localizado!")
	}

	return movie.ImdbID, nil
}

func Login(username, password string) (string, error) {
	if username == "" || password == "" {
		log.Print("As informações de login não foram preenhcidas")
		return "", errors.New("as informações de login não foram preenchidas")
	}

	token, err := api.LoginOpenSub(username, password)

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetSub(userToken, imdbID string) error {
	sub, err := api.GetSubByImdbId(userToken, imdbID)

	if err != nil {
		return err
	}

	// TODO: Tratar isso aqui
	print(sub)
	return nil
}
