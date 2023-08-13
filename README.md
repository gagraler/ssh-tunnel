<h2 align="center">SSH Tunnel</h2>
SSH Tunnel 是一个简单的工具，用于通过 SSH 将本地端口转发到远程服务器。适用于以本地方式访问远程服务器上运行的服务，可在只通过 SSH 端口访问远程服务器的情况下，访问远程服务器上的服务

## 运行
1. 克隆此仓库：
```shell
   git clone https://github.com/your-username/ssh-tunnel.git
   
   cd ssh-tunnel
   
    go mod tidy
```

2. 编译：
```shell

    # mac amd64
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ssh-tunnel-darwin -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-w -s" -v .
    
    # mac arm64
    CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ssh-tunnel-darwin -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-w -s" -v .
    
    # linux amd64
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ssh-tunnel-linux -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-w -s" -v .
    
    # linux arm64
    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ssh-tunnel-arm64 -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-w -s" -v .
    
     # windows amd64
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ssh-tunnel-windows -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-w -s" -v .
    
```

## License
[MIT License](https://github.com/keington/ssh-tunnel/blob/4ce3c32da7e9b60921e83e83f0011e7d0bdc1090/LICENSE)

## Copyright
Copyright (c) 2023 许怀安
