package etcd

import (
	"context"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc/resolver"
)

// ? simply factory

type IBuilder interface {
	Scheme() string
}

// implement grpc.builder interface
type builder struct {
	cli *clientv3.Client
}

func (b *builder) Scheme() string {
	return "truffle"
}

func (b *builder) Build(target resolver.Target, cc resolver.ClientConn,
	opts resolver.BuildOptions) (resolver.Resolver, error) {
	prefix := "/" + target.URL.Scheme + target.URL.Path
	res, err := b.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		zap.S().Warn("Build etcd get addr failed; err:", err)
		return nil, err
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	es := &eResolver{
		cc:     cc,
		cli:    b.cli,
		ctx:    ctx,
		cancel: cancelFunc,
		scheme: prefix,
	}
	zap.S().Infof("etcd res:%+v\n", res)
	for _, kv := range res.Kvs {
		es.store(kv.Key, kv.Value)
	}
	es.updateState()
	go es.watch()
	return &eResolver{}, nil
}

type vbuilder struct{}

func (b *vbuilder) Scheme() string {
	return "truffle"
}
