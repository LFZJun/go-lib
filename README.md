## 目录
1. [设计模式](gof)
    * [策略模式](gof/strategy)
    * [单例模式](gof/singleton)
    * [装饰者模式](gof/decorator)
    * [工厂模式](gof/factory)
    * [命令模式](gof/command)
    * [适配器模式](gof/adapter)
1. [concurrentMap](cache/memoryCache.go) ([优化方法](./README.md#17))
1. [池](pool)
    * [连接池](pool/conncet-pool.go) (不能存map，因为复用的时候不会删除entry)
    * [协程池](pool/coroutine-pool.go) (只实现了size，没实现min, max，凑活用，hhh)
1. [分布式锁](lock/redis_mutex.go) (redis实现)
1. [ioc](ioc)
1. [编译原理前端](parser)
    * [计算器](parser/calculator/)
    * [go tag](parser/tag/)
1. [go slice通用方法](slice/)
1. [数据结构与算法](algorithms/)
    * [环形队列](algorithms/queue)
    * [搜索](algorithms/search)
        * [对分搜索](algorithms/search/binary.go)
        * [前缀树](algorithms/search/trie.go)
    * [栈](algorithms/stack)
    * [排序](algorithms/sort)
        * [归并排序](algorithms/sort/merge_sort.go)
        * [快速排序](algorithms/sort/quick_sort.go)
    * [编辑距离](algorithms/levenshtein/)
1. [leetcode](leetcode)
    * [1. Two Sum](leetcode/1.%20Two%20Sum.go)
    * [2. Add Two Numbers](leetcode/2.%20Add%20Two%20Numbers.go)
    * [3. Longest Substring Without Repeating Characters](leetcode/3.%20Longest%20Substring%20Without%20Repeating%20Characters.go)
    * [4. Median of Two Sorted Arrays](leetcode/4.%20Median%20of%20Two%20Sorted%20Arrays.go)
    * [5. Longest Palindromic Substring](leetcode/5.%20Longest%20Palindromic%20Substring.go)
    * [6. ZigZag Conversion](leetcode/6.%20ZigZag%20Conversion.go)
    * [7. Reverse Integer](leetcode/7.%20Reverse%20Integer.go)
    * [8. String to Integer (atoi)](leetcode/8.%20String%20to%20Integer%20(atoi).go)
    * [9. Palindrome Number](leetcode/9.%20Palindrome%20Number.go)
    * [10. Regular Expression Matching](leetcode/10.%20Regular%20Expression%20Matching.go)
    * [11. Container With Most Water](leetcode/11.%20Container%20With%20Most%20Water.go)
    * [72. Edit Distance](leetcode/72.%20Edit%20Distance.go)

## 思考
#### concurrentMap
1. 过期策略用的是多个timer，即定时删除策略。
2. 定时策略很尴尬，一个timer一个goroutine。当timer过多时会导致sched调度过多，性能严重下降。
3. 优化方向是，改变定时策略，采用懒汉策略和定期删除策略。
> 即get时对比时间辍，删除过期数据。(其实也不是很懒hhh，有点像引用计数GC的即时触发)

> 同时有且只有一个定时器，定期遍历，删除过期数据。(hhh遍历这件事儿有点儿像标记清除，从根节点遍历) 这个定时器，有个致命点就是当数据过多时，遍历时间过长(这点没跑了太像标记清除了，最大暂停时间过长hhh)。

> 所以可以试试环形定时器，一个环像一个时钟，假设！一圈一个小时，一刻度1毫秒，一个刻度对应一个任务slot即任务队列。每次插入一个数据都是转换成对应的插槽，插入该任务队列。这样保证了每次定期操作不会有过多的无效操作，比如没有数据过期但是仍然遍历了全部数据。（这样太变态了不是么）
4. 数据存储方面我根本没有优化hhh，给祖国丢人了。
