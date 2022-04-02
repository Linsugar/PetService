#拉取基础镜像
FROM golang:1.17.3-alpine

#作者名字 唐帝
MAINTAINER Tangdi

#要进行工作的目录-为当前下的PetService
WORKDIR /PetService

#设置打包为linux环境
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY="https://goproxy.cn,direct"


ADD . .

#对依赖进行更新下载
RUN  go mod download

#COPY --from=build-nev /server /
#暴露端口8000
EXPOSE 8000

#执行镜像的命令
ENTRYPOINT ["./mygo"]

#进行build
RUN go build -o /server







