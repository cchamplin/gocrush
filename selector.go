package gocrush

type Selector interface {
	Select(input int64, round int64) Node
}
