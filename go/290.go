package leetcode

import "strings"

func wordPattern(pattern string, s string) bool {
    p2smap := map[byte]string{}
    s2pmap := map[string]byte{}
    words := strings.Split(s, " ")
    if len(pattern) != len(words) {
        return false
    }
    for i, word := range words {
        p := pattern[i]
        if s2pmap[word] > 0 && p2smap[p] != word || p2smap[p] != "" && s2pmap[word] != p {
            return false
        }
        p2smap[p] = word
        s2pmap[word] = p
    }
    return true
}
