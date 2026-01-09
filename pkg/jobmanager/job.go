package jobmanager

type Job interface {
	Process() error
	ID() string
}
