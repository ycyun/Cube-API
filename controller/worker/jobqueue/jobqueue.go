package jobqueue

func (q *JobQueue) Push(job *Job) {
	q.Jobs = append(q.Jobs, job)
}
func (q *JobQueue) Pop() *Job {
	if len(q.Jobs) > 0 {
		job := q.Jobs[0]
		q.Jobs = q.Jobs[1:]
		return job
	}
	return nil
}
