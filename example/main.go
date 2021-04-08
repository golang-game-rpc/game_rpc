package main

import (
	"encoding/json"
	"fmt"
	"game_rpc"
)

type SPQueryRoleDetail struct {
	RoleID uint `json:"RoleID"`
}

func main() {
	url := "192.168.4.120:12245"
	data, _ := json.Marshal(SPQueryRoleDetail{RoleID: 110725221648588646})

	message := &game_rpc.Message{
		Api:      "SPQueryPlayerDetailsReq",
		Data:     string(data),
		ServerID: -1,
	}

	// 编码、发送socket、解码
	resp, err := game_rpc.SendSocket(url, message)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(*resp.Result)
	fmt.Println(*resp.Data)
}
