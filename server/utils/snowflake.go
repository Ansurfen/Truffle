package utils

import (
	"time"

	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

func init() {
	st, err := time.Parse("2006-01-02", "2023-01-01")
	AssertWithPanic(err)
	snowflake.Epoch = st.UnixNano() / 1e6
}

type SnowFlake struct {
	*snowflake.Node
}

func NewSnowFlake(num int64) *SnowFlake {
	node, err := snowflake.NewNode(num)
	if err != nil {
		zap.S().Fatal(err)
	}
	zap.S().Infof("init snowflake node_%d", num)
	return &SnowFlake{node}
}
