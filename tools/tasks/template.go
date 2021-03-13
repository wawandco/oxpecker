package tasks

var taskTemplate string = `package tasks
var _ = grift.Namespace("{{.}}", func() error{
	return nil
})
}`
