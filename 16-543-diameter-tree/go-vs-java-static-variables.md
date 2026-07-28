# Go 中为什么不需要 Java 式的 static 变量？

## 1. Java 为什么需要 static ？

在 LeetCode 刷题或日常开发中，Java 用 static 主要有以下几个场景：

### 场景1：保存全局共享状态
```java
public class Solution {
    // 多个方法、多次调用之间要共享这个值
    static int count = 0;
    static Map<Integer, Integer> memo = new HashMap<>();
    
    public int fib(int n) {
        count++;
        if (memo.containsKey(n)) return memo.get(n);
        // ...
    }
}
```

### 场景2：递归函数需要访问外部状态
```java
class Solution {
    static int maxDepth = 0;
    
    public int diameterOfBinaryTree(TreeNode root) {
        maxDepth = 0;  // 每次调用要重置
        depth(root);
        return maxDepth;
    }
    
    private int depth(TreeNode node) {
        if (node == null) return 0;
        int left = depth(node.left);
        int right = depth(node.right);
        maxDepth = Math.max(maxDepth, left + right);  // 递归中修改外部状态
        return Math.max(left, right) + 1;
    }
}
```

### 场景3：工具类的方法不需要实例化就能调用
```java
class MathUtils {
    public static int add(int a, int b) {
        return a + b;
    }
}
// 调用：MathUtils.add(1, 2)
```

## 2. Go 是怎么解决这些需求的？

### 方式一：使用结构体（代替 static 成员变量）

当你需要在多次调用之间保存状态时，Go 推荐使用**结构体**：

```go
// Go 风格：用结构体保存状态
type Solution struct {
    count int
    memo  map[int]int
}

func NewSolution() *Solution {
    return &Solution{
        memo: make(map[int]int),
    }
}

func (s *Solution) Fib(n int) int {
    s.count++
    if val, ok := s.memo[n]; ok {
        return val
    }
    // ... 计算逻辑
    s.memo[n] = result
    return result
}

// 使用
sol := NewSolution()
sol.Fib(10)
sol.Fib(20)
fmt.Println(sol.count)  // 统计了所有调用次数
```

**对比分析：**
- Java 用 static：所有实例共享，线程不安全，多次调用需要手动重置
- Go 用结构体：每个实例独立，线程安全，多次调用互不干扰

**LeetCode 中使用结构体：**
```go
func (s *Solution) diameterOfBinaryTree(root *TreeNode) int {
    // ... 方法实现
}
```

### 方式二：使用闭包（代替递归中的 static 变量）

**闭包是什么？**
闭包就是一个函数"捕获"了它外部作用域的变量。即使外部函数已经返回，内部函数依然可以访问和修改这些变量。

```go
func diameterOfBinaryTree(root *TreeNode) int {
    maxDiam := 0  // 这个变量会被内部的 depth 函数"捕获"
    
    // 声明一个函数变量
    var depth func(node *TreeNode) int
    depth = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        left := depth(node.Left)
        right := depth(node.Right)
        
        // 这里可以直接访问和修改 maxDiam！
        // 因为 depth 函数"捕获"了 maxDiam 这个变量
        if left+right > maxDiam {
            maxDiam = left + right
        }
        return max(left, right) + 1
    }
    
    depth(root)
    return maxDiam
}
```

**对比分析：**
- Java 递归：需要把 maxDiam 声明为 static，每次调用要手动重置
- Go 闭包：每次调用 diameterOfBinaryTree 都会创建一个新的 maxDiam，自动隔离，不需要手动重置

### 方式三：使用指针传参

这是 Go 另一种常见做法：

```go
func diameterOfBinaryTree(root *TreeNode) int {
    maxDiam := 0
    depth(root, &maxDiam)  // 把 maxDiam 的地址传给 depth
    return maxDiam
}

func depth(node *TreeNode, maxDiam *int) int {
    if node == nil {
        return 0
    }
    left := depth(node.Left, maxDiam)
    right := depth(node.Right, maxDiam)
    
    // 解引用指针，修改 maxDiam 指向的值
    if left+right > *maxDiam {
        *maxDiam = left + right
    }
    return max(left, right) + 1
}
```

**指针传参的原理：**
- `&maxDiam` 取的是 maxDiam 的内存地址
- `*int` 表示这个参数是一个指向 int 类型的指针
- `*maxDiam` 是解引用，访问指针指向的那个 int 值
- 修改 `*maxDiam` 就等于修改外面的 `maxDiam`

### 方式四：包级变量（代替 public static 常量/配置）

```go
package mypackage

// 包级变量，类似于 Java 的 static
const MaxRetries = 3

var DefaultConfig = &Config{
    Timeout: 30,
    Retries: MaxRetries,
}

// 包级函数，不需要 struct 实例化就能调用
func HelperFunc(a, b int) int {
    return a + b
}
```

其他文件引用：
```go
import "mypackage"

func main() {
    mypacket.HelperFunc(1, 2)  // 直接调用，不需要创建实例
    fmt.Println(mypackage.MaxRetries)
}
```

## 3. 为什么 Go 不推荐 static 风格？

### 原因1：并发安全

```java
// Java static 变量在并发时的问题
class Solution {
    static int count = 0;
    public void increment() { count++; }  // 多线程不安全！
}
```

```go
// Go 结构体实例独立，天然隔离
type Counter struct {
    count int
}
func (c *Counter) Increment() { c.count++ }

c1 := &Counter{}
c2 := &Counter{}
// c1 和 c2 各自计数，互不影响
```

### 原因2：测试友好

```java
// Java 测试时需要手动清理 static 状态
@Test
void testFib() {
    Solution.count = 0;  // 必须手动重置！
    assertEquals(5, solution.fib(5));
}
```

```go
// Go 每次创建新实例，干净的状态
func TestFib(t *testing.T) {
    sol := NewSolution()  // 全新的、干净的状态
    got := sol.Fib(5)
    if got != 5 { t.Errorf(...) }
}
```

### 原因3：更符合 Go 的设计哲学

Go 的设计理念：
- 倾向于**组合**，而不是继承和共享
- 倾向于**显式传递状态**，而不是隐式全局共享
- 倾向于**小而清晰的函数**，而不是庞大的类

## 4. LeetCode 中 Go 的常见写法

### 写法1：直接写函数（最简单）
```go
func diameterOfBinaryTree(root *TreeNode) int {
    maxDiam := 0
    var depth func(*TreeNode) int
    depth = func(node *TreeNode) int {
        // ... 使用闭包访问 maxDiam
    }
    depth(root)
    return maxDiam
}
```

### 写法2：使用 structure
```go
type Solution struct {
    maxDiam int
}

func (s *Solution) depth(node *TreeNode) int {
    // ... 使用 s.maxDiam
}

func diameterOfBinaryTree(root *TreeNode) int {
    s := &Solution{}
    s.depth(root)
    return s.maxDiam
}
```

## 总结

| Java 的 static 用途       | Go 的替代方案             |
| ------------------------- | ------------------------- |
| static 计数器/状态        | 结构体字段                |
| static 缓存/memo          | 结构体字段 或 闭包捕获    |
| static 递归辅助变量       | 闭包 或 指针传参          |
| static 工具方法           | 包级函数（直接在 main 外写） |
| static final 常量         | const 或 var 包级变量     |
| public static 配置        | var 包级变量              |

**核心思想：** Go 用"显式传递"和"结构体"代替了 Java 的"隐式共享 static 状态"。
