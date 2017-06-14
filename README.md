### 随笔
1. [设计模式](gof)
    * [策略模式](gof/strategy)
    * [单例模式](gof/singleton)
    * [装饰者模式](gof/decorator)
    * [工厂模式](gof/factory)
    * [命令模式](gof/command)
    * [适配器模式](gof/adapter)
1. [缓存](cache/memoryCache.go) ([优化方法](./README.md#17))
1. [池](pool)
    * [连接池](pool/conncet-pool.go) (不能存map，因为复用的时候不会删除entry)
    * [协程池](pool/coroutine-pool.go) (只实现了size，没实现min, max，凑活用，hhh)
1. [分布式锁](lock/redis_mutex.go) (redis实现)
1. [环形队列](queue/RingQueue.go)
1. [ioc](ioc)
1. [parse tag](parser/tag/tag.go)
1. [计算器](parser/calculator/calculator.go)
1. [编辑距离](levenshtein/distance_test.go)

#### 缓存思考
1. 过期策略用的是多个timer，即定时删除策略。
2. 定时策略很尴尬，一个timer一个goroutine。当timer过多时会导致sched调度过多，cpu处理不过来，性能严重下降。
3. 优化方向是，改变定时策略，采用懒汉策略和定期删除策略。
> 即get时对比时间辍，删除过期数据。(其实也不是很懒hhh，有点像引用计数GC的即时触发)

> 同时有且只有一个定时器，定期遍历，删除过期数据。(hhh遍历这件事儿有点儿像标记清除，从根节点遍历) 这个定时器，有个致命点就是当数据过多时，遍历时间过长(这点没跑了太像标记清除了，最大暂停时间过长hhh)。

> 所以可以试试环形定时器，一个环像一个时钟，假设！一圈一个小时，一刻度1毫秒，一个刻度对应一个任务slot即任务队列。每次插入一个数据都是转换成对应的插槽，插入该任务队列。这样保证了每次定期操作不会有过多的无效操作，比如没有数据过期但是仍然遍历了全部数据。（这样太变态了不是么）
4. 数据存储方面我根本没有优化hhh，给祖国丢人了。
