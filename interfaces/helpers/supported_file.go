package helpers

import "path/filepath"

func CheckSupportedFile(filename string) (string, bool) {
	supportedFileTypes := map[string]bool{
		".png":  true,
		".jpeg": true,
		".jpg":  true,
	}
	fileExtension := filepath.Ext(filename)

	return fileExtension, !supportedFileTypes[fileExtension]
}
