package main

import (
	"log"

	"account_rpc/dal"
	account_rpc "account_rpc/kitex_gen/project/ecommerce/account_rpc/ecommerceaccountrpc"
	"account_rpc/logger"
)

func main() {
	svr := account_rpc.NewServer(new(EcommerceAccountRpcImpl))

	dal.InitRedis()
	logger.InitKlog()
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
