// Чтение, скан и закрытие файла
package pkg

import (
	"bufio"
	"log"
	"os"
)

func GetStrings(filename string) []string {
	var lines []string
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return nil
	}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err != nil {
		log.Fatal(scanner.Err())
	}
	return lines
}
