# 如何捕获Ctrl-C/Kill


简要来说就这3行：

```go
sigCh := make(chan os.Signal, 1)
signal.Notify(sigCh, os.Interrupt, os.Kill) // 能够捕获ctrl-c和kill
<-sigCh
```

当从`sigCh`读到数据时，说明收到了中断信号。写了两个Demo：

- control_c.go: 中断直接由子goroutine捕获，缺点main函数接不到ctrl-c无法退出。
- control_c_2.go: 中断由main捕获，main再通知子routine退出。
