package transcode

type Svc interface {
	CreateJob(CreateJobRequest) (string, error)
	JobStatus(string) (JobStatusResponse, error)
}
