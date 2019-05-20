#!/usr/bin bash
BINARYNAME="ipservice"
# 根据当前目录的Makefile创建二进制文件并进行服务启动
rm $BINARYNAME
make build

# 启动服务
./$BINARYNAME &

# check service avaiable
RESPONSE=`curl -Ss "http://localhost:8080/ping"`
if [ "$RESPONSE" != "pong" ];then
    #echo "NOT EQUAL"
    exit 1
fi

