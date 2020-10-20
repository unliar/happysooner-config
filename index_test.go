package happyConfig

import (
	"errors"
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	"testing"
)

var c Conf

func init() {
	fmt.Println("init")
	conf, _ := config.NewConfig()
	_ = conf.Load(
		file.NewSource(
			file.WithPath("./examples/data.json"),
		),
	)
	c.Init = conf
}

func TestConf_GetMySQL(t *testing.T) {
	r, err := c.GetMySQL("mysql")
	if err != nil {
		t.Error(err)
	}
	if r.Port != 3316 {
		t.Error(errors.New("mysql 错误的端口读取"))
	}
}

func TestConf_GetSMTP(t *testing.T) {
	r, err := c.GetSMTP("smtp")
	if err != nil {
		t.Error(err)
	}
	if r.Password != "1234" {
		t.Error(errors.New("smtp 错误的密码读取"))
	}
}

func TestConf_GetRedis(t *testing.T) {
	r, err := c.GetRedis("redis")
	if err != nil {
		t.Error(err)
	}
	if r.Port != 6379 {
		t.Error(errors.New("redis 错误的端口读取"))
	}
}
