package gan

import (
	"reflect"
	"testing"
)

func newTestNode() *node {
	r := &node{
		children: make(map[string]*node),
	}
	r.insert("GET", "/")
	r.insert("GET", "/hello/:name")
	r.insert("GET", "/hello/b/c")
	r.insert("GET", "/assets/*filepath")
	return r
}

func TestFind(t *testing.T) {
	r := newTestNode()
	node, params := r.find("GET", "/hello/lei")
	if node == nil {
		t.Fatal("node shouldn't be nil")
	}
	if node.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}
	if params["name"] != "lei" {
		t.Fatal("name should be equal to 'lei'")
	}
	t.Logf("matched path: %s, params['name']: %s\n", node.pattern, params["name"])

}

func TestFind1(t *testing.T) {
	r := newTestNode()
	node, params := r.find("GET", "/hello/b/c")
	if node == nil {
		t.Fatal("node shouldn't be nil")
	}
	if node.pattern != "/hello/b/c" {
		t.Log(node.pattern)
		t.Fatal("should match /hello/b/c")
	}
	t.Logf("matched path: %s, params['name']: %s\n", node.pattern, params["name"])
}

func TestFind2(t *testing.T) {
	r := newTestNode()
	node, params := r.find("GET", "/assets/file1.txt")
	ok1 := node.pattern == "/assets/*filepath" && params["filepath"] == "file1.txt"
	if !ok1 {
		t.Fatal("pattern shoule be /assets/*filepath & filepath shoule be file1.txt")
	}
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}
