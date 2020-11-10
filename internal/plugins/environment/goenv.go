package environment

type GoEnv struct{}

func (g GoEnv) Name() string {
	return "goenv"
}
