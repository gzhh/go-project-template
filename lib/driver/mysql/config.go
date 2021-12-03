package mysql

import (
	"github.com/micro/go-micro/v2/config"
	"strings"
	"time"
)

const driverName = "mysql"

type Config struct {
	Conn        []string      `json:"conn"`
	MaxIdle     int           `json:"max_idle"`
	MaxOpen     int           `json:"max_open"`
	MaxLifetime time.Duration `json:"max_lifetime"`
}

func (c *Config) Load(k string, v interface{}) {
	switch k {
	case "conn":
		v1 := v.(string)
		conn := make([]string, 0)
		conn = append(conn, strings.Split(v1, ",")...)
		c.Conn = conn
	case "max_lifetime":
		v1 := v.(int)
		c.MaxLifetime = time.Duration(v1) * time.Nanosecond
	case "max_idle":
		c.MaxIdle = v.(int)
	case "max_open":
		c.MaxOpen = v.(int)
	}
}

func getConfig(key string) (*Config, error) {
	var c Config
	err := config.Get(driverName, key).Scan(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
