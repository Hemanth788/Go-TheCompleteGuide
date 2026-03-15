package iomanager

type IIOManager interface {
	WriteResultToOutput(data any) error
	ReadResultFromInput() ([]string, error)
}