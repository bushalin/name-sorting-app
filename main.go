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
	scanner.Split(bufio.ScanLines)

	var nameSlice []string

	for scanner.Scan() {
		// trimmedName := TrimName(string(scanner.Bytes()), ",")
		nameSlice = append(nameSlice, scanner.Text())
	}

	// sort.Strings(nameSlice)

	sss := PrintResult(nameSlice)
	sort.Slice(sss, func(i, j int) bool {
		asc1 := 0
		asc2 := 0
		for _, v := range sss[i] {
			asc1 += int(v)
		}
		for _, v := range sss[j] {
			asc2 += int(v)
		}

		return asc1 > asc2
	})

	for _, s := range sss {
		fmt.Println(s)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
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

func PrintResult(s []string) []string {
	sss := []string{}
	for _, v := range s {
		ss := strings.Split(v, ",")
		sss = append(sss, ss...)
	}

	return sss
}
