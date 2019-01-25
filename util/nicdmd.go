package util

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
)

func NiceMD(ctx context.Context) metautils.NiceMD{
	return metautils.ExtractIncoming(ctx)
}
