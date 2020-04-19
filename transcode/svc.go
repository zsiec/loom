package transcode

type Svc interface {
	Create(CreateRequest) (CreateResponse, error)
	Status(string) (StatusResponse, error)
}
