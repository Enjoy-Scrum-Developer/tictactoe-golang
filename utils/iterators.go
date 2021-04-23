package utils

type Iterator interface {
	Iterate(func(int, string))
}

type DiagonalIndexIterator struct {
	Size    int
	Squares []string
}

type BackwardDiagonalIndexIterator struct {
	Size    int
	Squares []string
}

type HorizontalIndexIterator struct {
	Size    int
	Squares []string
	Row     int
}

type VerticalIndexIterator struct {
	Size    int
	Squares []string
	Column  int
}

func NewDiagonalIndexIterator(size int, squares []string) *DiagonalIndexIterator {
	return &DiagonalIndexIterator{Size: 3, Squares: squares}
}

func (iter DiagonalIndexIterator) Iterate(iteratorFunc func(int, string)) {
	iterateByIndexFunc(func(i int) int {
		return i*iter.Size + i
	}, iter.Size, iter.Squares, iteratorFunc)
}

func NewBackwardDiagonalIndexIterator(size int, squares []string) *BackwardDiagonalIndexIterator {
	return &BackwardDiagonalIndexIterator{Size: 3, Squares: squares}
}

func (iter BackwardDiagonalIndexIterator) Iterate(iteratorFunc func(int, string)) {
	iterateByIndexFunc(func(i int) int {
		return i*iter.Size + (iter.Size - i - 1)
	}, iter.Size, iter.Squares, iteratorFunc)
}

func NewHorizontalIndexIterator(size int, squares []string, row int) *HorizontalIndexIterator {
	return &HorizontalIndexIterator{Size: size, Squares: squares, Row: row}
}

func (iter HorizontalIndexIterator) Iterate(iteratorFunc func(int, string)) {
	iterateByIndexFunc(func(i int) int {
		return iter.Row*iter.Size + i
	}, iter.Size, iter.Squares, iteratorFunc)
}

func NewVerticalIndexIterator(size int, squares []string, column int) *VerticalIndexIterator {
	return &VerticalIndexIterator{Size: size, Squares: squares, Column: column}
}

func (iter VerticalIndexIterator) Iterate(iteratorFunc func(int, string)) {
	iterateByIndexFunc(func(i int) int {
		return iter.Size*i + iter.Column
	}, iter.Size, iter.Squares, iteratorFunc)
}

func iterateByIndexFunc(indexFunc func(i int) int, size int, squares []string, iteratorFunc func(int, string)) {
	for i := 0; i < size; i++ {
		index := indexFunc(i)
		iteratorFunc(index, squares[index])
	}
}
