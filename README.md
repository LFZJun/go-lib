[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/ljun20160606/go-lib/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/ljun20160606/go-lib.svg?branch=master)](https://travis-ci.org/ljun20160606/go-lib)
[![codecov](https://codecov.io/gh/ljun20160606/go-lib/branch/master/graph/badge.svg)](https://codecov.io/gh/ljun20160606/go-lib)

# go-lib

## 目录

1. [设计模式](gof)
    * [策略模式](gof/strategy)
    * [单例模式](gof/singleton)
    * [装饰者模式](gof/decorator)
    * [工厂模式](gof/factory)
    * [命令模式](gof/command)
    * [适配器模式](gof/adapter)
1. [cache](cache)
    * [timing wheel](cache/timer/timing-wheel.go)
    * [ccache](cache/ccache.go)
1. [池](pool)
    * [连接池](pool/conncet-pool.go) (不能存map，因为复用的时候不会删除entry)
    * [协程池](pool/coroutine-pool.go) (只实现了size，没实现min, max)
1. [分布式锁](lock/redis_mutex.go) (redis实现)
1. [编译原理前端](parser)
    * [计算器](parser/calculator/)
    * [go tag](parser/tag/)
1. [go slice通用方法](slice/)
    * [filter](slice/filter.go)
    * [groupby](slice/groupby.go)
    * [shuffle](slice/shuffle.go)
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

## 说明

公共方法库