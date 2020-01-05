package vppagent

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/networkservicemesh/networkservicemesh/controlplane/api/crossconnect"
	"github.com/networkservicemesh/networkservicemesh/forwarder/api/forwarder"
)

//RequestValidator returns Forwarder Server with validation for Request and Close
func RequestValidator() forwarder.ForwarderServer {
	return &requestValidator{}
}

type requestValidator struct {
}

func (n *requestValidator) Request(ctx context.Context, request *crossconnect.CrossConnect) (*crossconnect.CrossConnect, error) {
	if err := request.IsValid(); err != nil {
		Logger(ctx).Errorf("Reuqest: %v is not valid, reason: %v", request, err)
		return nil, err
	}
	if next := Next(ctx); next != nil {
		return next.Request(ctx, request)
	}
	return request, nil
}

func (n *requestValidator) Close(ctx context.Context, request *crossconnect.CrossConnect) (*empty.Empty, error) {
	if err := request.IsValid(); err != nil {
		Logger(ctx).Errorf("Close: %v is not valid, reason: %v", request, err)
		return new(empty.Empty), err
	}
	if next := Next(ctx); next != nil {
		return next.Close(ctx, request)
	}
	return new(empty.Empty), nil
}
