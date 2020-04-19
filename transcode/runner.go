package transcode

type TaskRunner interface {
	Run([]Task) (string, error)
}

type DefaultTaskRunner struct{}

func (DefaultTaskRunner) Run([]Task) (string, error) {
	return "", nil //todo
}
