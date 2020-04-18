package transcode

type Service interface {
	CreateJob(CreateJobRequest) (string, error)
	JobStatus(string) (JobStatusResponse, error)
}
