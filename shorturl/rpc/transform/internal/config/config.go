package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf                 // 默认配置
	DataSource         string          // 手动代码
	Table              string          // 手动代码
	Cache              cache.CacheConf // 手动代码
}
