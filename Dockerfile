#
# 编译应用
#
FROM registry.cn-shanghai.aliyuncs.com/xm69/go-build AS build

ENV GOPROXY=https://goproxy.cn
COPY ./go.mod ./go.sum ./
RUN set -eux && \
    # 打印环境
    echo "环境：$(uname -a)" && \
    # 加载依赖
    go mod download && go mod verify

COPY . .
RUN set -eux && \
  #编译
  go build -tags musl -o /home/main

################

#
# 封装镜像
#
FROM alpine:3
WORKDIR /home
RUN set -eux && \
  #设置源
  echo "http://mirrors.ustc.edu.cn/alpine/edge/main/" > /etc/apk/repositories && \
  echo "http://mirrors.ustc.edu.cn/alpine/edge/community/" >> /etc/apk/repositories && \
  echo "http://mirrors.ustc.edu.cn/alpine/edge/testing/" >> /etc/apk/repositories && \
  apk update && \
  \
  #设置时区
  apk add tzdata && \
  cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
  echo "Asia/Shanghai" > /etc/timezone && \
  apk del tzdata

#从编译应用阶段中复制程序
COPY --from=build /home/main .

ENTRYPOINT ["./main"]