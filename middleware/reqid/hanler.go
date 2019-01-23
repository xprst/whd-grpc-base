package req_uuid

import (
	"context"
	"github.com/xprst/whd-grpc-base/util"
	"google.golang.org/grpc/metadata"
)

var DefaultXRequestIDKey = "x-request-id"

func HandleRequestID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return newRequestID()
	}

	header, ok := md[DefaultXRequestIDKey]
	if !ok || len(header) == 0 {
		return newRequestID()
	}

	requestID := header[0]
	if requestID == "" {
		return newRequestID()
	}

	return requestID
}

func newRequestID() string {
	return util.ShotUuid()
}
