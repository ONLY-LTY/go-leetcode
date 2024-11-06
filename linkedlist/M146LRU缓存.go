package linkedlist

type Element struct {
	Key, Val  int
	Pre, Next *Element
}

type LRUCache struct {
	head, tail *Element
	cache      map[int]*Element
	capacity   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{cache: make(map[int]*Element), capacity: capacity}
}

// Get 访问元素之后 将元素移动到最前面
func (c *LRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		c.moveToFront(node)
		return node.Val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.cache[key]; ok {
		//更新元素
		node.Val = value
		c.moveToFront(node)
	} else {
		//新增元素
		newNode := &Element{Val: value, Key: key}
		c.cache[key] = newNode
		c.add(newNode)
	}
	//超出容量 删除末尾元素
	if len(c.cache) > c.capacity {
		delete(c.cache, c.tail.Key)
		c.remove(c.tail)
	}
}

// moveToFront 移动到头节点 先删除 在添加
func (c *LRUCache) moveToFront(node *Element) {
	c.remove(node)
	c.add(node)
}

// add 新增元素 插入到头结点
func (c *LRUCache) add(node *Element) {
	node.Pre = nil
	node.Next = c.head
	//如果头结点不为null 头结点的前一个节点指向node
	if c.head != nil {
		c.head.Pre = node
	}
	//头结点指向node
	c.head = node
	//如果尾节点为null 尾结点指向node
	if c.tail == nil {
		c.tail = node
	}
}

func (c *LRUCache) remove(node *Element) {
	//如果删除的是头结点
	if node == c.head {
		//头结点指向删除节点的下一个节点
		c.head = node.Next
		//删除节点
		node.Next = nil
		return
	}
	//如果删除的是尾节点
	if node == c.tail {
		c.tail = node.Pre
		node.Pre.Next = nil
		node.Pre = nil
		return
	}
	//中间节点
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}
