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

	return value, err

}

func comprobeStringInput(r io.Reader) (input string) {

	readThis := bufio.NewScanner(r)

	readThis.Scan()

	input = readThis.Text()

	return

}
