package main

import (
	"github.com/qq51529210/video-monitor/onvif/api"
	"github.com/qq51529210/video-monitor/onvif/db"
	"github.com/qq51529210/video-monitor/onvif/task"

	"github.com/qq51529210/log"
	"github.com/qq51529210/util"
)

// @Title		接口文档
// @version	1.0.0
func main() {
	defer func() {
		// 抓异常
		log.Recover(recover())
	}()
	// 配置
	err := loadCfg()
	if err != nil {
		panic(err)
	}
	// 日志
	err = util.InitLog(&_cfg.Log)
	if err != nil {
		panic(err)
	}
	// 数据库
	err = db.Init(_cfg.DB.URL, _cfg.DB.EnableCache)
	if err != nil {
		panic(err)
	}
	// 检查
	err = task.Run(&_cfg.Task)
	if err != nil {
		panic(err)
	}
	// api 服务
	api.Serve(_cfg.Port)
}
