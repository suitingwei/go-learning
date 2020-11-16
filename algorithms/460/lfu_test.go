package _60

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

//带有head和tail的双向链表
type Node struct {
	Previous  *Node
	Next      *Node
	Key       int
	Value     int
	Frequency int //节点的频率，这个频率在get/put的时候需要进行更新
}

type DoubleLinkedList struct {
	Head  *Node
	Tail  *Node
	Count int
}

//
func NewDoubleLinkedList() *DoubleLinkedList {
	head := &Node{}
	tail := &Node{}

	head.Next = tail
	tail.Previous = head

	return &DoubleLinkedList{
		Head: head, Tail: tail, Count: 0,
	}
}

func (list *DoubleLinkedList) RemoveLast() *Node {
	lastNode := list.Tail.Previous

	//异常case
	if lastNode == list.Head {
		return nil
	}

	lastNode.Previous.Next = lastNode.Next
	lastNode.Next.Previous = lastNode.Previous

	list.Count--

	return lastNode
}

func (list *DoubleLinkedList) print() {
	p := list.Head

	for {
		if p == nil {
			break
		}

		if p == list.Head {
			fmt.Printf("[ %d ", p.Value)
		} else if p == list.Tail {
			fmt.Printf(" %d ]", p.Value)
		} else {
			fmt.Printf(" <-[k_%d,v_%d,f_%d]-> ", p.Key, p.Value, p.Frequency)
		}

		p = p.Next
	}
}

//从双向连表里删除这个节点
func (list *DoubleLinkedList) Remove(node *Node) {
	node.Previous.Next = node.Next
	node.Next.Previous = node.Previous

	list.Count--

	//断开这个节点的前后指针，避免有bug不好查
	node.Next = nil
	node.Previous = nil
}

//把节点加到连表的头部
func (list *DoubleLinkedList) AddToHead(node *Node) {
	node.Next = list.Head.Next
	node.Previous = list.Head

	//这两行调到顺序，就等着debug到死吧
	//记住点，不能先去改要用到的节点
	list.Head.Next.Previous = node
	list.Head.Next = node

	list.Count++

}

func (list *DoubleLinkedList) Empty() bool {
	return list.Count == 0
}

type FrequencyMap map[int]*DoubleLinkedList

type KeyMap map[int]*Node

type LFUCache struct {
	Capacity     int
	FrequencyMap FrequencyMap
	KeyMap       KeyMap
	MinFrequency int //记录当前最小的频率
}

func Constructor(capacity int) LFUCache {

	return LFUCache{
		Capacity:     capacity,
		FrequencyMap: make(FrequencyMap),
		KeyMap:       make(KeyMap, capacity),
		MinFrequency: 0,
	}
}

//Get操作
//如果找不到这个key，直接返回-1
//如果找到了这个Key，那么需要更新他的频率到对应的FrequencyMap
//这里需要在O(1)的时间内知道这个节点所在的是哪一个FrequencyMap
//所以需要在节点中记录现在的频率
func (this *LFUCache) Get(key int) int {
	//这个key不存在
	if _, ok := this.KeyMap[key]; !ok {
		this.print(fmt.Sprintf("__GET(%d)__\tResponse=-1", key))
		return -1
	}

	//找到这个节点
	node := this.KeyMap[key]

	//提高他的频率
	this.increaseNodeFrequency(node)

	this.print(fmt.Sprintf("__GET(%d)__\t Response=%d", key, node.Value))

	return node.Value
}

func (this *LFUCache) increaseNodeFrequency(node *Node) {
	//从低频的list里移动到高频的list
	oldList := this.getFrequencyList(node.Frequency)

	//从低频的list里删除这个节点
	oldList.Remove(node)

	//如果低频的list空了，那么全局的最低频率要更新
	if oldList.Empty() && node.Frequency == this.MinFrequency {
		this.MinFrequency = node.Frequency + 1
	}

	//获取下一个频率的对应连表
	newList := this.getFrequencyList(node.Frequency + 1)

	//把这个节点添加到这个频率的连表最前面，这样频率相同的时候，淘汰最后使用的
	newList.AddToHead(node)

	//这个节点对应的频率要更新
	node.Frequency++
}

//找到这个频率对应的节点列表
func (this *LFUCache) getFrequencyList(frequency int) *DoubleLinkedList {
	//如果还没有达到这个频率的，那么就创建一个新的频率list
	if _, ok := this.FrequencyMap[frequency]; !ok {
		this.FrequencyMap[frequency] = NewDoubleLinkedList()
	}

	return this.FrequencyMap[frequency]
}

//对于一个已经存在的节点，更新他的频率，移动frequencyList
func (this *LFUCache) Put(key int, value int) {
	if this.Capacity == 0 {
		return
	}
	//对于已经存在的节点，put跟get基本一样，除了还需要更新value
	if node, ok := this.KeyMap[key]; ok {
		this.increaseNodeFrequency(node)

		//更新value
		node.Value = value

		this.print(fmt.Sprintf("__PUT(%d,%d)__", key, value))

		return
	}

	//对于不存在的节点,首先要创建这个节点，然后加到频率为1的节点里
	//淘汰最小频率的的list里的最后一个节点
	//更新LFU的总量
	newNode := &Node{
		Key: key, Value: value, Frequency: 1,
	}

	//如果LFU已经满了，那么淘汰频率最小的队列里的最后一个元素
	if len(this.KeyMap) >= this.Capacity {
		minList := this.getFrequencyList(this.MinFrequency)

		lastNode := minList.RemoveLast()

		//删除索引map
		if lastNode != nil {
			delete(this.KeyMap, lastNode.Key)
		}
	}

	//加到频率为1的连表头
	this.getFrequencyList(1).AddToHead(newNode)

	//注意每次新增一个节点的时候，最小频率都是1
	this.MinFrequency = 1

	this.KeyMap[key] = newNode

	this.print(fmt.Sprintf("__PUT(%d,%d)__", key, value))
}

func (this *LFUCache) print(prefix string) {
	fmt.Println(prefix)

	fmt.Printf("MinFrequency=%d\n", this.MinFrequency)
	for fre, list := range this.FrequencyMap {
		fmt.Printf("Frequency[%d]= ", fre)
		list.print()
		fmt.Println()
	}
	fmt.Println()
}

func TestLFU(t *testing.T) {
	obj := Constructor(2)
	ref := reflect.ValueOf(&obj)

	//operates:=`[[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]`
	operates := `[[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]`

	var operatesArr [][]int
	_ = json.Unmarshal([]byte(operates), &operatesArr)

	for i, operate := range operatesArr {
		fmt.Printf("Step=%d\t", i)
		if len(operate) == 1 {
			params := make([]reflect.Value, 1)
			params[0] = reflect.ValueOf(operate[0])
			ref.MethodByName("Get").Call(params)
		} else {
			params := make([]reflect.Value, 2)
			params[0] = reflect.ValueOf(operate[0])
			params[1] = reflect.ValueOf(operate[1])
			ref.MethodByName("Put").Call(params)
		}
	}
}
