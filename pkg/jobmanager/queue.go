package jobmanager

type Queue interface {
	Enqueue(job Job)
	Dequeue() Job
}
