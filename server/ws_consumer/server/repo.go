package server

import (
	"log"
	"path"
	"sync"
	"truffle/ws/ddd/po"
)

const ROOT string = "/"

type PayLoad struct {
	msgs     []po.Message
	lastTime int64
}

type Trie struct {
	isWord   bool
	cnt      *Counter
	children [26]*Trie
	payload  PayLoad
}

func NewRoot() *Trie {
	t := &Trie{cnt: &Counter{}}
	t.Insert(ROOT)
	return t
}

func NewTrie(topic bool) *Trie {
	// ? s 即 / 不需要创建 topic
	if topic {
		return &Trie{cnt: &Counter{}}
	}
	return &Trie{cnt: &Counter{}}
}

func (t *Trie) AddPayLoad(path string, msg po.Message) {
	if node := t.FindNode(path); node != nil {
		node.payload.msgs = append(node.payload.msgs, msg)
	} else {
		t.Insert(path)
		t.AddPayLoad(path, msg)
	}
}

func (t *Trie) InspectPayLoad(path string) int {
	if node := t.FindNode(path); node != nil {
		return len(node.payload.msgs)
	} else {
		return -1
	}
}

func (t *Trie) Insert(word string) {
	cur := t
	for i, c := range word {
		if c == '/' {
			c = 's'
		}
		n := c - 'a'
		if cur.children[n] == nil {
			if c == 's' {
				cur.children[n] = NewTrie(true)
			} else {
				cur.children[n] = NewTrie(false)
			}
		}
		cur = cur.children[n]
		if i == len(word)-1 {
			if cur.isWord {
				log.Println("Duplicate key")
				return
			}
			cur.isWord = true
		}
	}
}

func (t *Trie) Delete(word string) {
	cur := t
	for _, c := range word {
		if c == '/' {
			c = 's'
		}
		n := c - 'a'
		if cur.children[n] == nil {
			return
		}
		cur = cur.children[n]
	}
	if cur != nil && cur.isWord {
		cur.isWord = false
	}
}

func (t *Trie) Search(word string) bool {
	cur := t
	for _, c := range word {
		if c == '/' {
			c = 's'
		}
		n := c - 'a'
		if cur.children[n] == nil {
			return false
		}
		cur = cur.children[n]
	}
	return cur.isWord
}

func (t *Trie) StartsWith(prefix string) bool {
	cur := t
	for _, c := range prefix {
		if c == '/' {
			c = 's'
		}
		n := c - 'a'
		if cur.children[n] == nil {
			return false
		}
		cur = cur.children[n]
	}
	return true
}

func (t *Trie) DumpRoot(word string) (ret []string) {
	cur := t
	for i := 'a'; i <= 'z'; i++ {
		n := i - 'a'
		if cur.children[n] != nil {
			if i != 's' {
				word += string(i)
			} else {
				word += "/"
			}
			cur = cur.children[n]
			ret = append(ret, cur.DumpRoot(word)...)
			word = word[:len(word)-1]
			cur = t // last layer
		}
	}
	if cur.isWord {
		ret = append(ret, word)
	}
	return ret
}

func (t *Trie) FindNode(word string) *Trie {
	cur := t
	for _, c := range word {
		if c == '/' {
			c = 's'
		}
		n := c - 'a'
		if cur.children[n] == nil {
			return nil
		}
		cur = cur.children[n]
	}
	return cur
}

func (t *Trie) FindChildren(word string) (ret []string) {
	cur := t
	root := ""
	for _, c := range word {
		root += string(c)
		if c == '/' {
			c = 's'
		}
		n := c - 'a'
		if cur.children[n] == nil {
			return ret
		}
		cur = cur.children[n]
	}
	return cur.DumpRoot(root)
}

func (t *Trie) FindParent(word string) string {
	parent := path.Dir(path.Dir(word))
	if parent[len(parent)-1] != '/' {
		parent += "/"
	}
	return parent
}

func (t *Trie) FindAncestors(word string) (ret []string) {
	for {
		if word = path.Dir(word); word != "/" {
			ret = append(ret, word)
		} else {
			break
		}
	}
	return ret
}

func (t *Trie) MigrateClan(from, to string) {
	t.migrateClan(from, t.FindNode(to))
}

func (t *Trie) migrateClan(from string, target *Trie) {
	children := t.FindChildren(from)
	p := target.cnt.Gen()
	for _, c := range children {
		t.Delete(c)
		target.Insert(p + RemoveBind(c))
	}
}

func (t *Trie) Migrate(from, to string) string {
	return t.migrate(from, t.FindNode(to))
}

func (t *Trie) migrate(from string, target *Trie) string {
	p := target.cnt.Gen()
	t.Delete(from)
	node := p + RemoveBind(from)
	target.Insert(node)
	return node
}

func RemoveBind(word string) string {
	word = word[1:]
	flag := false
	str := ""
	for _, c := range word {
		if c == '/' {
			flag = true
		}
		if flag {
			str += string(c)
		}
	}
	return str[1:]
}

func BuildNode(word ...string) string {
	node := ROOT
	for _, c := range word {
		node += c
	}
	return node
}

type Counter struct {
	lock int
	sync.Mutex
}

func (c *Counter) Gen() string {
	c.Lock()
	defer c.Unlock()
	str := ""
	if c.lock == 0 {
		str += "a"
	}
	for i := c.lock; i > 0; i /= 26 {
		ch := i % 26
		if ch == 's'-'a' {
			c.lock++
			return c.Gen()
		} else {
			str += string(rune(ch + 'a'))
		}
	}
	c.lock++
	return str + "/"
}
