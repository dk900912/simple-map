<h1 id="bidqS"><font style="color:#DF2A3F;">如何发布个人开发的模块</font></h1>
<h3 id="tSkTD"><font style="color:#601BDE;">github上建立代码仓库</font></h3>

![](https://cdn.nlark.com/yuque/0/2025/png/27955462/1736927498124-09a28c2f-4d6d-4e3d-bda8-615a6df8d453.png)

<h3 id="zGM5m"><font style="color:#601BDE;">克隆代码仓库</font></h3>

```shell
git@github.com:dk900912/simple-map.git
```

<h3 id="zzc0u"><font style="color:#601BDE;">编写代码</font></h3>

```go
package sorted

type SortedMap struct {
	Keys   []string
	Values []int
}

func (sm *SortedMap) Set(key string, value int) {
	sm.Keys = append(sm.Keys, key)
	sm.Values = append(sm.Values, value)
}

func (sm *SortedMap) Get(key string) (int, bool) {
	for i, k := range sm.Keys {
		if k == key {
			return sm.Values[i], true
		}
	}
	return 0, false
}
```

```go
package sorted

import "testing"

func TestSortedMap(t *testing.T) {
	sm := &SortedMap{}

	// 测试 Set 方法
	sm.Set("key1", 1)
	sm.Set("key2", 2)

	// 测试 Get 方法
	value, ok := sm.Get("key1")
	if !ok || value != 1 {
		t.Errorf("Expected value 1 for key1, got %d, ok: %v", value, ok)
	}

	value, ok = sm.Get("key2")
	if !ok || value != 2 {
		t.Errorf("Expected value 2 for key2, got %d, ok: %v", value, ok)
	}

	// 测试不存在的键
	value, ok = sm.Get("key3")
	if ok {
		t.Errorf("Expected key3 to not exist, but got value: %d", value)
	}

	// 测试插入顺序
	sm.Set("key3", 3)
	expectedKeys := []string{"key1", "key2", "key3"}
	for i, key := range expectedKeys {
		if sm.Keys[i] != key {
			t.Errorf("Expected key %s at index %d, got %s", key, i, sm.Keys[i])
		}
	}
}
```

<h3 id="husU6"><font style="color:#601BDE;">初始化Go模块</font></h3>

```shell
go mod init github.com/dk900912/simple-map
```

<h3 id="HHLVV"><font style="color:#601BDE;">提交到代码仓库</font></h3>
<h6 id="pMgyH">提交代码到远程仓库</h6>

```bash
git push origin main:main
```

<h3 id="OmV8z"><font style="color:#601BDE;">其他项目使用该Go模块</font></h3>
<font style="color:#601BDE;">首</font>先，在代码中直接引入该模块即可，如下所示：

```go
package main

import (
	"fmt"
	sm "github.com/dk900912/simple-map/sorted"
)

func main() {
	sortedmap := &sm.SortedMap{}
	sortedmap.Set("a", 1)
	sortedmap.Set("b", 2)
	fmt.Println(sortedmap.Get("a"))
}
```

<font style="color:#601BDE;">然</font>后，下载该Go模块，如下所示：

```shell
hello-module> go get github.com/dk900912/simple-map
go: downloading github.com/dk900912/simple-map v0.0.0-20250115025209-168bc3f7b649
go: added github.com/dk900912/simple-map v0.0.0-20250115025209-168bc3f7b649
```

<font style="color:#601BDE;">最</font>后，`go.mod`内容已经被自动追加了依赖，如下所示：

```shell
module hello-module

go 1.22.0

require github.com/dk900912/simple-map v0.0.0-20250115025209-168bc3f7b649 // indirect
```

<h4 id="lWNFc"><font style="color:#270070;">疑问一：</font>`<font style="color:#270070;">v0.0.0-20250115025209-168bc3f7b649</font>`<font style="color:#270070;">有什么含义</font></h4>

`v0.0.0-20250115025209-168bc3f7b649`是Go自动生成的版本信息，中间是`git commit time`，尾部是长度40位`git commit id`的前12位。

<h4 id="TXczi"><font style="color:#270070;">疑问二：为什么是</font>`<font style="color:#270070;">v0.0.0</font>`<font style="color:#270070;">开头</font></h4>
因为我们刚才提交代码的时候，一是没有在本地打标签，二是没有没有将标签提交到远程仓库。接下来，咱们来完善这两步，如下：

```git
hello-module> git push origin --tags
Total 0 (delta 0), reused 0 (delta 0), pack-reused 0 (from 0)
To github.com:dk900912/simple-map.git
 * [new tag]         v1.0.0 -> v1.0.0
```

![](https://cdn.nlark.com/yuque/0/2025/png/27955462/1736911526668-45ee25e8-126c-4b70-bd50-0bfc3ea5e0bd.png)

紧接着，再次执行`go get`，如下所示：

```shell
hello-module> go get github.com/dk900912/simple-map
go: downloading github.com/dk900912/simple-map v1.0.0
go: upgraded github.com/dk900912/simple-map v0.0.0-20250115025209-168bc3f7b649 => v1.0.0
```

```shell
module hello-module

go 1.22.0

require github.com/dk900912/simple-map v1.0.0 // indirect
```

最后，如果`simple-map`这一Go模块又追加了一次提交，但并没有发布新的Tag，而我们`hello-module`项目中有需要该最新追加提交的功能，此时我们可以通过12位`git commit id`来获取，你会发现版本号自动变为了`v1.0.1`。

```git
> go get github.com/dk900912/simple-map@c2d95cbcd20f
go: downloading github.com/dk900912/simple-map v1.0.1-0.20250115033734-c2d95cbcd20f
go: upgraded github.com/dk900912/simple-map v1.0.0 => v1.0.1-0.20250115033734-c2d95cbcd20f
```

```git
module hello-module

go 1.22.0

require github.com/dk900912/simple-map v1.0.1-0.20250115033734-c2d95cbcd20f // indirect

```

> **疑问一与疑问二中提到的**`**v0.0.0**`**和**`**v1.0.1**`**是具有普适性的，只要符合文中这两种场景，那么伪版本号中的第一部分就具有这样的规律！！！**
>

<h3 id="HcqIT"><font style="color:#601BDE;">继续发布v2.0.0</font></h3>
一定要更新`go.mod`文件，即以`v2`作为后缀。

```go
module github.com/dk900912/simple-map/v2
```

如果只是单纯`git tag v2.0.0`，后续试图引入该模块的`main module`会报错的，报错信息如下：

```go
hello-module> go get github.com/dk900912/simple-map@v2.0.0
go: github.com/dk900912/simple-map@v2.0.0: reading https://mirrors.tencent.com/go/github.com/dk900912/simple-map/@v/v2.0.0.info: 404 Not Found
        server response:
        go mod download -json github.com/dk900912/simple-map@v2.0.0:
        {
                "Path": "github.com/dk900912/simple-map",
                "Version": "v2.0.0",
                "Error": "github.com/dk900912/simple-map@v2.0.0: invalid version: module contains a go.mod file, so major version must be compatible: should be v0 or v1, not v2"
        }
```

主要就是因为，你认为你打了`v2.0.0`的标签且已发布到代码仓库了，但是由于没有在`go.mod`文件中声明`v2`，那么Go工具链还是认为这是一个`v1`版本：_invalid version: module contains a go.mod file, so major version must be compatible: should be v0 or v1, not v2。_

不仅`go get`会报错，当你直接把`require github.com/dk900912/simple-map v2.0.0`手动加入到模块文件中时，依然会报错：_version "v2.0.0" invalid: should be v0 or v1, not v2。_

<h3 id="eMk40"><font style="color:#601BDE;">继续发布v3.0.0</font></h3>
<font style="color:#DF2A3F;">删除</font><font style="color:#DF2A3F;">go.mod文件，详细指令如下：</font>

```go
simple-map> git commit -a -m "del mod file"
simple-map> git tag v3.0.0
simple-map> git push origin main --tags
```

![](https://cdn.nlark.com/yuque/0/2025/png/27955462/1737203001100-c9fc764d-8e3b-431b-b04b-7b3cd05d96ff.png)

接下来，我们在`hello-module`这一`main module`来引入`v3.0.0`的`simple-map`：

```go
package main

import (
	"fmt"
	sm "github.com/dk900912/simple-map/sorted"
)

func main() {
	sortedmap := &sm.SortedMap{}
	sortedmap.Set("a", 1)
	sortedmap.Set("b", 2)
	fmt.Println(sortedmap.Get("a"))
	fmt.Println(sortedmap.Get("a"))
}
```

```shell
hello-module>go get github.com/dk900912/simple-map@v3.0.0
go: downloading github.com/dk900912/simple-map v3.0.0+incompatible
go: added github.com/dk900912/simple-map v3.0.0+incompatible
```

最后，看一下`hello-module`下的`go.mod`，确实看到了期望看到的`incompatible`：

```shell
module hello-module

go 1.22.0

require (
	github.com/dk900912/simple-map v3.0.0+incompatible // indirect
)
```

<font style="color:#DF2A3F;">在 Go 语言中，当你看到 </font><font style="color:#DF2A3F;">+incompatible</font><font style="color:#DF2A3F;"> 的标记时，这通常意味着所下载的模块没有遵循 Go Modules 的版本控制规范。具体来说，这种情况可能发生在以下几种情况下：</font>

1. **<font style="color:#DF2A3F;">没有 </font>**<font style="color:#DF2A3F;">go.mod</font>**<font style="color:#DF2A3F;"> 文件</font>**<font style="color:#DF2A3F;">：如果模块的根目录中没有 </font><font style="color:#DF2A3F;">go.mod</font><font style="color:#DF2A3F;"> 文件，Go 将无法确定该模块的版本，因此会将其标记为不兼容。</font>
2. **<font style="color:#DF2A3F;">使用的版本不符合语义版本控制</font>**<font style="color:#DF2A3F;">：Go Modules 期望模块遵循语义版本控制（SemVer）。如果模块的版本号不符合 SemVer 的格式（例如，版本号实际已经大于等于 2 但是却不以 </font><font style="color:#DF2A3F;">v2</font><font style="color:#DF2A3F;">或者</font><font style="color:#DF2A3F;">v3</font><font style="color:#DF2A3F;">开头），Go 可能会将其视为不兼容。</font>

