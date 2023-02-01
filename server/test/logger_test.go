package test

import (
	"errors"
	"testing"
	"time"
	"truffle/utils"

	"go.uber.org/zap"
)

func TestInitLogger(t *testing.T) {
	// 1. 配置log参数
	conf := utils.LogConfigs{
		Level:          "debug",    // 输出日志级别 "debug" "info" "warn" "error"
		Format:         "logfmt",   // 输出日志格式 logfmt, json
		Path:           "./log",    // 输出日志文件位置
		FileName:       "test.log", // 输出日志文件名称
		FileMaxSize:    1,          // 输出单个日志文件大小，单位MB
		FileMaxBackups: 10,         // 输出最大日志备份个数
		MaxAge:         1000,       // 日志保留时间，单位: 天 (day)
		Compress:       false,      // 是否压缩日志
		Stdout:         false,      // 是否输出到控制台
	}
	// 2. 初始化log
	if err := utils.InitLogger(conf); err != nil {
		t.Fatal(err)
	}

	// 3. 调用 Logger 打印日志测试
	zap.S().Infof("测试 Infof 用法：%s", "111")   // logger Infof 用法
	zap.S().Debugf("测试 Debugf 用法：%s", "111") // logger Debugf 用法
	go func() {
		for i := 0; i < 100000; i++ {
			zap.S().Infof("(1)协程内部调用测试 Infof 用法：%s", "111")
			time.Sleep(time.Millisecond)
		}
	}()
	zap.S().Errorf("测试 Errorf 用法：%s", "111") // logger Errorf 用法
	zap.S().Warnf("测试 Warnf 用法：%s", "111")   // logger Warnf 用法
	zap.S().Infof("测试 Infof 用法：%s, %d, %v, %f", "111", 1111, errors.New("collector returned no data"), 3333.33)
	// logger With 用法
	logger := zap.S().With("collector", "cpu", "name", "主机")
	logger.Infof("测试 (With + Infof) 用法：%s", "测试")
	zap.S().Errorf("测试 Errorf 用法：%s", "111")
	go func() {
		for i := 0; i < 100000; i++ {
			zap.S().Infof("(2)协程内部调用测试 Infof 用法：%s", "111")
			time.Sleep(time.Millisecond)
		}
	}()
	time.Sleep(time.Minute)
}
