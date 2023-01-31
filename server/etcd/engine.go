package etcd

import (
	"context"
	"log"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// ? simply factory

type IEEngine interface {
	Run(string, string, int64) error
}

type engine struct {
	cli    *clientv3.Client
	ld     clientv3.LeaseID
	ctx    context.Context
	cancel context.CancelFunc
}

func (e *engine) Run(name, addr string, expire int64) error {
	err := e.CreateLease(expire)
	if err != nil {
		return err
	}
	err = e.BindLease(name, addr)
	heartbeat, err := e.KeepAlive()
	if err != nil {
		return err
	}
	go e.Watch(heartbeat)
	return nil
}

func (e *engine) CreateLease(expire int64) error {
	res, err := e.cli.Grant(e.ctx, expire)
	if err != nil {
		log.Println("Fail to create ld")
		return err
	}
	e.ld = res.ID
	return nil
}

func (e *engine) BindLease(k, v string) error {
	res, err := e.cli.Put(e.ctx, k, v, clientv3.WithLease(e.ld))
	if err != nil {
		log.Println("Fail to bind lease")
		return err
	}
	log.Println("Success to bind lease", res)
	return err
}

func (e *engine) KeepAlive() (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	res, err := e.cli.KeepAlive(e.ctx, e.ld)
	if err != nil {
		log.Println("Fail to maintain keepalive")
		return res, err
	}
	return res, nil
}

func (e *engine) Close() error {
	e.cancel()
	log.Println("Close...")
	e.cli.Revoke(e.ctx, e.ld)
	return e.cli.Close()
}

func (e *engine) Watch(heartbeat <-chan *clientv3.LeaseKeepAliveResponse, events ...string) {
	for {
		select {
		case ld := <-heartbeat:
			if ld == nil {
				e.Close()
			}
			log.Printf("Heartbeat: %v \n", ld)
		case <-e.ctx.Done():
			log.Println("Server close...")
			return
		}
	}
}

type vengine struct{}

func (e *vengine) Run(string, string, int64) error {
	return nil
}
