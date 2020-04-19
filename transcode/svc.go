package transcode

import "context"

type Svc interface {
	Create(context.Context, CreateRequest) (CreateResponse, error)
	Status(context.Context, string) (StatusResponse, error)
}
