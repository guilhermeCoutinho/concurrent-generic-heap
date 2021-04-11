package heap

import "golang.org/x/xerrors"

var (
	EmptyHeap       = xerrors.New("Empty heap")
	ElementNotFound = xerrors.New("Element not found")
	IndexOutOfRange = xerrors.New("Index out of range")
)
