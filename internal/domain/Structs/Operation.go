package domainstructs

type Operation struct {
	FirstOperand  int64 `json:",omitempty"`
	SecondOperand int64 `json:",omitempty"`
}

type Result struct {
	Is int64
}
