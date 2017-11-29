package model

type (
	Boot struct {
		OSStart string `json:"os-start"`
	}
)

func NewBoot() *Boot {
	return &Boot{
		nil,
	}
}
