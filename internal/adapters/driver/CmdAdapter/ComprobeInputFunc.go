package cmdadapter

import (
	"bufio"
	"io"
	"strconv"
)

func comprobeNumberInput(r io.Reader) (value int64, err error) {

	readThis := bufio.NewScanner(r)

	readThis.Scan()

	input := readThis.Text()

	convertedInput, _ := strconv.ParseInt(input, 10, 64)

	if err != nil {
		return
	}

	return convertedInput, err

}

func comprobeStringInput(r io.Reader) (input string) {

	readThis := bufio.NewScanner(r)

	readThis.Scan()

	input = readThis.Text()

	return

}
