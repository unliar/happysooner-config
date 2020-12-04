package happyConfig

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"os"
)

type Conf struct {
	Init config.Config
}

// 获取 MySQL 配置
func (c *Conf) GetMySQL(path ...string) (*MySQL, error) {
	m := MySQL{}
	err := c.Init.Get(path...).Scan(&m)
	if err != nil {
		return &m, err
	}
	return &m, nil
}

// 获取 Redis 配置
func (c *Conf) GetRedis(path ...string) (*Redis, error) {
	m := Redis{}
	err := c.Init.Get(path...).Scan(&m)
	if err != nil {
		return &m, err
	}
	return &m, nil
}

// 获取 smtp 配置
func (c *Conf) GetSMTP(path ...string) (*SMTP, error) {
	m := SMTP{}
	err := c.Init.Get(path...).Scan(&m)
	if err != nil {
		return &m, err
	}
	return &m, nil
}

// 获取 NSQ 配置
func (c *Conf) GetNSQ(path ...string) (*NSQ, error) {
	m := NSQ{}
	err := c.Init.Get(path...).Scan(&m)
	if err != nil {
		return &m, err
	}
	return &m, nil
}

// 获取微信公众号 配置
func (c *Conf) GetWechatMP(path ...string) (*WechatMP, error) {
	m := WechatMP{}
	err := c.Init.Get(path...).Scan(&m)
	if err != nil {
		return &m, err
	}
	return &m, nil
}

func NewHappyDefaultConf() (Conf, error) {
	conf, _ := config.NewConfig()
	consulSrc := consul.NewSource(
		consul.WithAddress(os.Getenv("MICRO_REGISTRY_ADDRESS")),
		consul.WithPrefix("/happy/config"),
		consul.StripPrefix(true),
	)
	err := conf.Load(consulSrc)
	if err != nil {
		return Conf{}, err
	}
	return Conf{
		Init: conf,
	}, nil
}
