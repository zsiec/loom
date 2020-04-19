package analyze

import "context"

type Svc interface {
	Analyze(ctx context.Context, url string) (Info, error)
}
