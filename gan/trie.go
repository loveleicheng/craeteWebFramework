package gan

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string           // 完整uri
	part     string           // 被/分割的字符串
	children map[string]*node // 子节点
	isWild   bool             // 是否为:或*等形式的模糊匹配
}

// n作为根节点的角色，从根节点开始插入
// /a/:b 与 /a/:b/c 可以并存
// /a/:b 与 /a/b/c 可以并存
// /a/:b 与 /a/:d/c 冲突
// 即:表示的动态变量在子节点只能有一个
// /a/*b 与 /a/c/d 冲突
func (n *node) insert(method string, pattern string) {
	parts := parsePattern(pattern)
	if _, ok := n.children[method]; !ok {
		n.children[method] = &node{
			children: make(map[string]*node),
		}
	}
	root := n.children[method]
	for _, part := range parts {
		if root.children[part] == nil {
			// 新节点模糊没有模糊匹配
			if part[0] != ':' && part[0] != '*' {
				// 有带*号的模糊匹配子节点，则不需要继续创建了
				if root.checkWildChildAsterisk() == true {
					panic(fmt.Sprintf("router existed：%s", pattern))
				} else {
					root.children[part] = &node{
						part:     part,
						children: make(map[string]*node),
						isWild:   part[0] == ':' || part[0] == '*',
					}
				}
			} else {
				// 新节点有模糊匹配
				// 没有模糊匹配的子节点，则新建
				if root.checkWildChild() == false {
					root.children[part] = &node{
						part:     part,
						children: make(map[string]*node),
						isWild:   part[0] == ':' || part[0] == '*',
					}
				} else {
					// 有带*号的模糊匹配子节点
					// 或者有带:号的模糊匹配子节点，且变量名字不同
					panic(fmt.Sprintf("router existed：%s", pattern))
				}

			}
		}
		root = root.children[part]
	}
	root.pattern = pattern
}

// 查看子节点是否有模糊匹配
func (n *node) checkWildChild() bool {
	for _, val := range n.children {
		if val.part[0] == ':' || val.part[0] == '*' {
			return true
		}
	}
	return false
}

// 查看子节点是否有冒号:模糊匹配
func (n *node) checkWildChildColon() bool {
	for _, val := range n.children {
		if val.part[0] == ':' {
			return true
		}
	}
	return false
}

// 查看子节点是否有星号*模糊匹配
func (n *node) checkWildChildAsterisk() bool {
	for _, val := range n.children {
		if val.part[0] == '*' {
			return true
		}
	}
	return false
}

// n作为根节点，从根开始查询
// pattern is ture visited uri
// n作为根节点的角色，从根节点开始插入
// /a/:b 与 /a/:b/c 可以并存
// /a/:b 与 /a/b/c 可以并存
// /a/:b/c 与 /a/b/c 可以并存, 但是/a/b/c的精准匹配优先，优先级暂时不实现
// /a/:b 与 /a/:d/c 冲突
// 即:表示的动态变量在子节点只能有一个
// /a/*b 与 /a/c/d 冲突
func (n *node) find(method, pattern string) (child *node, params map[string]string) {
	params = make(map[string]string)
	if _, ok := n.children[method]; !ok {
		return nil, params
	}
	parts := parsePattern(pattern)
	root := n.children[method]
	child = &node{
		children: make(map[string]*node),
	}
	high := len(parts)
	readyQueue := make([]*node, 1)
	waitQueue := []*node{root}
	for i, part := range parts {
		readyQueue = waitQueue
		waitQueue = waitQueue[:0]
		// 每一层的part都和对应层的子节点们进行比较
		for _, root = range readyQueue {
			for _, child = range root.children {
				if child.part == part || child.isWild {
					if child.part[0] == ':' {
						params[child.part[1:]] = part
					}
					if child.part[0] == '*' {
						params[child.part[1:]] = strings.Join(parts[i:], "/")
						return
					}
					waitQueue = append(waitQueue, child)
					// 已经是最后一段，并且符合结尾项目特征，有pattern
					if i == high-1 && child.pattern != "" {
						return
					}
				}
			}
		}
	}
	return nil, params
}

// 解析uri路径
func parsePattern(pattern string) (parts []string) {
	values := strings.Split(pattern, "/")
	for _, val := range values {
		if val != "" {
			parts = append(parts, val)
			if val[0] == '*' {
				break
			}
		}
	}
	return parts
}
