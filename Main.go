package main

import (
	"flag"
	"github.com/golang/glog"
)

func main() {
	// 解析命令行参数
	logDir := flag.String("l", "log", "log directory")
	start := flag.Uint("s", 1, "start page index")
	end := flag.Uint("e", 10, "end page index")

	if *start > *end {
		glog.Fatalln("start page over than end page")
	}

	glog.Info("log dir:", *logDir)

	if *logDir == "" || *logDir == "stderr" {
		flag.Lookup("logtostderr").Value.Set("true")
	} else {
		flag.Lookup("log_dir").Value.Set(*logDir)
	}

	manager := NewDouBanManager(*start, *end)
	manager.Run()
}
