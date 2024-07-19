package io

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

const videoExtensions string = ".mp4 .avi .mkv .mov .wmv .flv"

func SearchAllFileName(path string) ([]string, error) {
	// Verifica se o diretorio existe
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("O diretório informado não existe: %s", path)
		return nil, err
	}

	files, err := os.ReadDir(path)

	if err != nil {
		log.Printf("Não foi possível realizar a leitura do diretório: %s", path)
		return nil, err
	}

	// cria um slice de strings com o tamanho da quantidade de arquivos
	filesNames := make([]string, len(files))

	// preenche o slice com o nome de todos os arquivos
	for i, file := range files {
		if isVideoFile(file.Name()) {
			filesNames[i] = strings.Split(file.Name(), ".")[0]
		}
	}

	// retorna todos os nomes
	return filesNames, nil
}

func isVideoFile(fileName string) bool {
	// Obtém a extensão do arquivo
	ext := strings.ToLower(filepath.Ext(fileName))

	// Verifica se a extensão está na string de extensões de vídeo
	return strings.Contains(videoExtensions, ext)
}
