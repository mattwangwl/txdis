package txdis

import "errors"

const (
	defaultSaltString = "gogopowerranger"
)

var (
	ErrOutOfPartitionRange = errors.New("value out of partition range")
	ErrNoWorkerInPool      = errors.New("no worker in pool")
)
