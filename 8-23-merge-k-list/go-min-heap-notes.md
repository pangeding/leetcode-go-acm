# Go 最小堆实现：Push/Pop 方法详解

## 代码背景

在 LeetCode "合并K个升序链表"（Merge k Sorted Lists）题目中，使用 Go 标准库 `container/heap` 实现最小堆。关键代码片段：

```go
type NodeHeap []*ListNode

func (h *NodeHeap) Push(x interface{}) {
    *h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
```

## 核心问题：`*` 指针用法

### 为什么接收者要是指针？

```go
func (h *NodeHeap) Push(x interface{}) {
    *h = append(*h, x.(*ListNode))
}
```

- `(h *NodeHeap)`：接收者是指针，因为堆需要**修改底层切片**。如果不加 `*`，函数只操作切片的拷贝，修改不会生效。
- `*h`：解引用操作，获取切片本体后执行 `append`，将新节点追加进去。

### Pop 中的切片操作

```go
old := *h          // 获取当前堆的切片副本（仍指向同一底层数组）
n := len(old)
x := old[n-1]      // 取出最后一个元素
*h = old[0 : n-1]  // 将堆切片缩小，"删除"最后一个元素
return x
```

## 核心问题：`interface{}` 用法

### 为什么是 `interface{}`？

`container/heap` 是 Go 1.18 泛型推出之前的老牌标准库，其接口定义固定要求 `Push(x interface{})` 和 `Pop() interface{}`。

- `x interface{}`：类似 Java 的 `Object` 或 "任意类型"，可以接收任何值。
- `x.(*ListNode)`：**类型断言**，将通用的 `interface{}` 转回具体的 `*ListNode` 类型。

### 使用示例

```go
// 存入时隐式转换为 interface{}
heap.Push(h, head)

// 取出时是 interface{}，必须手动断言回 *ListNode
node := heap.Pop(h).(*ListNode)
```

## 深入对比

### 1. 和 Java 泛型 `<T>` 的区别

| 特性 | Java `<T>` | Go `interface{}` |
| :--- | :--- | :--- |
| **类型安全** | 编译期检查，强类型 | 运行期断言，断言错误会导致 panic |
| **实现机制** | 类型擦除（但 API 保持类型） | 动态装箱，完全丢失类型信息 |
| **使用体验** | 声明一次，自动推导 | 存入自动、取出手动（需要 `.(T)`） |
| **性能** | 基本无额外开销 | 存在少量内存分配和反射开销 |

> **注意**：Go 1.18+ 已经支持真正的泛型语法（如 `func Do[T any](val T)`），但 `container/heap` 尚未使用泛型重构。

### 2. `interface{}` 和 Go 有方法的 `interface` 的区别？

* `interface{}`：**空接口**，没有定义任何方法。任何类型都自动满足它，所以能装任何东西。
* 普通 `interface`：定义了方法约束（如 `Stringer` 要求实现 `String() string`），只有实现了这些方法的类型才能传入。
* **本质相同**：它们都是 Go 的 interface 机制，区别仅在于对类型的约束程度（完全无约束 vs 有方法约束）。

### 3. 乱传参数的后果是 panic 崩溃
编译时没问题，但运行时会直接 panic（崩溃）。

在你的代码里：
```
func (h *NodeHeap) Push(x interface{}) {
    *h = append(*h, x.(*ListNode))  // <--- 这里挂掉
}
```

当你传入 heap.Push(h, "hello")：

程序执行到 x.(*ListNode) 时，发现 x 实际上是字符串而不是 *ListNode，会抛出 panic：
```
panic: interface conversion: interface {} is string, not *ListNode
```
为什么会这样？

因为 x.(*ListNode) 是一种强硬断言。Go 认为你十分确定这就是 *ListNode，如果不是，它认为发生了严重错误，直接终止程序。

如何安全处理？

如果想防止崩溃，可以用 "comma ok" 写法捕获错误：
```
func (h *NodeHeap) Push(x interface{}) {
    node, ok := x.(*ListNode)
    if !ok {
        // 类型不对，自己处理（比如报错、忽略），而不会 panic
        panic("传进来的不是 ListNode 节点!")
    }
    *h = append(*h, node)
}
```