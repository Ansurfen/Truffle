package etcd

import (
	"context"
	"log"
	"sync"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

type eResolver struct {
	ctx    context.Context
	cancel context.CancelFunc
	cc     resolver.ClientConn
	cli    *clientv3.Client
	scheme string
	ips    sync.Map
}

func (e *eResolver) ResolveNow(resolver.ResolveNowOptions) {
	log.Println("Trying reconnect service to fail connect...")
}

func (e *eResolver) Close() {
	log.Println("Etcd resolver close...")
	e.cancel()
}

func (e *eResolver) store(k, v []byte) {
	e.ips.Store(string(k), string(v))
}

func (s *eResolver) del(key []byte) {
	s.ips.Delete(string(key))
}

func (e *eResolver) updateState() {
	var addrList resolver.State
	e.ips.Range(func(k, v interface{}) bool {
		tA, ok := v.(string)
		if !ok {
			return false
		}
		log.Printf("conn.UpdateState key[%v];val[%v]\n", k, v)
		addrList.Addresses = append(addrList.Addresses, resolver.Address{Addr: tA, ServerName: k.(string)})
		return true
	})
	e.cc.UpdateState(addrList)
}

func (e *eResolver) watch() {
	watchChan := e.cli.Watch(context.Background(), e.scheme, clientv3.WithPrefix())
	for {
		select {
		case val := <-watchChan:
			for _, event := range val.Events {
				switch event.Type {
				case mvccpb.PUT:
					e.store(event.Kv.Key, event.Kv.Value)
					log.Println("put:", string(event.Kv.Key))
					e.updateState()
				case mvccpb.DELETE:
					log.Println("del:", string(event.Kv.Key))
					e.del(event.Kv.Key)
					e.updateState()
				}
			}
		case <-e.ctx.Done():
			return
		}
	}
}
