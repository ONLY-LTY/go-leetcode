package linkedlist

type N struct {
	Key, Val  int
	Pre, Next *N
}

type LRUCache struct {
	head, tail *N
	cache      map[int]*N
	Cap        int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{cache: make(map[int]*N), Cap: capacity}
}

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
		newNode := &N{Val: value, Key: key}
		c.cache[key] = newNode
		c.add(newNode)
	}
	//超出容量 删除末尾元素
	if len(c.cache) > c.Cap {
		delete(c.cache, c.tail.Key)
		c.remove(c.tail)
	}
}

func (c *LRUCache) moveToFront(node *N) {
	c.remove(node)
	c.add(node)
}

func (c *LRUCache) add(node *N) {
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

func (c *LRUCache) remove(node *N) {
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
