package client

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"truffle/utils"

	"github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	"github.com/openzipkin/zipkin-go/reporter"
	httpreport "github.com/openzipkin/zipkin-go/reporter/http"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GClientConn struct {
	*grpc.ClientConn
}

func NewGClientConn(target string, opts ...grpc.DialOption) *GClientConn {
	var temp []grpc.DialOption
	temp = append(temp, grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if len(opts) > 0 {
		temp = append(temp, opts...)
	}
	temp = append(temp, grpc.WithTransportCredentials(insecure.NewCredentials()))
	dial, err := grpc.Dial(target, temp...)
	if err != nil {
		zap.S().Fatal(err)
		return nil
	}
	return &GClientConn{dial}
}

func NewGClientConns(paths []string, opts ...grpc.DialOption) map[string]*GClientConn {
	dials := make(map[string]*GClientConn)
	for _, path := range paths {
		dials[path] = NewGClientConn("truffle:///"+path, opts...)
	}
	return dials
}

func WithTracer(tracer *zipkin.Tracer) grpc.DialOption {
	return grpc.WithStatsHandler(zipkingrpc.NewClientHandler(tracer))
}

// ? simpley factory

type IService interface {
	Setup()
	Run()
	Shutdown()
}

// grpc base client
type GBClient struct {
	// opt
	srv *grpc.Server
	opt *utils.BaseOpt
}

func (c *GBClient) Setup()    {}
func (c *GBClient) Run()      {}
func (c *GBClient) Shutdown() {}

func NewGBClient() *GBClient {
	return &GBClient{}
}

// grpc client with tracer
type GTClient struct {
	Srv    *grpc.Server
	Rep    reporter.Reporter
	Tracer *zipkin.Tracer
	Opt    *utils.BaseOpt
}

func NewGTClient(opt *utils.BaseOpt) *GTClient {
	c := &GTClient{Opt: opt}
	if opt.Opt(utils.ENV).(utils.EnvOpt).UseTracer {
		tracer, rep, err := LoadZipkinTracer(
			opt.Opt(utils.TRACER).(TracerOpt).Addr,
			opt.Opt(utils.TRACER).(TracerOpt).Name,
			opt.Opt(utils.DEFAULT).(utils.DefaultOpt).Service.Addr)
		c.Srv = grpc.NewServer(grpc.StatsHandler(zipkingrpc.NewClientHandler(tracer)))
		if err != nil {
			zap.S().Fatal(err)
			return nil
		}
		c.Tracer, c.Rep = tracer, rep
	} else {
		c.Srv = grpc.NewServer()
	}
	return c
}

func (c *GTClient) Setup() {
}

func (c *GTClient) Run() {
	listen, err := net.Listen("tcp", c.Opt.Opt(utils.DEFAULT).(utils.DefaultOpt).Service.Addr)
	if err != nil {
		zap.S().Fatal("Fail to listen server")
		return
	}
	c.Srv.Serve(listen)
}

func (c *GTClient) Shutdown() {
	cs := make(chan os.Signal, 1)
	signal.Notify(cs, os.Interrupt)
	<-cs
	c.Rep.Close()
}

func (c *GTClient) Destory() {
	syscall.Exit(0)
}

type TracerOpt struct {
	Addr string
	Name string
}

func (opt TracerOpt) Scheme() string {
	return utils.TRACER
}

func (opt TracerOpt) Init(env string, c *utils.Conf) utils.IOpt {
	opt.Addr = c.GetString("tracer.develop.addr")
	opt.Name = c.GetString("tracer.develop.name")
	return opt
}

func LoadZipkinTracer(url, serviceName, addr string) (*zipkin.Tracer, reporter.Reporter, error) {
	// init zipkin reporter
	// use http service
	r := httpreport.NewReporter(url)
	// create a endpoint，which label current service
	endpoint, err := zipkin.NewEndpoint(serviceName, addr)
	if err != nil {
		return nil, r, err
	}
	// parse span，and context
	tracer, err := zipkin.NewTracer(r, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		return nil, r, err
	}
	return tracer, r, nil
}
