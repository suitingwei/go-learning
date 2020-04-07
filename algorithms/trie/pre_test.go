package trie

import (
	"fmt"
	"testing"
)

type Trie struct {
	children [26]*Trie //该字符的后续字符
	isEnd    bool      //该字符是否是终止字符，当树中有: apple,app的时候，如果搜索app，那么一定要知道最后一个p是终止字符
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		children: [26]*Trie{}, // a-z的字符数组，如果是叶子节点，存放nil
		isEnd:    false,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if word == "" {
		return
	}

	char := word[0]

	//计算出来这个字符所对应的child 索引

	index := char - 'a'

	//如果是叶子节点，那么挂上面
	if this.children[index] == nil {
		node := &Trie{
			children: [26]*Trie{},
			isEnd:    len(word) == 1,
		}

		this.children[index] = node
	} else {
		//如果该字符是原来的终止字符，那么更改原节点的isEnd
		if len(word) == 1 {
			this.children[index].isEnd = true
		}
	}

	//如果这个节点已经存在了，那么继续遍历他的子节点
	this.children[index].Insert(word[1:])
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if word == "" {
		return this.hasNoChildren()
	}

	char := word[0]

	//计算出来这个字符所对应的child 索引
	index := char - 'a'

	//结束条件,当前字符就是目标字符，并且没有后续
	childNode := this.children[index]
	if len(word) == 1 && childNode != nil && childNode.isEnd {
		return true
	}

	if this.children[index] == nil {
		return false
	} else {
		return this.children[index].Search(word[1:])
	}

}

//是否没有后续
func (this *Trie) hasNoChildren() bool {
	for i := 0; i < 26; i++ {
		if this.children[i] != nil {
			return false
		}
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(word string) bool {
	if word == "" {
		return true
	}

	char := word[0]

	//计算出来这个字符所对应的child 索引
	index := char - 'a'

	//结束条件,当前字符就是目标字符，并且没有后续
	childNode := this.children[index]
	if len(word) == 1 && childNode != nil {
		return true
	}

	if this.children[index] == nil {
		return false
	} else {
		return this.children[index].StartsWith(word[1:])
	}
}

func (this *Trie) PrintTree() {
	queue := []*Trie{}
	for i := 0; i < 26; i++ {
		if this.children[i] != nil {
			queue = append(queue, this.children[i])
		}
	}

	for {
		if len(queue) == 0 {
			break
		}
		levelNums := len(queue)
		tempQueue := []*Trie{}
		for i := 0; i < levelNums; i++ {
			fmt.Printf(" %c ", queue[i].data)

			for j := 0; j < 26; j++ {
				if queue[i].children[j] != nil {
					tempQueue = append(tempQueue, queue[i].children[j])
				}
			}
		}
		fmt.Println()
		queue = tempQueue
	}
}

func TestTrieTree(t *testing.T) {

	obj := Constructor()
	obj.Insert("apple")
	fmt.Println(obj.Search("apple"))
	fmt.Println(obj.Search("app"))
	fmt.Println(obj.StartsWith("app"))
	obj.Insert("app")
	fmt.Println(obj.Search("app"))
}
