package _46

import (
	"fmt"
	"testing"
)

type LRUCache struct {
	Capacity     int
	listHead     *node
	listTail     *node
	indexMap     map[int]*node
	currentCount int
}

//双向连表
type node struct {
	Key      int
	Next     *node
	Previous *node
	Value    int
}

func Constructor(capacity int) LRUCache {
	head := &node{
		Next:     nil,
		Previous: nil,
		Value:    -1, //dummy
	}
	tail := &node{
		Value:    -1,
		Next:     nil,
		Previous: head,
	}
	head.Next = tail
	return LRUCache{
		Capacity:     capacity,
		listHead:     head,
		listTail:     tail,
		indexMap:     make(map[int]*node, capacity),
		currentCount: 0,
	}
}

//把get的节点，挪到第一个位置
//1->2->3->4->5
//如果Get(3)，那么get之后，连表变成：3->1->2->4->5
func (this *LRUCache) Get(key int) int {

	tempNode, ok := this.indexMap[key]

	if !ok || tempNode == nil {
		return -1
	}

	return this.moveToHead(tempNode).Value
}

func (this *LRUCache) Put(key int, value int) {
	existedNode, ok := this.indexMap[key]

	//如果put的这个值存在，那么更新他的值，然后挪到连表最前头即可
	if ok {
		//挪到队列头部
		this.moveToHead(existedNode)

		//更新值
		existedNode.Value = value

		return
	}

	//空间足够用
	if this.currentCount >= this.Capacity {
		lastNode := this.listTail.Previous

		//从indexMap删除
		delete(this.indexMap, lastNode.Key)

		//从双向连表删除
		lastNode.Previous.Next = lastNode.Next
		lastNode.Next.Previous = lastNode.Previous

		this.currentCount--
	}
	newNode := &node{
		Key:      key,
		Value:    value,
		Previous: this.listHead,
		Next:     this.listHead.Next,
	}

	this.listHead.Next.Previous = newNode
	this.listHead.Next = newNode

	this.indexMap[key] = newNode
	this.currentCount++
}

func (this *LRUCache) moveToHead(tempNode *node) *node {

	//带有dummy head和dummy tail,无需检测是否为null
	tempNode.Previous.Next = tempNode.Next
	tempNode.Next.Previous = tempNode.Previous

	//优化一发，无需这个临时的中间节点
	//newNode := &node{
	//	Key:      tempNode.Key,
	//	Value:    tempNode.Value,
	//	Next:     this.listHead.Next, //这里一定要在上面删除节点的操作之后，再进行
	//	Previous: this.listHead,
	//}

	//不要忘了吧这个节点的前后指针进行修改
	tempNode.Previous = this.listHead
	tempNode.Next = this.listHead.Next

	//不要忘记修改head的前后
	this.listHead.Next.Previous = tempNode
	this.listHead.Next = tempNode
	return tempNode
}

func (this *LRUCache) print() {
	p := this.listHead

	for {
		if p == nil {
			break
		}

		if p == this.listHead {
			fmt.Printf("%d ", p.Value)
		} else if p == this.listTail {
			fmt.Printf(" %d ", p.Value)
		} else {
			fmt.Printf(" <=> [Key:%d]:[Value:%d]<=> ", p.Key, p.Value)
		}

		p = p.Next
	}
	fmt.Println()

	fmt.Printf("IndexMap:%v\n", this.indexMap)
	fmt.Println()
}

func TestLRU(t *testing.T) {
	cache := Constructor(2)

	cache.Put(2, 1)
	cache.print()
	cache.Put(2, 2)
	cache.print()
	cache.Get(2) // 返回  1
	cache.print()
	cache.Put(3, 3) // 该操作会使得密钥 2 作废
	cache.Get(2)    // 返回 -1 (未找到)
	cache.Put(4, 4) // 该操作会使得密钥 1 作废
	cache.Get(1)    // 返回 -1 (未找到)
	cache.Get(3)    // 返回  3
	cache.Get(4)    // 返回  4

}
