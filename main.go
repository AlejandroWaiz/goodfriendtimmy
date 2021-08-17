package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {

	// err := os.Setenv("MuxPort", "3000")
	// if err != nil {
	// 	log.Fatalf("Could not set env variables: %v", err)
	// }
	// domainPort := timmyadapter.CreateTimmyAdapter()
	// driverHandler := cmdadapter.CreateCmdAdapter(domainPort)
	// driverHandler.ListenAndServe()

	one := strings.NewReader("Hello")
	two := strings.NewReader("Mother")
	three := strings.NewReader("Ducker")

	ReadThis(one, two, three)

}

func ReadThis(one, two, three io.Reader) {

	fmt.Println(comprobeNumberInput(one))
	fmt.Println(comprobeNumberInput(two))
	fmt.Println(comprobeNumberInput(three))

}

func comprobeNumberInput(r io.Reader) (input string, err error) {

	readThis := bufio.NewScanner(r)

	readThis.Scan()

	input = readThis.Text()

	if err != nil {
		return
	}

	return

}
