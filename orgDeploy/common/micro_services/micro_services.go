package micro_services

import (
	"github.com/asim/go-micro/v3"
)

type SetupServer interface {
	Setup() micro.Service
}
