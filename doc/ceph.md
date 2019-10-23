# Ceph

Ceph 是一个分布式文件存储系统，由 Sage Weil 在 UCSC 提出。

## Ceph 的组成

- 存储管理器 Ceph OSD
    - Ceph OSD 守护进程的作用是存储数据，处理数据的复制，恢复，回填，再均衡，并检查其他 OSD 守护进程的心跳来向 Ceph Monitor 来提供监控信息
    - Ceph 集群默认有 3 个 OSD，但至少要有 2 个才能保持集群的状态为 active + clean
- 集群监视器 Ceph Monitor
    - CM 用来维护展示集群状态的各种图表
- 元数据服务器 MDS
    - MDS 为 Ceph 文件存储系统存储元数据

## Ceph 存储集群的组成

- 一系列的节点
- 存储设备
- 传输网络

## 相关概念

- 文件存储：适用于文件系统，读写慢，利于共享
- 块存储：适用于硬盘，读写快，不利于共享
- 对象存储：适用于键值数据库，读写快，利于共享
- iSCSI 存储协议：Internet Small Computer System Interface，是 IEFT 在 2003 年发布的在 TCP/IP 上进行数据块传输的存储协议（标准）。
    - iSCSI 协议的负载数据是 SCSI 命令，传输是基于 TCP/IP 协议。参照 OSI 七层网络模型，iSCSI 位于第五层会话层，将第七层的 SCSI 协议包封装映射到第四层的 TCP 协议包在以太网上进行传输。SCSI 协议是一种客户端/服务器模式的协议，iSCSI 在协议封装上大量沿用了 SCSI 命令的语义，以及 Initiator 和 Target 的概念。在一个交互会话中，会话的发起方称为 Initiator，而服务方称为 Target。IQN（iSCSI Qualified Name）则是用来标识 Initiator 和 Target 身份的唯一标识符。
    - 工作流程：
        - 在 iSCSI 系统中，由 SCSI 适配器发送一条 SCSI 命令，并将命令封装到 TCP/IP 包中，传递到以太网络中
        - 接收者收到 TCP/IP 包后抽取出其中的 SCSI 命令并执行相应的操作，将返回的 SCSI 命令和数据封装到 TCP/IP 包中，发回给发送者
        - 系统收到返回的 TCP/IP 包后提取出数据或者命令，并把它们传回 SCSI 子系统

## CRUSH 算法

> Controlled Replication Under Scalable Hashing

该算法的主要目的在于为给定的 pg 分配一组存储数据的 osd 节点，即用于控制数据分布的一种方法。


## References

- [http://docs.ceph.org.cn/](http://docs.ceph.org.cn/)
- [https://zhuanlan.zhihu.com/p/34188712](https://zhuanlan.zhihu.com/p/34188712)
- [https://blog.csdn.net/yuxin6866/article/details/79451006](https://blog.csdn.net/yuxin6866/article/details/79451006)
- [http://www.xuxiaopang.com/2016/11/08/easy-ceph-CRUSH/](http://www.xuxiaopang.com/2016/11/08/easy-ceph-CRUSH/)

