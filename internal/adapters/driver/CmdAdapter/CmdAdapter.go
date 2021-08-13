package cmdadapter

import (
	"bufio"
	"os"

	"github.com/AlejandroWaiz/goodfriendtimmy/internal/domain"
)

type CmdAdapter struct {
	port    domain.DomainPort
	scanner *bufio.Scanner
}

func CreateCmdAdapter(port domain.DomainPort) *CmdAdapter {

	scanner := bufio.NewScanner(os.Stdin)

	cmd := CmdAdapter{port: port, scanner: scanner}

	return &cmd

}

//This should be a CLI app where you can do math operations with the numbers given to the cmd in.
//This layer would be connected via port to the domain and there is where the business logic will reside.
