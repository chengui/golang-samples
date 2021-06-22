package gee

import (
    "fmt"
    "strings"
)

type Trie struct {
    isWild bool
    part string
    pattern string
    children []*Trie
}

func NewTrie() *Trie {
    return &Trie{
        part: "",
        isWild: false,
        pattern: "",
    }
}

func (trie *Trie) String() string {
    if trie == nil {
        return "nil"
    }
    strChildren := ""
    for i, child := range(trie.children) {
        if i == 0 {
            strChildren += child.String()
        } else {
            strChildren += " " + child.String()
        }
    }
    return fmt.Sprintf(
        "Trie{part:%v isWild:%v pattern:%v children:{%v}}",
        trie.part, trie.isWild, trie.pattern, strChildren)
}

func (trie *Trie) parsePattern(pattern string) []string {
    vs := strings.Split(pattern, "/")

    parts := make([]string, 0)
    for _, item := range vs {
        if item != "" {
            parts = append(parts, item)
            if item[0] == '*' {
                break
            }
        }
    }
    return parts
}

func (trie *Trie) Insert(pattern string) {
    parts := trie.parsePattern(pattern)
    cur := trie
    if len(parts) == 0 {
        cur.pattern = "/"
    }
    for i, part := range(parts) {
        var child *Trie
        for j := range(cur.children) {
            if cur.children[j].part == part {
                child = cur.children[j]
                break
            }
        }
        if child == nil {
            child = &Trie{
                part: part,
                isWild: part[0] == ':' || part[0] == '*',
            }
            cur.children = append(cur.children, child)
        }
        cur = child
        if i == len(parts) - 1 {
            cur.pattern = pattern
        }
    }
}

func (trie *Trie) Search(pattern string) (string, map[string]string) {
    parts := trie.parsePattern(pattern)
    params := make(map[string]string)
    cur := trie
    if len(parts) == 0 {
        return cur.pattern, params
    }
    for i, part := range(parts) {
        found := false
        for _, child := range(cur.children) {
            if child.part == part || child.isWild {
                if child.isWild {
                    if child.part[0] == ':' {
                        params[child.part[1:]] = part
                    }
                    if child.part[0] == '*' && len(child.part) > 1 {
                        params[child.part[1:]] = strings.Join(parts[i:], "/")
                    }
                }
                cur = child
                found = true
                break
            }
        }
        if !found {
            return "", params
        }
        if found && cur.part[0] == '*' {
            return cur.pattern, params
        }
    }
    return cur.pattern, params
}
