package main

import (
	"github.com/qq51529210/video-monitor/cloudrecord/api"
	"github.com/qq51529210/video-monitor/cloudrecord/db"

	"github.com/qq51529210/log"
)

// @Title		接口文档
// @version	1.0.0
func main() {
	defer func() {
		log.Recover(recover())
	}()
	// 配置
	err := loadCfg()
	if err != nil {
		panic(err)
	}
	// 日志
	// err = util.InitLog(&_cfg.Log)
	if err != nil {
		panic(err)
	}
	// 数据库
	err = db.Init(_cfg.DB.URL)
	if err != nil {
		panic(err)
	}
	// api 服务
	api.Serve(_cfg.Port)
}
