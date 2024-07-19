package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Busca arquivos de vídeo e sua legenda",
	Run: func(cmd *cobra.Command, args []string) {

		var path string

		// Caso não possua nenhum parametro, apenas retorna
		// TODO: Aqui é preciso mostrar alguma mensagem, dar alguma opção para o usuário
		if len(args) == 0 {
			return
		}

		// pega o path
		path = args[0]

		// TODO: Verifica se realmente é um path, vou adicionar a possibilidade de procurar direto pelo nome do filme

		// busca a lista de arquivos de vídeo existentes na pasta
		// e solicita para o usuário qual ele deseja

		files, err := PerformSearch(path)

		if err != nil {
			log.Fatal(err)
		}

		if len(files) == 0 {
			fmt.Printf("Esse diretório não possui arquivos de vídeo: %s\n", path)
			return
		}

		nameFile, err := AskForFileName(files)

		if err != nil {
			log.Fatal(err)
		}

		// TODO: Aqui é preciso mostrar alguma mensagem, dar alguma opção para o usuário
		if nameFile == "" {
			return
		}

		fmt.Printf("Nome do filme: %s\n", nameFile)

		// search for movie id
		movieId, err := GetMovieIdByName(nameFile)

		if err != nil {
			log.Fatal(err)
		}

		userToken, err := Login("evangelistamt7", "mtsec@1920")

		if err != nil {
			log.Fatal(err)
		}

		// manda api filmes
		GetSub(userToken, movieId)
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
}
