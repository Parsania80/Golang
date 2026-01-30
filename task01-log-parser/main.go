package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func readFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(content), nil
}

func OpenFile(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	return file
}

func WriteFile(filePath string, data string) error {

	inFile := OpenFile(filePath)
	_, err := inFile.WriteString(data)
	if err != nil {
		return err
	}
	defer inFile.Close()

	return nil
}

func parseFileContent(filePath string, currentDir string) {
	// Implement your parsing logic here

	inFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	errorFile := OpenFile(currentDir + ERROR_FILE)
	warningFile := OpenFile(currentDir + WARNING_FILE)
	infoFile := OpenFile(currentDir + INFO_FILE)
	errorCount := 0
	warningCount := 0
	infoCount := 0
	defer inFile.Close()
	defer errorFile.Close()
	defer warningFile.Close()
	defer infoFile.Close()

	scanner := bufio.NewScanner(inFile)
	lineinex := 0

	for scanner.Scan() {
		lineinex++
		line := scanner.Text()

		formattedLine := fmt.Sprintf("%d %s\n", lineinex, line)
		if strings.Contains(line, ERROR_KEY) {
			errorFile.WriteString(formattedLine)
			errorCount++
		} else if strings.Contains(line, WARNING_KEY) || strings.Contains(line, WARN_KEY) {
			warningFile.WriteString(formattedLine)
			warningCount++
		} else if strings.Contains(line, INFO_KEY) {
			infoFile.WriteString(formattedLine)
			infoCount++
		}

	}
	scannerErr := scanner.Err()
	if scannerErr != nil {
		panic(scannerErr)
	}

	errorFile.WriteString("Number of Errors: " + strconv.Itoa(errorCount))
	warningFile.WriteString("Number of Warnings: " + strconv.Itoa(warningCount))
	infoFile.WriteString("Number of Info: " + strconv.Itoa(infoCount))

}

func main() {

	var fileName string
	if len(os.Args) != 2 {
		panic("Please provide the log file name as an argument")
	}
	fileName = os.Args[1]

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := filepath.Join(dir, fileName)
	parseFileContent(path, dir+"/")
}
