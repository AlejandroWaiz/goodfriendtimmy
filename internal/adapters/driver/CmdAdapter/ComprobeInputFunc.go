package cmdadapter

import "strconv"

func (cmd *CmdAdapter) comprobeInput() (value int, err error) {

	cmd.scanner.Scan()

	input := cmd.scanner.Text()

	value, err = strconv.Atoi(input)

	if err != nil {
		return
	}

	return

}
