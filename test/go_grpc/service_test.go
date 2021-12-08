package go_grpc

import (
	"testing"

	service "github.com/please-build/proto-rules/test/go_grpc/proto"
)

func TestService(t *testing.T) {
	_ = service.Request{}
}
