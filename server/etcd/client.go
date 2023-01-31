package etcd

import (
	"context"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

type IEClient interface {
	Builder() IBuilder
	Engine() IEEngine
}

func NewEClient(heart, dial bool, eps []string) IEClient {
	// ? tactic pattern + facade pattern
	if heart && dial {
		return NewEUClient(eps)
	} else if !heart && dial {
		return NewEDClient(eps)
	} else if heart && !dial {
		return NewESClient(eps)
	}
	return nil
}

// etcd dial client
type EDClient struct {
	*builder
	*vengine
	cli *clientv3.Client
}

func (c *EDClient) Builder() IBuilder {
	return c.builder
}

func (c *EDClient) Engine() IEEngine {
	return c.vengine
}

func NewEDClient(eps []string) *EDClient {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   eps,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		zap.S().Fatal("fail to start client")
	}
	return &EDClient{
		cli:     cli,
		builder: &builder{cli: cli},
	}
}

// etcd service client with heart comparing EClient
type ESClient struct {
	*vbuilder
	*engine
	cli *clientv3.Client
}

func NewESClient(eps []string) *ESClient {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   eps,
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		zap.S().Fatalf("new etcd client failed,error %v \n", err)
		return nil
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &ESClient{
		cli: cli,
		engine: &engine{
			cli:    cli,
			ctx:    ctx,
			cancel: cancelFunc,
		},
	}
}

func (c *ESClient) Builder() IBuilder {
	return c.vbuilder
}

func (c *ESClient) Engine() IEEngine {
	return c.engine
}

// etcd ultimate client with dial comparing EPClient
type EUClient struct {
	*builder
	*engine
	cli *clientv3.Client
}

func NewEUClient(eps []string) *EUClient {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   eps,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		zap.S().Fatalf("new etcd client failed,error %v \n", err)
		return nil
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	return &EUClient{
		cli: cli,
		builder: &builder{
			cli: cli,
		},
		engine: &engine{
			cli:    cli,
			ctx:    ctx,
			cancel: cancelFunc,
		},
	}
}

func (c *EUClient) Builder() IBuilder {
	return c.builder
}

func (c *EUClient) Engine() IEEngine {
	return c.engine
}
