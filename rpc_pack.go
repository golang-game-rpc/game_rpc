package game_rpc

import (
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
)

const HeadUsedLen = 1
const DataUsedLen = 4

type Message struct {
	Api      string
	Data     string
	ServerID int32
}

// 打包二进制
func packData(message *Message) ([]byte, error) {
	// 头部信息
	head := &CSPkgHead{
		CmdID:               proto.Uint32(100),
		MsgSeqID:            proto.Uint32(101),
		NotifyMsgSeqID:      nil,
		EncryptCompressType: nil,
	}
	// 进行编码
	headProtobuf, err := proto.Marshal(head)
	if err != nil {
		return nil, err
	}

	// 请求体
	body := &SPGMReq{
		Api:      proto.String(message.Api),
		Data:     proto.String(message.Data),
		ServerID: proto.Int32(message.ServerID),
	}
	// 进行编码
	bodyProtobuf, err := proto.Marshal(body)
	if err != nil {
		return nil, err
	}

	// 计算二进制参数个长度
	headLen := len(headProtobuf)
	bodyLen := len(bodyProtobuf)
	dataLen := HeadUsedLen + headLen + bodyLen
	dataLenBytes, err := intToBytes(dataLen)
	if err != nil {
		return nil, err
	}
	// 组合
	data := string(dataLenBytes) + string(headLen) + string(headProtobuf) + string(bodyProtobuf)

	return []byte(data), nil
}

// 解包二进制
func unpackData(buff []byte) (*SPGMRes, error) {
	// 存放dataLen占 4
	dataLen, err := bytesToInt(buff[:DataUsedLen])
	if err != nil {
		return nil, err
	}
	bodyStart := DataUsedLen + HeadUsedLen + int(buff[DataUsedLen])
	bodyEnd := dataLen + DataUsedLen
	// 声明返回结果的指针
	resp := &SPGMRes{}
	// 解码
	err = proto.Unmarshal(buff[bodyStart:bodyEnd], resp)
	if err != nil {
		return nil, err
	}
	// 打印
	return resp, nil
}

// 整形转字节数组
func intToBytes(data int) ([]byte, error) {
	tmp := int32(data)
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, tmp)
	if err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

// 字节转换成整形
func bytesToInt(data []byte) (int, error) {
	var tmp int32
	bytesBuffer := bytes.NewBuffer(data)
	err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	if err != nil {
		return 0, err
	}
	return int(tmp), nil
}
