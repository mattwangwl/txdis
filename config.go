package txdis

type Config struct {
	// worker 數量
	WorkerNum int
	// 每一個worker 的queue 大小
	QueueLength int
	// 計算分區的鹽
	Salt []byte
}

func NewDefaultConfig() *Config {
	return &Config{
		WorkerNum:   8,
		QueueLength: 8,
		Salt:        []byte(defaultSaltString),
	}
}
