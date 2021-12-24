# localcache
基于go语言实现的localcache

## v1.0
参考go-cache实现
- 数据结构
map[string]interface{}
- 并发安全
用go提供的读写锁
- 高性能并发
暂时不支持
- 对象上限
暂不支持
- 淘汰策略
暂不支持
- 过期清除
使用go提供的定时器
- gc
没有gc机制

## v2.0
- 数据结构
  map[string]interface{}
- 并发安全
  用go提供的读写锁
- 高性能并发
  暂不支持
- 对象上限
  由构造函数提供对象上限
- 淘汰策略
  lru算法
- 过期清除
  暂不支持
- gc
  没有gc机制

## v3.0
- 数据结构
  list + map
- 并发安全
  用go提供的读写锁
- 高性能并发
  使用哈希桶来实现
- 对象上限
  由构造函数提供对象上限
- 淘汰策略
  lru算法
- 过期清除
  暂不支持
- gc
  没有gc机制