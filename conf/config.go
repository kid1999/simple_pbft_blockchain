package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/**
* @Description:  TODO 动态读取配置
* @author : kid1999
* @date Date : 2021/10/27 7:22 PM
* @version V1.0
 */
type Location struct {
	x float32
	y float32
}

type Config struct {
	NodeID        string
	BlockSize     int
	BlockInterval int
	Consensus     int
	Producers     []string
	NodeCount     int
	LeaderID      string
	Stake         int
	Location
	CPU int //MHz
	// 地理位置，权益的基尼系数阈值
	StakeMAX        float32
	LocationMAX     float32
	TransactionSize int
}

var GlobalConfig *Config

//节点池，主要用来存储监听地址
var NodeTable = map[string]string{
	"N0": "127.0.0.1:8000",
	"N1": "127.0.0.1:8001",
	"N2": "127.0.0.1:8002",
	"N3": "127.0.0.1:8003",
}

var ClientTable = map[string]string{
	"N0": "127.0.0.1:9000",
	"N1": "127.0.0.1:9001",
	"N2": "127.0.0.1:9002",
	"N3": "127.0.0.1:9003",
}

var ConfigServer = "http://127.0.0.1:5000/api"

func NewConfig() *Config {
	c := Config{}
	c.NodeID = "N0"
	c.BlockSize = 2
	c.BlockInterval = 2
	c.Consensus = 0
	c.Producers = []string{"N0", "N1", "N2"}
	c.LeaderID = "N0"
	c.NodeCount = len(c.Producers)
	c.CPU = 100
	c.Stake = 20
	c.Location = Location{0.5, 0.5}
	c.LocationMAX = 0.3
	c.StakeMAX = 0.1
	c.TransactionSize = 10
	GlobalConfig = &c
	return &c
}

func (c *Config) PlanToGetConfig() {
	for {
		c.RequestConfig()
		time.Sleep(time.Second * 10)
	}
}

func (c *Config) RequestConfig() {
	resp, err := http.Get(ConfigServer)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	json.Unmarshal(body, c)
}

func (c Config) IsProducer() bool {
	for i := 0; i < len(c.Producers); i++ {
		if c.NodeID == c.Producers[i] {
			return true
		}
	}
	return false
}
