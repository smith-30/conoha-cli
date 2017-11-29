package model

type (
	Halt struct {
		OSStop string `json:"os-stop"`
	}
)

func NewHalt() *Halt {
	return &Halt{
		"",
	}
}
