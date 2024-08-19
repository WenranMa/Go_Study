# 211. 添加与搜索单词 - 数据结构设计

### 中等

请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。

实现词典类 WordDictionary ：

- WordDictionary() 初始化词典对象
- void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
- bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回  false 。word 中可能包含一些 '.' ，每个 . 都可以表示任何一个字母。

### 示例：

    输入：
    ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
    [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
    输出：
    [null,null,null,null,false,true,true,true]

    解释：
    WordDictionary wordDictionary = new WordDictionary();
    wordDictionary.addWord("bad");
    wordDictionary.addWord("dad");
    wordDictionary.addWord("mad");
    wordDictionary.search("pad"); // 返回 False
    wordDictionary.search("bad"); // 返回 True
    wordDictionary.search(".ad"); // 返回 True
    wordDictionary.search("b.."); // 返回 True

### 提示：
- 1 <= word.length <= 25
- addWord 中的 word 由小写英文字母组成
- search 中的 word 由 '.' 或小写英文字母组成
- 最多调用 10^4 次 addWord 和 search

### 解：
前缀树

```go
type TrieNode struct {
	Children map[byte]*TrieNode
    IsLeaf bool
}

func NewTrieNode() *TrieNode {
    return &TrieNode{
        Children: make(map[byte]*TrieNode),
        IsLeaf: false,
    }
}

type WordDictionary struct {
    Root *TrieNode
}

func Constructor() WordDictionary {
    return WordDictionary{
        Root: NewTrieNode(),
    }
}

func (this *WordDictionary) AddWord(word string)  {
    n:= this.Root
    for i := 0; i < len(word); i++ {
        c := word[i]
        if _, ok:= n.Children[c]; !ok {
            n.Children[c] = NewTrieNode()
        }
        n = n.Children[c]
    }
    n.IsLeaf = true
}

func (this *WordDictionary) Search(word string) bool {
	return dfsSearch(this.Root.Children, word, 0)
}

func dfsSearch(children map[byte]*TrieNode, word string, start int) bool {
    if start == len(word) {
        return len(children) == 0
    }
    c := word[start]
    if trieN, ok := children[c]; ok {
        if start == len(word) -1  && trieN.IsLeaf {
            return true
        }
        return dfsSearch(trieN.Children, word, start + 1)
    } else if c == '.' {
        result := false
        for _, n := range children {
            if start == len(word) - 1 && n.IsLeaf {
                return true
            }
            if dfsSearch(n.Children, word, start + 1) {
                result = true
            }
        }
        return result
    }
    return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
```