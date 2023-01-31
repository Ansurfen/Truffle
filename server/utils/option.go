package utils

const (
	ENV         = "env"
	EMAIL       = "email"
	TIMER       = "timer"
	SERVICE     = "service"
	LOGGER      = "logger"
	DEFAULT     = "default"
	TRACER      = "tracer"
	BREAKER     = "breaker"
	ENV_DEVELOP = "develop"
	ENV_RELEASE = "release"
)

type IOpt interface {
	Scheme() string
	Init(string, *Conf) IOpt
	// Free()
}

type BaseOpt struct {
	conf *Conf
	opts map[string]IOpt
	env  string
}

func LoadOpt(env string, opts ...IOpt) *BaseOpt {
	baseOpt := &BaseOpt{env: env, conf: NewConf("application", "yaml", "."), opts: map[string]IOpt{}}
	var envOpt EnvOpt
	envOpt.Init(env, baseOpt.conf)
	baseOpt.opts[envOpt.Scheme()] = envOpt
	for _, opt := range opts {
		// if opt.Scheme() == ENV {
		// 	env = opt.Init(env, baseOpt.conf).(EnvOpt).Env
		// }
		baseOpt.opts[opt.Scheme()] = opt.Init(env, baseOpt.conf)
	}
	return baseOpt
}

func (opt *BaseOpt) Free(keys ...string) {
	for _, key := range keys {
		if _, ok := opt.opts[key]; ok {
			// TODO range every opt free
			delete(opt.opts, key)
		}
	}
}

func (opt *BaseOpt) Opt(key string) IOpt {
	return opt.opts[key]
}

func (opt *BaseOpt) Init() {

}

type ServiceOpt struct {
	Addr string
	Name string
	Pod  int
	Etcd struct {
		Addr   []string
		Expire int64
		Dial   int64
	}
	ENV      string // local env
	RPCPort  string
	HttpPort string
}

func (opt ServiceOpt) Scheme() string {
	return SERVICE
}

func (opt ServiceOpt) Init(env string, c *Conf) IOpt {
	opt.Name = c.GetString("server.name")
	opt.RPCPort = c.GetString("server.port")
	opt.HttpPort = c.GetString("server.http")
	opt.Etcd.Expire = c.GetInt64("etcd.time.expire")
	opt.Etcd.Dial = c.GetInt64("etcd.time.dial")
	// priority: local > global
	if len(opt.ENV) != 0 {
		env = opt.ENV
	}
	switch env {
	case ENV_DEVELOP:
		opt.Etcd.Addr = c.GetStringSlice("etcd.develop.addr")
		opt.Addr = c.GetString("server.develop.host") + ":" + opt.RPCPort
	case ENV_RELEASE:
		opt.Etcd.Addr = c.GetStringSlice("etcd.release.addr")
		opt.Addr = c.GetString("server.release.host") + ":" + opt.RPCPort
	}
	limit := c.GetInt("server.limit")
	port := c.GetInt("server.port")
	base := c.GetInt("server.base")
	if port+1 < base+limit {
		c.Set("server.port", port+1)
		c.WriteConfig()
	} else {
		c.Set("server.port", base)
		c.WriteConfig()
	}
	return opt
}

type LoggerOpt struct {
	Level          string
	Format         string
	Path           string
	FileName       string
	FileMaxSize    int
	FileMaxBackups int
	MaxAge         int
	Compress       bool
	Stdout         bool
}

func (opt LoggerOpt) Scheme() string {
	return LOGGER
}

func (opt LoggerOpt) Init(env string, c *Conf) IOpt {
	opt.Level = c.GetString("logger.level")
	opt.Format = c.GetString("logger.format")
	opt.Path = c.GetString("logger.path")
	opt.FileName = c.GetString("logger.filename")
	opt.FileMaxSize = c.GetInt("logger.maxsize")
	opt.Compress = c.GetBool("logger.compress")
	opt.Stdout = c.GetBool("logger.stdout")
	return opt
}

type EnvOpt struct {
	Env       string
	UseMQ     bool
	UseTracer bool
}

func (opt EnvOpt) Scheme() string {
	return ENV
}

func (opt EnvOpt) Init(env string, c *Conf) IOpt {
	opt.Env = c.GetString("env.this")
	opt.UseMQ = c.GetBool("env.mq")
	opt.UseTracer = c.GetBool("env.tracer")
	return opt
}

type DefaultOpt struct {
	Logger  LoggerOpt
	Service ServiceOpt
}

func (opt DefaultOpt) Scheme() string {
	return DEFAULT
}

func (opt DefaultOpt) Init(env string, c *Conf) IOpt {
	opt.Logger = opt.Logger.Init(env, c).(LoggerOpt)
	opt.Service = opt.Service.Init(env, c).(ServiceOpt)
	return opt
}
