package log

import (
	log "github.com/mikesen1994/blockchain_golang/logcustom"
	"sync"
	"testing"
)

func TestLogPrint(t *testing.T) {
	t.Log("测试打印效果")
	{
		log.IsColor(true)
		log.Trace("测试打印效果Trace")
		log.Tracef("%s","测试打印效果Tracef")
		log.Info("测试打印效果Info")
		log.Infof("%s","测试打印效果Infof")
		log.Debug("测试打印效果Info")
		log.Debugf("%s","测试打印效果Debugf")
		log.Warn("测试打印效果Warn")
		log.Warnf("%s","测试打印效果Warnf")
		log.Error("测试打印效果Error")
		log.Errorf("%s","测试打印效果Errorf")
		log.Fatal("测试打印效果Fatal")
		mylog:=log.New()
		mylog.Trace("测试打印效果Trace")
		mylog.Tracef("%s","测试打印效果Tracef")
		mylog.Info("测试打印效果Info")
		mylog.Infof("%s","测试打印效果Infof")
		mylog.Debug("测试打印效果Info")
		mylog.Debugf("%s","测试打印效果Debugf")
		mylog.Warn("测试打印效果Warn")
		mylog.Warnf("%s","测试打印效果Warnf")
		mylog.Error("测试打印效果Error")
		mylog.Errorf("%s","测试打印效果Errorf")
		mylog2:=log.New()
		mylog2.IsColor(true)
		mylog2.SetLogDiscardLevel(log.Levelerror)
		mylog2.Trace("测试打印效果Trace")
		mylog2.Tracef("%s","测试打印效果Tracef")
		mylog2.Info("测试打印效果Info")
		mylog2.Infof("%s","测试打印效果Infof")
		mylog2.Debug("测试打印效果Debug")
		mylog2.Debugf("%s","测试打印效果Debugf")
		mylog2.Warn("测试打印效果Warn")
		mylog2.Warnf("%s","测试打印效果Warnf")
		mylog2.Error("测试打印效果Error")
		mylog2.Errorf("%s","测试打印效果Errorf")
	}
}


func TestLogMultithread(t *testing.T) {
	t.Log("测试日志是否线程安全")
	{
		wait := sync.WaitGroup{}
		wait.Add(20000)
		for i := 0; i < 10000; i++ {
			go func(i int) {
				for j := 0; j <= 10; j++ {
					Tracef("1试试可以正常打印吗！这是第%d次 第%d次！", i, j)
				}
				wait.Done()
			}(i)
		}
		for i := 0; i < 10000; i++ {
			go func(i int) {
				for j := 0; j <= 10; j++ {
					Tracef("2试试可以正常打印吗！这是第%d次 第%d次！", i, j)
				}
				wait.Done()
			}(i)
		}
		wait.Wait()
	}
}
