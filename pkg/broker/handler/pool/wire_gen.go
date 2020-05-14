// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package pool

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/cloudevents/sdk-go/v2/protocol/http"
	"github.com/google/knative-gcp/pkg/broker/config"
)

// Injectors from wire.go:

func InitializeTestFanoutPool(ctx context.Context, targets config.ReadonlyTargets, pubsubClient *pubsub.Client, opts ...Option) (*FanoutPool, error) {
	v := _wireValue
	protocol, err := http.New(v...)
	if err != nil {
		return nil, err
	}
	v2 := _wireValue2
	deliverClient, err := NewDeliverClient(protocol, v2...)
	if err != nil {
		return nil, err
	}
	retryClient, err := NewRetryClient(ctx, pubsubClient, v2...)
	if err != nil {
		return nil, err
	}
	fanoutPool, err := NewFanoutPool(targets, pubsubClient, deliverClient, retryClient, opts...)
	if err != nil {
		return nil, err
	}
	return fanoutPool, nil
}

var (
	_wireValue  = []http.Option(nil)
	_wireValue2 = DefaultCEClientOpts
)

func InitializeTestRetryPool(targets config.ReadonlyTargets, pubsubClient *pubsub.Client, opts ...Option) (*RetryPool, error) {
	v := _wireValue3
	protocol, err := http.New(v...)
	if err != nil {
		return nil, err
	}
	v2 := _wireValue4
	deliverClient, err := NewDeliverClient(protocol, v2...)
	if err != nil {
		return nil, err
	}
	retryPool, err := NewRetryPool(targets, pubsubClient, deliverClient, opts...)
	if err != nil {
		return nil, err
	}
	return retryPool, nil
}

var (
	_wireValue3 = []http.Option(nil)
	_wireValue4 = DefaultCEClientOpts
)
