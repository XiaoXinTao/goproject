package main

import (
	"log"

	"github.com/XiaoXinTao/goproject/account_rpc/dal"
	"github.com/XiaoXinTao/goproject/account_rpc/db"
	account_rpc "github.com/XiaoXinTao/goproject/account_rpc/kitex_gen/project/ecommerce/account_rpc/ecommerceaccountrpc"
	"github.com/XiaoXinTao/goproject/account_rpc/logger"
)

func main() {
	svr := account_rpc.NewServer(new(EcommerceAccountRpcImpl))
	logger.InitKlog()
	dal.InitRedis()
	db.InitMysql()
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
