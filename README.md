# 交配引擎
## 目录结构
├── conf                     # 配置文件存放目录
│   ├── config.yaml          # 配置文件
├── engine                   # 引擎包
│   ├── init.go              # 初始化
│   ├── order.go             # 委托单
│   ├── order_book.go        # 交易委托账本
│   ├── order_queue.go       # 订单队列
│   ├── run.go               # 具体交易对的撮合引擎启动入口
│   └── trade.go             # 成交记录
├── enum                     # 枚举类型的包
│   ├── order_action.go      # 订单行为，create为下单，cancel为撤单
│   ├── order_side.go        # 买卖方向，buy为买入，sell为卖出
│   ├── order_type.go        # 订单类型，MVP版本(1.0版本)只支持limit
│   └── sort_direction.go    # 排序方向，asc为升序，desc为降序
├── errcode                  #
│   ├── code.go              # 定义了各种不同的错误码
│   └── errcode.go           # 错误码的数据结构，包括code和msg两个字段
├── handler                  #
│   ├── close_matching.go    # 接收关闭撮合的请求
│   ├── handle_order.go      # 接收处理订单的请求
│   └── open_matching.go     # 接收开启撮合的请求
├── main.go                  # Go程序唯一入口
├── middleware               # 中间件的包
│   ├── cache                # 缓存包
│   │   └── cache.go         # 缓存操作
│   ├── mq                   # 消息队列包
│   │   └── mq.go            # MQ操作
│   └── redis.go             # 主要做Redis初始化操作
└── process                  #
    ├── close_engine.go      # 关闭引擎
    ├── dispatch.go          # 分发订单
    ├── init.go              # 初始化
    └── new_engine.go        # 启动新引擎