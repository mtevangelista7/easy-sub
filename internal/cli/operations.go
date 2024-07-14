package cli

import "easysub/internal/io"

func PerformSearch(path string) {
	// pega uma lista com os arquivos de áudio
	filesNames := io.SearchAllFileName(path)

	for _, name := range filesNames {
		print("%s", name)
	}
}
