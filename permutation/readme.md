# 几种计算全排列的方法

基于golang实现，有非并发实现和并发实现

## 递归

全排列问题，比如打印1-9的共9个字母的全排列。先输出1，然后是2-9的全排列，输出2，然后1-9中去除2的全排列。于是很自然想到递归的方法。写出伪代码：

```

permutaion(prefix, set) {
    if set 为空
        print prefix

    for s in set:
        permuetaion(prefix+s, set - s)
}
```

## go递归实现

```
func permutaionImpl(prefix []byte, s []byte, cur int) {
	if cur == len(s) {
		fmt.Println(prefix)
		return
	}

	for _, b := range s {
		exist := false
		for i := 0; i < cur; i++ {
			if prefix[i] == b {
				exist = true
				break
			}
		}

		if !exist {
			prefix[cur] = b
			permutaionImpl(prefix, s, cur+1)
		}

	}
}

func Permutation(s []byte) {
	//前缀slice
	p := make([]byte, len(s), len(s))
	permutaionImpl(p, s, 0)
}
```

## 并发1

刚才实现的是go单线程执行的，改成并发版本的, 参考rob pike讲[Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide#39)，多个goroutine通过channel连接起来。goroutine1发送1给goroutine2, goroutine2发送12给goroutine3，goroutine3发送123给goroutine4， 以此类推。

``` go
func prefixIncrement(in []byte, s []byte, next chan []byte) {
	for _, c := range s {
		exist := false
		for _, e := range in {
			if e == c {
				exist = true
				break
			}
		}
		if exist {
			continue
		}

		temp := make([]byte, 0)
		temp = append(temp, in...)
		temp = append(temp, c)
		next <- temp
	}
}

func permutaionConImpl(req chan []byte, out chan []byte, s []byte) {
	go func() {
		//递归退出条件: len(v) == len(s)-1
		v, ok := <-req
		if !ok {
			return
		}

		next := out
		if len(v) != len(s)-1 {
			next = make(chan []byte)
			permutaionConImpl(next, out, s)
		}

		prefixIncrement(v, s, next)
		for in := range req {
			prefixIncrement(in, s, next)
		}
		close(next)
	}()
}

// PermutationConcurrency  并发计算全排列
func PermutationConcurrency(s []byte) {
	req, out := make(chan []byte), make(chan []byte)

	//开启goroutine计算
	permutaionConImpl(req, out, s)

    over := make(chan struct{})
    
    //要开goroutine读取out，如果放在主函数中，会导致死锁。
	go func() {
		for res := range out {
			fmt.Println(res)
		}
		close(over)
	}()

	for _, c := range s {
		sl := []byte{c}
		req <- sl
	}
	close(req)

	<-over

}
```

## 并发2

并发1中把每个排列的阶段拆分到不同的goroutine， 从goroutine1开始每个goroutine越来越繁忙，最后一个goroutine要输出n!个slice，任务轻重很不均衡。于是也可以考虑让每个goroutine做同样的事情，比如下面的实现

```
func PermutationConcurrencyVertical(s []byte) {
	var sg sync.WaitGroup
	for _, c := range s {
		sg.Add(1)

		//要传参数进去，避免只取到最后一个值
		go func(k byte) {
			p := make([]byte, len(s), len(s))
			p[0] = k
			permutaionImpl(p, s, 1)
			defer sg.Done()
		}(c)
	}

	sg.Wait()
}
```

## 并发3

如果能提前知道要开几个goroutine，那就可以不用递归的方式创建goroutine了，代码逻辑会更清晰易懂。

```go
func permutaionCal(left chan []byte, right chan []byte, s []byte) {
	for v := range left {
		prefixIncrement(v, s, right)
	}
	close(right)
}

func PermutaionChain(s []byte) {
	leftmost := make(chan []byte)
	left := leftmost
	right := leftmost

	for i := 0; i < len(s)-1; i++ {
		right = make(chan []byte)
		go permutaionCal(left, right, s)
		left = right
	}

	go func(c chan []byte) {
		for _, v := range s {
			c <- []byte{v}
		}
		close(c)
	}(leftmost)

	for v := range right {
		fmt.Println(v)
	}
}
```



## 代码地址

https://github.com/amither/go-slice/tree/master/permutation

