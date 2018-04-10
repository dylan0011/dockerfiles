# 配置文件说明
	- clientPort=2181 客户端连接 Zookeeper 服务器的端口，Zookeeper 会监听这个端口，接受客户端的访问请求

	- dataDir=/data 数据存放地址,默认情况下，Zookeeper将写数据的日志文件也保存在这个目录里

	- dataLogDir=/datalog 日志文件目录

	- tickTime=2000 服务器之间或客户端与服务器之间维持心跳的时间间隔(ms)

	- initLimit=5 集群中的follower服务器(F)与leader服务器(L)之间初始连接时能容忍的最多心跳数(tickTime的数量)

	- syncLimit=2 集群中的follower服务器与leader服务器之间请求和应答之间能容忍的最多心跳数(tickTime的数量）

	- minSessionTimeout 最小会话超时时间 默认情况下,minSession=2*tickTime

	- maxSessionTimeout 最大会话超时时间 默认情况下,maxSession=20*tickTime

	- maxClientCnxns=60  限制连接到Zookeeper的客户端数量，并限制并发连接的数量

	- server.N=YYY:A:B   
		- N: 服务器编号
		- YYY: 服务器的IP地址
		- A: LF通信端口(F:follower服务器，L:Leader服务器)，即 该服务器与集群中的leader交换的信息的端口
		- B: 选举端口，当leader挂掉后，选举新leader时服务器间相互通信的端口
		- 一般来说，集群中每个服务器的A端口都是一样，每个服务器的B端口也是一样。但是当所采用的为伪集群时，IP地址都一样，只能使A端口和B端口不一样
	- e.g.
		server.1=zoo1:2888:3888
		server.2=zoo2:2888:3888
		server.3=zoo3:2888:3888


# docker-compose.yml 参考 https://docs.docker.com/samples/library/zookeeper/
