package hub

import (
	"fmt"
	"log"
	"path"
	"sync"
)

const ROOT string = "/"

type TopicMeta struct {
	user map[string]bool
	Name string // vitrual name
}

type Trie struct {
	isWord   bool
	cnt      *Counter
	children [26]*Trie
	Topic    *TopicMeta
}

func NewRoot() *Trie {
	t := &Trie{cnt: &Counter{}}
	t.Insert(ROOT)
	return t
}

func NewTrie(topic bool) *Trie {
	// ? s 即 / 不需要创建 topic
	if topic {
		return &Trie{cnt: &Counter{}, Topic: &TopicMeta{user: make(map[string]bool)}}
	}
	return &Trie{cnt: &Counter{}}
}

func (t *Trie) UpdateTopic(topic string, user []string) {
	if node := t.FindNode(topic); node != nil && node.Topic != nil {
		for _, u := range user {
			node.Topic.user[u] = true
		}
	} else {
		println("fail to join topic")
	}
}

func (t *Trie) DelUser(topic, user string) {
	if node := t.FindNode(topic); node != nil && node.Topic != nil {
		if _, ok := node.Topic.user[user]; ok {
			delete(node.Topic.user, user)
		} else {
			println("user don't exist")
		}
	}
}

func (t *Trie) Multicast(user, msg string, topics ...string) {
	for _, topic := range topics {
		t.Unicast(user, topic, msg)
	}
}

func (t *Trie) Unicast(user, topic, msg string) ([]string, bool) {
	var ret []string
	if node := t.FindNode(topic); node != nil && node.Topic != nil {
		if _, ok := node.Topic.user[user]; user == "truffle" || ok {
			for u := range node.Topic.user {
				if user != u {
					fmt.Println(u, "get a msg:", msg, "by", user, "at", topic)
				}
				ret = append(ret, u)
			}
			return ret, true
		} else {
			println("user don't exist in topic")
		}
	} else {
		println("fail to find topic")
	}
	return ret, false
}

func (t *Trie) BroadCast(user, topic, msg string) {
	topics := t.FindChildren(topic)
	for _, topic := range topics {
		t.Unicast("truffle", topic, msg)
	}
}

func (t *Trie) JoinTopic(user string, topics []string) {
	for _, topic := range topics {
		if node := t.FindNode(topic); node != nil {
			node.Topic.user[user] = true
		}
	}
}

func (t *Trie) TopicStatus(topic string) {
	if node := t.FindNode(topic); node != nil && node.Topic != nil {
		fmt.Print("Online user: ")
		for user := range node.Topic.user {
			fmt.Printf("%s ", user)
		}
	}
	println()
}

func (t *Trie) AddChild(parent string) string {
	node := t.FindNode(parent)
	child := node.cnt.Gen()
	node.Insert(child)
	return parent + child
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

// ? clone pattern

func (t *Trie) Clone(word string) *Trie {
	return t.FindNode(word)
}

func (t *Trie) Fork(word string) {
	if word == ROOT {
		return
	}
	node := t.Clone(word)
	p := t.FindParent(word)
	key := t.FindNode(p).cnt.Gen()
	t.FindNode(p).Insert(key + RemoveBind(word))
	t.FindNode(p).FindNode(key + RemoveBind(word)).cnt = node.cnt
	for user := range node.Topic.user {
		t.FindNode(p).FindNode(key + RemoveBind(word)).Topic.user[user] = true
	}
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

func Encoding(path []byte) []byte {
	for i, ch := range path {
		switch ch {
		case '/':
			path[i] = 'a'
		case 'z':
			path[i] = 's'
		default:
			path[i] = ch + 1
		}
	}
	return path
}

func Decoding(path []byte) []byte {
	for i, ch := range path {
		switch ch {
		case 'a':
			path[i] = '/'
		case 's':
			path[i] = 'z'
		default:
			path[i] = ch - 1
		}
	}
	return path
}

type Counter struct {
	lock int
	sync.Mutex
}

func (c *Counter) SetIdx(lk int) {
	c.lock = lk
}

func (c *Counter) Gen() string {
	c.Lock()
	str := ""
	if c.lock == 0 {
		str += "a"
	}
	for i := c.lock; i > 0; i /= 26 {
		ch := i % 26
		if ch == 's'-'a' {
			c.lock++
			c.Unlock()
			return c.Gen()
		} else {
			str += string(rune(ch + 'a'))
		}
	}
	c.Unlock()
	c.lock++
	return str + "/"
}
