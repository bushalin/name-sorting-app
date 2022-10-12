package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// selecting the file to read from the arguments
	var filename string

	if len(os.Args) > 1 {
		args := os.Args[1:]
		if args[0] == "--filename" {
			filename = args[1]
		}
	}

	if len(filename) == 0 {
		filename = "names.txt"
	}

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	// reading file contents
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	var nameSlice []string

	for scanner.Scan() {
		trimmedName := TrimName(string(scanner.Bytes()), ",")
		nameSlice = append(nameSlice, trimmedName)
	}

	sort.Strings(nameSlice)

	PrintResult(nameSlice)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("testy")
}

func TrimName(s, suffix string) string {
	if strings.HasSuffix(s, ",") {
		s = s[:len(s)-len(suffix)]
	}
	if strings.HasPrefix(s, "'") {
		s = s[1:]
	}
	if strings.HasSuffix(s, "'") {
		s = s[:len(s)-1]
	}
	return s
}

func PrintResult(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}
