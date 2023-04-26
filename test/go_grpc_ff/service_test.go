package go_grpc_ff

import (
	"testing"

	service "github.com/please-build/proto-rules/test/go_grpc_ff/proto"
)

func TestService(t *testing.T) {
	_ = service.Request{}
}
