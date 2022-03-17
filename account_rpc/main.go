package main

import (
	"log"

	account_rpc "project.ecommerce.account_rpc/kitex_gen/project/ecommerce/account_rpc/ecommerceaccountrpc"
	"project.ecommerce.account_rpc/logger"
)

func main() {
	svr := account_rpc.NewServer(new(EcommerceAccountRpcImpl))
	logger.InitKlog()

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
