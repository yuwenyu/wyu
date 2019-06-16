package kernel

type Test interface {
	T() string
}

type test struct {
	t string
}

var _ Test = &test{}

func NewTest() *test {
	return &test{t:"test success"}
}

func (t *test) T() string {
	return t.t
}
