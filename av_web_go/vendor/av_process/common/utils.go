package common

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os/exec"
	"strings"
)

// GetLocalIPAddress 获取本机 IP 地址
func GetLocalIPAddress() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		Logger.Fatal(err)
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	defer func() {
		if err := conn.Close(); err != nil {
			Logger.Fatal(err)
		}
	}()
	localIPAddress := localAddr.IP
	return localIPAddress
}

// ExecuteCommand 执行命令
func ExecuteCommand(command string) (executeResult string) {
	Logger.Infof("当前执行的命令: %s", command)
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	// 标准输出管道
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		Logger.Errorf("创建命令输出管道失败, 错误原因: %v", err)
		return err.Error()
	}
	if err := cmd.Start(); err != nil {
		Logger.Errorf("执行命令: %v", err)
		return err.Error()
	}
	stdOutBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		Logger.Errorf("输出流错误原因: %v", err)
		return err.Error()
	}
	if err := cmd.Wait(); err != nil {
		Logger.Errorf("等待执行结果: %v", err)
		return err.Error()
	}
	return fmt.Sprintf("%s", stdOutBytes)
}

// GenRandomString 生成
func GenRandomString() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// GenMD5 获取 MD5 字符串
func GenMD5(input string) string {
	input = input + "av_web+"
	out := md5.Sum([]byte(input))
	return fmt.Sprintf("%x", out)
}
