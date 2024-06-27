package iomanager

type IOManager interface {
	ReadLines() ([]string, error)
	WriteJson(any) error
}
