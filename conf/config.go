package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
* @Description:  TODO 动态读取配置
* @author : kid1999
* @date Date : 2021/10/27 7:22 PM
* @version V1.0
 */
type Node struct {
	ID         string `json:"id"`
	NodeAddr   string `json:"node_addr"`
	ClientAddr string `json:"client_addr"`
	// 计算性能 MHz
	CPU int `json:"cpu"`
	// 存储能力 MB
	Disk int `json:"disk"`
}

type Config struct {
	NodeID        string
	BlockSize     int
	BlockInterval int
	Consensus     int
	Producers     []string
	NodeCount     int
	LeaderID      string
	DBPath        string
}

var GlobalConfig *Config

//节点池，主要用来存储监听地址
var Nodes = []Node{
	{ID: "N0", NodeAddr: "127.0.0.1:8000", ClientAddr: "127.0.0.1:9000", CPU: 10 * 1024 * 1024, Disk: 1000 * 1024},
	{ID: "N1", NodeAddr: "127.0.0.1:8001", ClientAddr: "127.0.0.1:9001", CPU: 5 * 1024 * 1024, Disk: 500 * 1024},
	{ID: "N2", NodeAddr: "127.0.0.1:8002", ClientAddr: "127.0.0.1:9002", CPU: 3 * 1024 * 1024, Disk: 300 * 1024},
	{ID: "N3", NodeAddr: "127.0.0.1:8003", ClientAddr: "127.0.0.1:9003", CPU: 1 * 1024 * 1024, Disk: 100 * 1024},
}

var ConfigServer = "http://127.0.0.1:5000/api"

// 读取配置信息
func NewConfig() *Config {
	c := Config{}
	c.NodeID = "N0"
	c.BlockSize = 2
	c.BlockInterval = 2
	c.Consensus = 0
	c.Producers = []string{"N0", "N1", "N2"}
	c.LeaderID = "N0"
	c.NodeCount = len(c.Producers)
	c.DBPath = "Keys/db/my.db"
	GlobalConfig = &c
	return &c
}

// 上传配置信息
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

func GetNode(ID string) *Node {
	for _, node := range Nodes {
		if node.ID == ID {
			return &node
		}
	}
	return nil
}
