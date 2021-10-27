## Simple PBFT Blockchain

> 复现论文  “Performance Optimization for Blockchain-Enabled Industrial Internet of Things (IIoT) Systems: A Deep Reinforcement Learning Approach” 的区块链部分。

## 任务

完成本论文中的网络部分，包括：区块链节点和网络的构建、PBFT共识协议的实现与区块链相结合、动态调整区块链中的参数配置。

通过DQN动态调节区块链中的出块时间、区块大小、共识协议、出块节点。



### DQN模型

state space：

* Transaction size
* stake
* 地理位置 (x,y)
* cpu性能
* 传输速率

action space：

* 出块节点
* 共识协议
* 区块大小
* 出块时间

target = max(吞吐量)





参考:

[论文分析](https://kid1999.github.io/2021/10/24/IIOT%E5%8C%BA%E5%9D%97%E9%93%BE%E7%9A%84%E5%BC%BA%E5%8C%96%E5%AD%A6%E4%B9%A0%E4%BC%98%E5%8C%96/)

[Practical Byzantine Fault Tolerance 论文](https://dblp.org/rec/conf/osdi/CastroL99)

