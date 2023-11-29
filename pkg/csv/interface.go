package csv

type Reader interface {
	Read() ([][]string, map[string]int, error)
}
