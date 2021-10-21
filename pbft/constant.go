package pbft

/**
* @Description: 常量
* @author : kid1999
* @date Date : 2021/10/21 15:19
* @version V1.0
 */

const (
	//客户端的监听地址
	clientAddr = "127.0.0.1:8888"
	//节点总数
	nodeCount = 4
)

//节点池，主要用来存储监听地址
var NodeTable = map[string]string{
	"N0": "127.0.0.1:8000",
	"N1": "127.0.0.1:8001",
	"N2": "127.0.0.1:8002",
	"N3": "127.0.0.1:8003",
}
