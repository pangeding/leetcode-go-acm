# `container/list` 实现 LRU Cache

## `container/list` 核心 API

```go
list.New()           // 创建一个空的双向链表
list.Len()           // 获取链表长度

// 插入操作
list.PushFront(v)    // 在链表头部插入元素 v，返回 *list.Element
list.PushBack(v)     // 在链表尾部插入元素

// 获取端点
list.Front()         // 返回头部元素 (*list.Element)
list.Back()          // 返回尾部元素 (*list.Element)

// 移动/删除
list.MoveToFront(e)  // 将元素 e 移动到头部
list.Remove(e)       // 从链表中删除元素 e
```

## 关于 `*list.Element`

`list.Element` 的结构是：
```go
type Element struct {
    Value interface{}  // 存储的实际值，类型是空接口（任意类型）
    // 还有 next/prev 指针，但内部使用，不需要管
}
```

所以代码里的 `elem.Value.(*entry)` 是**类型断言**——把 `interface{}` 转回 `*entry` 类型。

## 代码逻辑说明

1. **`map[int]*list.Element`**：key → 链表节点的映射，O(1) 查找
2. **`list.Front()`** = 最近使用的元素（LRU 的"头部"）
3. **`list.Back()`** = 最久未使用的元素（需要淘汰的）
4. **`Get`**：命中时 `MoveToFront` 标记为最近使用
5. **`Put`**：已存在则更新值 + `MoveToFront`；不存在则 `PushFront`，满了就淘汰 `Back()`

## 需要注意的点

- `interface{}` 是空接口类型，可以存储任意类型的值
- 取出时需要用类型断言 `elem.Value.(*entry)` 转回原始类型
- 这个文件没有 `package` 声明，LeetCode 核心模式一般是 `package main` 或者不需要写
