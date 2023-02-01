package middleware

import (
	"truffle/utils"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/logging"
	"go.uber.org/zap"
)

type BreakerOpt struct {
	Logger struct {
		Path string
	}
	Conf *config.Entity
}

func (opt BreakerOpt) Scheme() string {
	return "breaker"
}

func (opt BreakerOpt) Init(env string, c *utils.Conf) utils.IOpt {
	opt.Logger.Path = c.GetString("breaker.logger.path")
	opt.Conf = config.NewDefaultConfig()
	opt.Conf.Sentinel.Log.Logger = logging.NewConsoleLogger()
	opt.Conf.Sentinel.Log.Dir = opt.Logger.Path
	err := sentinel.InitWithConfig(opt.Conf)
	if err != nil {
		zap.S().Fatal(err)
	}
	return opt
}

func LoadRule(path string, threshold float64, interval uint32) {
	if _, err := flow.LoadRules([]*flow.Rule{
		{
			Resource:               path,
			Threshold:              threshold,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			StatIntervalInMs:       interval,
		},
	}); err != nil {
		zap.S().Fatalf("Unexpected error: %+v", err)
		return
	}
}
