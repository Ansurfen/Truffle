package test

import (
	"fmt"
	"testing"
	"truffle/ws/hub"
)

func TestTrie(t *testing.T) {
	trie := hub.NewRoot()
	a := trie.AddChild("/")
	a_a := trie.AddChild(a)
	a_a_a := trie.AddChild(a_a)
	a_a_b := trie.AddChild(a_a)
	fmt.Println(a_a_a, a_a_b)
	trie.Migrate(a_a_a, "/")
	fmt.Println(trie.FindChildren("/a/"))
	trie.Migrate("/a/a/", "/a/")
	fmt.Println(trie.FindChildren("/a/"))
	trie.Fork("/a/")
	trie.AddChild("/c/")
	fmt.Println(trie.FindChildren("/"))
	trie.MigrateClan("/c/", "/")
	fmt.Println("/a/a/b/ parent:", trie.FindParent("/a/a/b/"))
	fmt.Println(trie.FindChildren("/"))
	users := []string{"a1", "a2", "a3"}
	// ? 更新群聊的用户
	trie.UpdateTopic("/a/", users)
	trie.TopicStatus("/a/")
	// ? 一个用户订阅他加入的所有群聊
	topics := []string{"/a/a/b/", "/a/", "/b/a/a/"}
	trie.JoinTopic("jobs", topics)
	trie.TopicStatus("/a/")
	trie.Fork("/a/")
	trie.TopicStatus("/e/")
	trie.DelUser("/a/", "a")
	trie.DelUser("/a/", "a1")
	// trie.DelUser("/e/", "a2")
	trie.TopicStatus("/e/")
	trie.TopicStatus("/a/")
	trie.Unicast("a2", "/a/", "这是一则单播")
	trie.BroadCast("a2", "/a/", "测试广播")
	trie.Multicast("a2", "测试组播", "/a/", "/b/", "/e/")
	// ! 广播的消息需要多路复用发回去
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(trie.AddChild("/"))
	// }
}
