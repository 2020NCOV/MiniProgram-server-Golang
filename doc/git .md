## git —— 如何与上游保持同步

**1. 添加上游仓库**（添加过的可以略过，直接第二步）

```
git remote add upstream [upstream_url]
```

**2. Fetch 上游代码**

```
git fetch upstream
```

**3. 切换到本地master分支** （如果只有一个主分支就不用切换）

```
git checkout master
```

**4. 将upstream/master merge到本地master分支** 

```
git merge upstream/master
```

**5. push到自己的github仓库**

```
git push origin master
```
