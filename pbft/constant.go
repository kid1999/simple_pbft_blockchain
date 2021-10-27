package pbft

/**
* @Description: 常量
* @author : kid1999
* @date Date : 2021/10/21 15:19
* @version V1.0
 */

const (
	//客户端的监听地址
	ClientAddr = "127.0.0.1:9999"
	//节点总数
	nodeCount = 4
	// 主节点信息
	LeaderID = "N0"
	// 区块包含的交易数
	BLOCK_SIZE = 2
	// 出块间隔
	BLOCK_INTERVAL = 10
	// 交易大小
	TransactionSize = 200
	// 节点地理信息
	X, Y = 500, 500
	// 节点权益
	Stake = 50
	// 节点计算资源
	CPU = 1
	// 地理位置，权益的基尼系数阈值
	StakeMAX, LocationMAX = 0.1, 0.3
)

const (
	ReplyTimeout  = 10
	LeaderTimeout = 500
)

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
