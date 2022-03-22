package main

import (
	"log"

	"github.com/XiaoXinTao/goproject/item_rpc/db"
	item_rpc "github.com/XiaoXinTao/goproject/item_rpc/kitex_gen/xintao/project/item_rpc/ecommerceitemrpc"
	"github.com/XiaoXinTao/goproject/item_rpc/logger"
)

func main() {
	svr := item_rpc.NewServer(new(EcommerceItemRpcImpl))
	logger.InitKlog()
	db.InitMysql()
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
