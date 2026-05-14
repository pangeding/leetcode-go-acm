# 2026-05-14 Git Merge Workflow

## 问题背景

执行 `git pull origin main` 时遇到分支 diverged（分叉）错误：

```
From github.com:pangeding/leetcode-go-acm
 * branch            main       -> FETCH_HEAD
hint: You have divergent branches and need to specify how to reconcile them.
fatal: Need to specify how to reconcile divergent branches.
```

## 原因分析

本地分支和远程分支有不同的提交，产生了分叉：

- **本地 (HEAD → main)**: `6b7c9f5` - "24 swap pairs"
- **远程 (origin/main)**: `21256e0` - "206"、`48e140c` - "206 wrong why??"、`c52c9fe` - "206 reverse"

两个分支从 `4932ba2` ("1") 之后开始分叉。

## 解决方案

### 选择 Merge 方式合并

执行以下命令将远程分支合并到本地：

```bash
git merge origin/main --no-edit
```

合并结果：
- 合并策略：`ort`
- 新增文件：
  - `1-160-link.go` (更新)
  - `2-206-reverse/2-206-reverse.go`
  - `2-206-reverse/wrong-answer.md`
  - `3-234-huiwen/3-234-huiwen.go`
- 共 4 个文件变更，+105 行，-2 行

### 合并后的 Git 历史

```
*   2955a6a Merge remote-tracking branch 'origin/main'
|\
| * 21256e0 206
| * 48e140c 206 wrong why?? 递归失败了，为什么只有1个
| * c52c9fe 206 reverse
* | 6b7c9f5 24 swap pairs
|/
* 4932ba2 1
* b2180a8 160
* 6639cf7 first commit
```

### 推送到远程

合并成功后，推送更新到远程仓库：

```bash
git push origin main
```

## 其他可选方案

### 1. Rebase（变基）

保持线性历史：

```bash
git pull --rebase origin main
```

### 2. 配置默认 Pull 策略

```bash
# 默认 merge
git config pull.rebase false

# 默认 rebase
git config pull.rebase true

# 仅允许 fast-forward
git config pull.ff only
```

## 经验总结

1. 当本地和远程都有新提交时，`git pull` 会因分支分叉而失败
2. 可以使用 `git log --oneline --graph --decorate --all` 查看分叉情况
3. Merge 会创建合并提交，保留完整历史；Rebase 会重放本地提交，保持线性历史
4. 对于个人项目，Rebase 历史更清晰；协作项目通常用 Merge 保留完整记录
