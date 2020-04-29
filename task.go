package txdis

type JobFunc func() (interface{}, error)

// NewTaskR 隨機分配 worker
func NewTaskR(job JobFunc) ITask {
	return NewTask(_random.Int63(), job)
}

// NewTask 保證相同的id 在同一個worker 的隊列中及執行
func NewTask(id int64, job JobFunc) ITask {
	return &Task{
		job: job,
		id:  id,
		c:   make(chan TaskResult, 1),
	}
}

type ITask interface {
	ID() int64
	Exec()
	Result() TaskResult
}

type TaskResult struct {
	Payload interface{}
	Error   error
}

type Task struct {
	job    JobFunc
	id     int64
	c      chan TaskResult
	result TaskResult
}

func (t *Task) ID() int64 {
	return t.id
}

func (t *Task) Exec() {
	defer func() {
		close(t.c)
	}()

	payload, err := t.job()

	t.c <- TaskResult{
		Payload: payload,
		Error:   err,
	}
}

func (t *Task) Result() TaskResult {
	if result, ok := <-t.c; ok {
		t.result = result
	}
	return t.result
}
