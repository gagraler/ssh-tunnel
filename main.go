package main

import (
	"dev-proxy/config"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
)

func init() {
	config.InitConfig()
}

var (
	// SSH服务器地址
	sshServer = config.GetConfig("server")
	// SSH服务器端口
	sshPort = config.GetConfig("port")
	// SSH用户
	sshUser = config.GetConfig("user")
	// SSH密码
	sshPassword = config.GetConfig("keington")
)

func main() {

	// 远程服务器开放的端口列表
	remotePorts := map[int]map[string]int{
		3306:  {"MySQL": 3306},
		8080:  {"Web": 8080},
		15080: {"After End": 15080},
	}

	// 建立SSH连接
	sshConfig := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	sshConn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", sshServer, sshPort), sshConfig)
	if err != nil {
		log.Fatalf("Failed to dial SSH server: %s", err)
	}
	defer func(sshConn *ssh.Client) {
		err := sshConn.Close()
		if err != nil {
			return
		}
	}(sshConn)

	// 启动端口转发监听
	for remotePort, info := range remotePorts {
		for serviceName, localPort := range info {
			go func(remotePort int, localPort int, serviceName string) {
				listenAndServe(localPort, "localhost", remotePort, sshConn, serviceName)
			}(remotePort, localPort, serviceName)
		}
	}

	// 打印访问地址
	fmt.Println("Forwarding ports:")
	for remotePort, info := range remotePorts {
		for serviceName, localPort := range info {
			fmt.Printf("\t- localhost:%d -> %s:%d [%s]\n", localPort, "remote host", remotePort, serviceName)
		}
	}

	// 等待程序退出
	select {}
}

// 复制输入流到输出流
func copyStream(src net.Conn, dst net.Conn, errChan chan error) {
	_, err := io.Copy(dst, src)
	if err != nil {
		errChan <- err
	}
}

// 启动端口隧道监听
func listenAndServe(localPort int, remoteHost string, remotePort int, sshConn *ssh.Client, serviceName string) {
	localAddr := fmt.Sprintf("127.0.0.1:%d", localPort)
	remoteAddr := fmt.Sprintf("%s:%d", remoteHost, remotePort)
	listener, err := net.Listen("tcp", localAddr)
	if err != nil {
		log.Fatalf("Failed to start local listener: %s", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Failed to accept incoming connection: %s", err)
		}
		remoteConn, err := sshConn.Dial("tcp", remoteAddr)
		if err != nil {
			log.Fatalf("Failed to dial remote server: %s", err)
		}
		go func() {
			defer func(conn net.Conn) {
				err := conn.Close()
				if err != nil {
					return
				}
			}(conn)
			defer func(remoteConn net.Conn) {
				err := remoteConn.Close()
				if err != nil {
					return
				}
			}(remoteConn)
			errChan := make(chan error, 1)
			go copyStream(conn, remoteConn, errChan)
			go copyStream(remoteConn, conn, errChan)
			<-errChan
		}()
	}
}
