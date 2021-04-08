package game_rpc

import (
	"errors"
	"net"
	"time"
)

func SendSocket(url string, req *Message) (*SPGMRes, error) {
	// 建立socket连接
	conn, err := net.Dial("TCP", url)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// 编译
	data, err := packData(req)
	if err != nil {
		return nil, err
	}

	// 设置超时时间
	err = conn.SetDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		return nil, err
	}

	// 发送数据
	_, err = conn.Write(data)
	if err != nil {
		return nil, err
	}

	// 接受数据，缓冲区必须要分配内存
	var buf [4096]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		return nil, err
	}

	// 回包不为空，进行解码
	if n <= 0 {
		return nil, errors.New("server return null")
	}

	resp, err := unpackData(buf[:n])
	if err != nil {
		return nil, err
	}
	return resp, nil
}
