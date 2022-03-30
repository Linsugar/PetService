FROM golang:1.17.3-alpine

MAINTAINER Tangdi

WORKDIR /PetService

ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY="https://goproxy.cn,direct"

ADD . .


RUN  go mod download

#COPY --from=build-nev /server /
EXPOSE 8000

ENTRYPOINT ["./mygo"]


RUN go build -o /server







