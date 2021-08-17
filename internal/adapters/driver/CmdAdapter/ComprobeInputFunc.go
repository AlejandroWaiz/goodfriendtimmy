package cmdadapter

import (
	"bufio"
	"io"
	"strconv"
)

func comprobeNumberInput(r io.Reader) (value float64, err error) {

	readThis := bufio.NewScanner(r)

	readThis.Scan()

	input := readThis.Text()

	intInput, err := strconv.Atoi(input)

	value = float64(intInput)

	if err != nil {
		return
	}

	return

}

func comprobeStringInput(r io.Reader) (input string, err error) {

	readThis := bufio.NewScanner(r)

	readThis.Scan()

	input = readThis.Text()

	if err != nil {
		return
	}

	return

}