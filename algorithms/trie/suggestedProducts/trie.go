package suggestedProducts

//Leetcode主方法
func suggestedProducts(products []string, searchWord string) [][]string {
	trieTree := buildTrieTree(products)

	return findSuggested(trieTree, searchWord)
}

type TrieNode struct {
	data     uint8
	children [26]*TrieNode
}

func (t *TrieNode) isChildrenEmpty() bool {
	for i := 0; i < 26; i++ {
		if t.children[i] != nil {
			return false
		}
	}
	return true
}

//构建字典树
func buildTrieTree(products []string) *TrieNode {
	root := &TrieNode{
		data:     '/',
		children: [26]*TrieNode{},
	}

	p := root

	for _, product := range products {
		for i := 0; i < len(product); i++ {
			index := product[i] - 'a'

			if p.children[index] == nil {
				newNode := &TrieNode{
					data:     index,
					children: [26]*TrieNode{},
				}

				p.children[index] = newNode
			}
			p = p.children[index]
		}
		p = root
	}
	return root
}

func findSuggested(trieTree *TrieNode, searchWord string) [][]string {

}

//计算以searchWord开头的所有字符串
//如果searchWord不在字典树内，返回空
//在查询某一个节点下面的所有子串的时候，需要采取N叉树的前序遍历
func startWith(trieTree *TrieNode, searchWord string) []string {
	if trieTree == nil {
		return nil
	}

	p := trieTree

	result := []string{}

	for i := 0; i < len(searchWord); i++ {
		index := searchWord[i] - 'a'

		//顺藤摸瓜继续往下走
		if p.children[index] != nil {
			p = p.children[index]
		} else {
			return []string{}
		}
	}

}
