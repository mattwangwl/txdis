package txdis

import (
	"sync"
	"testing"
	"time"
)

func TestDispatcher(t *testing.T) {
	p := New(NewDefaultConfig())

	var mx sync.Mutex
	var wg sync.WaitGroup
	validator := make(map[int64]struct{})

	for i := int64(0); i < 10000; i++ {
		wg.Add(1)
		go func(num int64) {
			defer wg.Done()

			mx.Lock()
			validator[num] = struct{}{}
			mx.Unlock()

			t := NewTask(num, func() (interface{}, error) {
				// 假裝有在跑邏輯
				<-time.After(100 * time.Microsecond)
				return num, nil
			})
			p.Delegate(t)

			r := t.Result()
			mx.Lock()
			if _, ok := validator[r.Payload.(int64)]; ok {
				delete(validator, r.Payload.(int64))
			}
			mx.Unlock()
		}(i)
	}

	wg.Wait()

	if len(validator) > 0 {
		t.Error("validator size is not zero, size:", len(validator))
	}

}

func Test_dispatcher_Delegate(t *testing.T) {
	type fields struct {
		workers   []*worker
		partition iPartition
	}

	type args struct {
		task ITask
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "delegate",
			fields: fields{
				workers:   []*worker{newWorker(1)},
				partition: &mockPartition{calculateResult: 0},
			},
			args:    args{task: NewTask(12345, func() (interface{}, error) { return nil, nil })},
			wantErr: false,
		},
		{
			name: "partition out of range",
			fields: fields{
				workers:   []*worker{newWorker(1)},
				partition: &mockPartition{calculateResult: 1},
			},
			args:    args{task: NewTask(12345, func() (interface{}, error) { return nil, nil })},
			wantErr: true,
		},
		{
			name: "empty worker pool",
			fields: fields{
				workers:   []*worker{},
				partition: &mockPartition{calculateResult: 1},
			},
			args:    args{task: NewTask(12345, func() (interface{}, error) { return nil, nil })},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dispatcher{
				workers:   tt.fields.workers,
				partition: tt.fields.partition,
			}
			if err := d.Delegate(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("Delegate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type mockPartition struct {
	calculateResult int
}

func (m *mockPartition) Calculate(id int64) int {
	return m.calculateResult
}
