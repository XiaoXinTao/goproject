module project.ecommerce.account_rpc

go 1.16

require (
	github.com/apache/thrift v0.16.0
	github.com/cloudwego/kitex v0.2.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.18.1 // indirect
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.2
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0
