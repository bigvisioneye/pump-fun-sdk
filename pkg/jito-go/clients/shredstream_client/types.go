package shredstream_client

import (
	"github.com/bigvisioneye/pump-fun-sdk/pkg/jito-go/pkg"
	"github.com/bigvisioneye/pump-fun-sdk/pkg/jito-go/proto"
	"github.com/gagliardetto/solana-go/rpc"
	"google.golang.org/grpc"
)

type client struct {
	GrpcConn *grpc.ClientConn
	RpcConn  *rpc.Client

	ShredstreamClient proto.ShredstreamClient

	Auth *pkg.AuthenticationService
}
