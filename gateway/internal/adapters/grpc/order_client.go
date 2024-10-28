// grpc/order_client.go
package grpc

import (
	"github.com/superdev/ecommerce/gateway/internal/adapters/grpc/proto"
	"github.com/superdev/ecommerce/gateway/internal/config"
	"go.uber.org/zap"

	"google.golang.org/grpc"
)

func NewOrderGrpcClient(
	conf *config.Config,
	logger zap.Logger,
) (proto.OrderServiceClient, error) {
	conn, err := grpc.Dial(conf.OrderServiceUrl, grpc.WithInsecure())
	if err != nil {
		// logger.Error(fmt.Sprintf("Could not connect to Order Service: %v", err))
		return nil, err
	}
	return proto.NewOrderServiceClient(conn), nil
}
