package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "find all movies files",
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		if len(args) == 0 {
			return
		}

		// pega o path
		path = args[0]

		// busca a lista de arquivos de vídeo existentes na pasta
		// e solicita para o usuário qual ele deseja
		nameFile := AskForFileName(PerformSearch(path))

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
