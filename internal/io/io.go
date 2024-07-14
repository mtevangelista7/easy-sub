package io

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func SearchAllFileName(path string) []string {
	// Verifica se o diretorio existe
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("Não existe: %s", path)
		return nil
	}

	// realiza a leitura do diretorio
	files, err := os.ReadDir(path)

	// caso tenha erro
	if err != nil {
		log.Fatalf("sla %s", err)
		return nil
	}

	// cria um slice de strings
	filesNames := make([]string, len(files))

	// preenche o slice com o nome de todos os arquivos
	for i, file := range files {
		if isVideoFile(file.Name()) {
			filesNames[i] = file.Name()
		}
	}

	// retorna todos os nomes
	return filesNames
}

func isVideoFile(fileName string) bool {
	// String com extensões de vídeo comuns, separadas por espaços
	videoExtensions := ".mp4 .avi .mkv .mov .wmv .flv"

	// Obtém a extensão do arquivo
	ext := strings.ToLower(filepath.Ext(fileName))

	// Verifica se a extensão está na string de extensões de vídeo
	return strings.Contains(videoExtensions, ext)
}
