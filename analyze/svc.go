package analyze

type Svc interface {
	Analyze(url string) (Info, error)
}
