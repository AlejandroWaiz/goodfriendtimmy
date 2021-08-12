package domainstructs

type Operation struct {
	FirstOperand  float64 `json:",omitempty"`
	SecondOperand float64 `json:",omitempty"`
}

type Result struct {
	Is float64
}
