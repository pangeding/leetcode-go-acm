# Go "comma ok" 惯用法完整介绍

Go 中有 **三种** 场景使用 comma ok 语法。

---

## 1. Map 访问

### 语法

```go
value, ok := myMap[key]
```

### 含义

`myMap[key]` 返回两个值：

| 返回值 | 类型 | 含义 |
|--------|------|------|
| `value` | 任意 | key 对应的 value |
| `ok` | `bool` | key 是否存在于 map 中 |

- key **存在**：`ok` 为 `true`，`value` 是对应的值
- key **不存在**：`ok` 为 `false`，`value` 是该类型的零值

### 为什么需要 ok？

Go 中访问不存在的 key 不会报错，而是返回零值。如果只写一个返回值，就无法区分 "key 不存在" 和 "key 存在但 value 是零值" 这两种情况。

```go
m := map[string]int{"a": 0}

v1 := m["a"]   // v1 = 0, 但不知道是存在还是不存在
v2 := m["b"]   // v2 = 0, 同样无法判断

// 使用 comma ok
v1, ok1 := m["a"]  // v1 = 0,  ok1 = true  (key 存在)
v2, ok2 := m["b"]  // v2 = 0,  ok2 = false (key 不存在)
```

### 示例

```go
func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.cache[key]; ok {
        // key 存在，更新值并移到头部
        node.value = value
        this.moveToHead(node)
        return
    }
    // key 不存在，创建新节点
    // ...
}
```

### 常见用法

```go
// 检查 key 是否存在
if _, ok := m[key]; ok {
    // key 存在时执行
}

// 删除指定 key
delete(m, key)

// 安全取值
if val, ok := m[key]; ok {
    // 使用 val
}
```

---

## 2. 类型断言（Type Assertion）

### 语法

```go
value, ok := interface{}.(Type)
```

### 含义

将一个 `interface{}` 转换为具体类型时：

| 结果 | 含义 |
|------|------|
| `ok` 为 `true` | 转换成功，`value` 是目标类型的值 |
| `ok` 为 `false` | 转换失败，`value` 是目标类型的零值，**不会 panic** |

### 对比：不用 comma ok 的情况

如果只用单返回值，类型不匹配时会 **panic**（程序崩溃）：

```go
var i interface{} = "hello"

// 安全写法，不会 panic
s, ok := i.(string)      // s = "hello", ok = true
n, ok := i.(int)         // n = 0, ok = false (n 是 int 的零值)

// 危险写法，类型不匹配会 panic
s := i.(string)          // 正常，s = "hello"
n := i.(int)             // panic: interface conversion: interface {} is string, not int
```

### 示例

```go
func describe(i interface{}) {
    if s, ok := i.(string); ok {
        fmt.Println("是字符串:", s)
    } else if n, ok := i.(int); ok {
        fmt.Println("是整数:", n)
    } else {
        fmt.Println("未知类型")
    }
}
```

### Type Switch 中的 comma ok

```go
switch v := i.(type) {
case string:
    fmt.Println("字符串:", v)
case int:
    fmt.Println("整数:", v)
default:
    fmt.Println("其他类型")
}
```

---

## 3. Channel 接收

### 语法

```go
value, ok := <-ch
```

### 含义

从 channel 接收数据时：

| 结果 | 含义 |
|------|------|
| `ok` 为 `true` | 接收到了值，channel 未关闭 |
| `ok` 为 `false` | channel 已关闭且没有更多值可接收，`value` 是零值 |

### 示例

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
close(ch)

for {
    val, ok := <-ch
    if !ok {
        fmt.Println("channel 已关闭")
        break
    }
    fmt.Println("收到:", val)
}
// 输出:
// 收到: 1
// 收到: 2
// channel 已关闭
```

### 更常用的写法：range

```go
// range 会在 channel 关闭时自动退出循环
for val := range ch {
    fmt.Println("收到:", val)
}
```

---

## 总结

| 场景 | 语法 | ok = true | ok = false |
|------|------|-----------|------------|
| Map 访问 | `v, ok := m[key]` | key 存在 | key 不存在 |
| 类型断言 | `v, ok := i.(T)` | 类型匹配 | 类型不匹配 |
| Channel 接收 | `v, ok := <-ch` | 有值 | channel 已关闭 |

共同点：**用第二个返回值 `ok` 来判断操作是否成功，避免 panic 或误判。**
