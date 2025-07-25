package service

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func FindLastUpdated() (string, error) {
	// Путь к папке с сохранениями
	saveDir := `C:\Users\Andry\AppData\Local\Potionomics\Saved\SaveGames`

	// Шаблон имени файла
	prefix := "PotionomicsSaveData"
	suffix := ".sav"

	var latestFile string
	var latestModTime time.Time

	entries, err := os.ReadDir(saveDir)
	if err != nil {
		fmt.Println("Ошибка при чтении директории:", err)
		return "", err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()

		if strings.HasPrefix(name, prefix) && strings.HasSuffix(name, suffix) {
			info, err := entry.Info()
			if err != nil {
				fmt.Println("Ошибка при получении информации о файле:", err)
				continue
			}

			modTime := info.ModTime()
			if modTime.After(latestModTime) {
				latestModTime = modTime
				latestFile = filepath.Join(saveDir, name)
			}
		}
	}

	return latestFile, nil
}
