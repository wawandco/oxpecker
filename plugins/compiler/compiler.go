package compiler

// Buffalo
type Tool struct {
	output string
}

func (g Tool) Name() string {
	return "compiler"
}
