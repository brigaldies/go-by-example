package main

// Source: https://dev.to/napicella/go-for-java-developers-lambda-functions-to-the-rescue-655

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	inputPathname := "./data/config.txt"
	outputPathname := "./data/config-output.txt"
	inputFile, err := os.Open(inputPathname)
	if err != nil {
		fmt.Println("file open error", inputPathname, err)
		return
	}
	fmt.Println(fmt.Sprintf("Opened input file %s", inputPathname))

	outputFile, err := os.OpenFile(outputPathname, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("file create error", outputPathname, err)
		return
	}

	fmt.Println(fmt.Sprintf("Opened output file %s", outputPathname))

	err = findSectionAndKeys("HEADER", inputFile, outputFile)
	if err != nil {
		fmt.Println("findSectionAndKeys error", err)
		return
	}

	// Close files
	err = inputFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read and print the filtered output
	outputContent, err := os.ReadFile(outputPathname)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Output")
	fmt.Println(string(outputContent))

	fmt.Println("Success and done.")
}

func findSectionAndKeys(profileName string, in io.Reader, out io.Writer) error {
	var (
		line       string
		readError  error
		writeError error
	)

	profileFilter := newProfileFilter(profileName)

	fmt.Println(fmt.Sprintf("Profile filter: %v", profileFilter))

	reader := bufio.NewReader(in)

	for {
		line, readError = reader.ReadString('\n')

		fmt.Println(fmt.Sprintf("Processing line: %q", line))
		if profileFilter.match(line) {
			fmt.Println("\tMatch!")
			_, writeError = io.WriteString(out, line)
		}

		if readError != nil || writeError != nil {
			break
		}
	}

	if writeError != nil {
		return writeError
	}

	if readError != io.EOF {
		return readError
	}

	return nil
}

func newProfileFilter(profileName string) *profileFilter {
	var matchers []func(line string) bool

	matchers = append(matchers,
		matcher(startsWith).apply("#"+profileName),
		matcher(startsWith).apply("Key2="),
		matcher(startsWith).apply("Key3="),
	)

	return &profileFilter{matchers, profileName}
}

func startsWith(line string, toMatch string) bool {
	return strings.HasPrefix(
		strings.ToLower(strings.TrimSpace(line)),
		strings.ToLower(toMatch),
	)
}

type matcher func(line string, toMatch string) bool

func (f matcher) apply(toMatch string) func(line string) bool {
	return func(line string) bool {
		return f(line, toMatch)
	}
}

type profileFilter struct {
	matchers    []func(line string) bool
	profileName string
}

func (p *profileFilter) match(text string) bool {
	if len(p.matchers) == 0 {
		return false
	}

	shouldFilter := p.matchers[0](text)

	if shouldFilter {
		p.matchers = p.matchers[1:len(p.matchers)]
	}

	return shouldFilter
}
