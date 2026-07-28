# LeetCode 23. Merge k Sorted Lists & Heap 源码深度解析

---

## 🟢 核心重点：Go 标准库 `container/heap` 源码学习

在 Go 中实现堆非常方便，但有几个**反直觉**的设计在刷题时极易踩坑。我们通过解读 `/usr/local/go/src/container/heap/heap.go` 核心源码来彻底理解它。

### 1. 接口定义：`heap.Interface`

源码中定义了一个最小堆所需的接口：

```go
// heap.go
type Interface interface {
    sort.Interface       // 嵌入了排序接口：Len(), Less(), Swap()
    Push(x any)          // 添加 x 作为最后一个元素
    Pop() any            // 移除并返回最后一个元素
}
```

**💡 关键避坑点 (Naming Confusion)：**
请注意，接口里的 `Push(x)` 和 `Pop()` **并不负责维护堆的结构**（如上浮下沉）！

*   **小写方法 (`h.Push`)**: 是 `heap` 包内部调用的，你实现它只需写切片 append 逻辑。
*   **大写方法 (`heap.Push(h, x)`)**: 是你调用的，它会先调用小写 `h.Push` 添加数据，然后**自动执行**维护堆结构的代码。

---

### 2. 源码核心逻辑拆解

#### A. `Init` —— $O(n)$ 建堆
```go
func Init(h Interface) {
    n := h.Len()
    // 从最后一个非叶子节点开始，倒序遍历到根节点
    for i := n/2 - 1; i >= 0; i-- {
        down(h, i, n)
    }
}
```
*   **为什么从 `n/2 - 1` 开始？**
    对于大小为 $n$ 的数组，索引大于 `n/2 - 1` 的节点全是叶子节点。叶子节点本身就是满足堆性质的子树（只有一个元素），不需要处理。
*   **复杂度**: 虽然看起来像 $O(n \log n)$，但根据堆的性质，越往下的节点高度越小，数学证明总复杂度其实是严格的 **$O(n)$**。

#### B. `down` —— 下沉操作 (核心)
这是维护堆性质的主要逻辑。
```go
func down(h Interface, i0, n int) bool {
    i := i0
    for {
        j1 := 2*i + 1       // 左孩子索引
        if j1 >= n || j1 < 0 { break } // 越界

        j := j1              
        // 选择左右孩子中更小的一个（注意这里调用了 Less）
        if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
            j = j2 // 右孩子更小，选右孩子
        }
        
        // 如果孩子不比父亲小，说明堆序性未破坏，停止
        if !h.Less(j, i) { break }
        
        h.Swap(i, j) // 交换
        i = j        // 继续向下
    }
    return i > i0
}
```
*   **设计细节**: 这里不仅比较了父子，还先比较了两个子节点，选出**更小**的那个进行交换。这保证了局部最优。
*   **最小堆本质**: 因为使用了 `h.Less(j, i)` (孩子 < 父亲)，所以小的元素会冒上去。

#### C. `Push` (对外暴露的函数) —— 插入 $O(\log n)$
```go
func Push(h Interface, x any) {
    h.Push(x)            // 1. 你定义的 append 逻辑
    up(h, h.Len()-1)     // 2. 对新加的末尾元素执行上浮
}
```

#### D. `Pop` (对外暴露的函数) —— 删除堆顶 $O(\log n)$
```go
func Pop(h Interface) any {
    n := h.Len() - 1
    h.Swap(0, n)         // 1. 把堆顶(0) 和 末尾(n) 交换
    down(h, 0, n)        // 2. 对现在的堆顶执行下沉
    return h.Pop()       // 3. 移除并返回末尾（即原堆顶）
}
```
*   **为什么要交换？** Go 的切片底层是数组，**删除头部元素代价是 $O(n)$** (需要搬运所有数据)。
*   **优化**: 将堆顶和堆尾交换，然后把堆尾（原堆顶）砍掉，切片尾部删除是 $O(1)$ 的。代价是堆顶变成了一个小元素，需要 `down` 修复。

---

## 🔵 LeetCode 23. 合并 K 个升序链表 (题目实战)

基于源码理解，我们编写解题代码。
由于我们要找**最小值**，所以我们需要实现一个**最小堆**。

### 1. 实现 `MinHeap`

因为存的是链表节点指针，我们的底层切片类型是 `[]*ListNode`。

```go
package main

import "container/heap"

type ListNode struct {
    Val  int
    Next *ListNode
}

// 定义最小堆
type MinHeap []*ListNode

// 1. Len: 返回长度
func (h MinHeap) Len() int { return len(h) }

// 2. Less: 定义最小堆 (Val 小的在上面)
func (h MinHeap) Less(i, j int) bool {
    return h[i].Val < h[j].Val
}

// 3. Swap: 交换
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// 4. Push: 注意这是给 package heap 调用的，只负责 append
func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(*ListNode))
}

// 5. Pop: 注意这是给 package heap 调用的，只负责移除末尾并返回
func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
```

### 2. 解题逻辑

利用堆来维护当前所有链表的 **"头节点"**，每次都能 $O(\log k)$ 的速度找到最小的那个节点。

```go
func mergeKLists(lists []*ListNode) *ListNode {
    h := &MinHeap{}
    heap.Init(h)

    // 1. 将所有链表的头节点入堆
    for _, node := range lists {
        if node != nil {
            heap.Push(h, node)
        }
    }

    dummy := &ListNode{}
    curr := dummy

    // 2. 只要堆里还有节点，就一直取出最小的
    for h.Len() > 0 {
        // 拿出最小节点 (堆顶)
        node := heap.Pop(h).(*ListNode)
        curr.Next = node
        curr = curr.Next

        // 3. 如果该节点后面还有后续节点，将其放入堆中继续竞争
        if curr.Next != nil {
            heap.Push(h, curr.Next)
        }
    }

    return dummy.Next
}
```

### 3. 复杂度总结

- **时间复杂度**: $O(N \log k)$。
  - $N$ 是节点总数。
  - 堆中始终最多只有 $k$ 个元素。
  - 每个节点都要进堆一次出堆一次，操作代价 $\log k$。
- **空间复杂度**: $O(k)$。
  - 堆中存储 $k$ 个指针。
  - (注：如果算递归栈或结果链表则可能更高，但堆本身的额外开销仅 $k$)。
