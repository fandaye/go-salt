package main

import (
	"github.com/fandaye/go-salt"
	"fmt"
	"log"
)

func main()  {
	//定义配置文件
	Config := make(map[string]string)
	Config["salt_addr"]="10.100.4.36"
	Config["salt_prot"]="8080"
	Config["salt_user"]="salt"
	Config["salt_passwd"]="salt"

	S:=go_salt.Salt{}
	S.Config=Config
	
	//saltstack api 判断文件是否存在
	post_cmd_1 := fmt.Sprintf(`{"fun": "%s", "client": "%s", "tgt": "%s" ,"arg": "/etc/passwd"}`, "file.file_exists", "local", "host01")
	if PostData, Err := S.CMD_SALT(post_cmd_1); Err == nil {
		fmt.Println(PostData)
	}else {
		log.Println(Err.Error())
	}

	//saltstack api 执行命令
	post_cmd_2 := fmt.Sprintf(`{"fun": "%s", "client": "%s", "tgt": "%s" ,"arg": "%s"}`, "cmd.run", "local", "host01","free -m")
	if PostData, Err := S.CMD_SALT(post_cmd_2); Err == nil {
		fmt.Println(PostData)
	}else {
		log.Println(Err.Error())
	}
}


