package txdis

/*
task id 經過hash 計算後，決定要給哪一個worker 執行
同樣的id 會落在同樣的worker
*/

func New(config *Config) IDispatcher {
	workers := make([]*worker, config.WorkerNum)
	for i := 0; i < config.WorkerNum; i++ {
		worker := newWorker(config.QueueLength)
		go worker.Run()
		workers[i] = worker
	}

	obj := &dispatcher{
		workers:   workers,
		partition: newPartition(config.Salt, config.WorkerNum),
	}

	return obj
}

type IDispatcher interface {
	Delegate(task ITask) error
}

type dispatcher struct {
	workers   []*worker
	partition iPartition
}

func (d *dispatcher) Delegate(task ITask) error {
	if len(d.workers) == 0 {
		return ErrNoWorkerInPool
	}

	worker, err := d.dispatch(task)
	if err != nil {
		return err
	}

	worker.queue <- task

	return nil
}

func (d *dispatcher) dispatch(task ITask) (*worker, error) {
	loc := d.partition.Calculate(task.ID())

	if loc >= len(d.workers) {
		return nil, ErrOutOfPartitionRange
	}

	worker := d.workers[loc]
	return worker, nil
}

func newWorker(queueLength int) *worker {
	return &worker{
		queue: make(chan ITask, queueLength),
	}
}

type worker struct {
	queue chan ITask
}

func (w *worker) Run() {
	for t := range w.queue {
		t.Exec()
	}
}
