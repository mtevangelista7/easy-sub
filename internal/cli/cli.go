package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "easysub",
	Short: "EasySub - Ferramenta para buscar legendas de filmes",
	Long: `EasySub é uma aplicação de linha de comando que permite aos usuários buscar 
e baixar legendas de filmes de forma fácil e rápida. Utilize os subcomandos para realizar
as operações desejadas.`,
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("Erro:%w", err)
	}

	return nil
}
