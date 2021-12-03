package mysql

import (
	"xorm.io/xorm"
)

func newEngineGroup(conf *Config, isShowSQL bool) (*xorm.EngineGroup, error) {
	engineGroup, err := xorm.NewEngineGroup("mysql", conf.Conn)
	if err != nil {
		return engineGroup, err
	}

	if conf.MaxIdle != 0 {
		engineGroup.SetMaxIdleConns(conf.MaxIdle)
	}
	if conf.MaxOpen != 0 {
		engineGroup.SetMaxOpenConns(conf.MaxOpen)
	}
	if conf.MaxLifetime != 0 {
		engineGroup.SetConnMaxLifetime(conf.MaxLifetime)
	}

	engineGroup.ShowSQL(isShowSQL)
	return engineGroup, engineGroup.Ping()
}

func EngineGroupInstance(key string, isShowSQL bool) (*xorm.EngineGroup, error) {
	c, err := getConfig(key)
	if err != nil {
		return nil, err
	}

	engine, err := newEngineGroup(c, isShowSQL)
	if err != nil {
		return nil, err
	}

	return engine, nil
}
