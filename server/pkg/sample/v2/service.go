package v2

type SampleService struct {
	version string
}

func NewSampleService() *SampleService {
	return &SampleService{
		version: "v2",
	}
}

func (s *SampleService) Hello() string {
	return "Hello! v2"
}
