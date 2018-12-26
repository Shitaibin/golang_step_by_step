一招教你无阻塞读写通道
----------------

### 文件介绍
```
├── block
│   ├── read_buffer_block.go
│   ├── read_unbuffer_block.go
│   ├── write_buufer_block.go
│   └── write_unbuffer_block.go
├── select_timer_unblock.go
└── select_unblock.go
```

- `block`：目录，4种情况的阻塞示例代码
- `select_unblock.go`：使用select的default语句实现无阻塞读写
- `select_timer_unblock.go`：使用select和超时实现无阻塞读写


正文
-----------------

