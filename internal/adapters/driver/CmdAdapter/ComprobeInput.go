package cmdadapter

import "strconv"

func (cmd *CmdAdapter) comprobeInput(input string) (value int, err error) {

	cmd.scanner.Scan()

	value, err = strconv.Atoi(input)

	if err != nil {
		return
	}

	return

}
