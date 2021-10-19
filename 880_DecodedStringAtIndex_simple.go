package main
import "fmt"

func main() {
	var s string
	var k int
	//s = "leet2code3"
	//k = 10
	//s = "ha22"
	//k = 5
	//s = "a2345678999999999999999"
	//k = 1
	s = "a2b3c4d5e6f7g8h9"
	k = 10
	ret := decodeAtIndex(s, k)
	fmt.Println(ret)
}

type repeatStrFragment struct {
	strStart int
	strEnd int
	repeat int
	totalSize int
}

func decodeAtIndex(s string, k int) string {
	k--
	i:=0
	for ;s[i] >='0'&&s[i]<='9';i++ {}
	var fragments []*repeatStrFragment
	preIsStr := false
	for ;i < len(s);i++ {
		if s[i] >= '0' && s[i] <= '9' {
			fragments[len(fragments) - 1].repeat *= int(s[i] - '0')
			if len(fragments) > 1 {
				fragments[len(fragments) - 1].totalSize = (fragments[len(fragments) - 2].totalSize + fragments[len(fragments) - 1].strEnd - fragments[len(fragments) - 1].strStart) * fragments[len(fragments) - 1].repeat
			} else {
				fragments[len(fragments) - 1].totalSize = (fragments[len(fragments) - 1].strEnd - fragments[len(fragments) - 1].strStart) * fragments[len(fragments) - 1].repeat
			}
			preIsStr = false
		} else {
			if preIsStr {
				fragments[len(fragments) - 1].strEnd++
				fragments[len(fragments) - 1].totalSize ++
			} else {
				var fragment *repeatStrFragment
				if len(fragments) > 0 {
					fragment = &repeatStrFragment{i, i + 1, 1, fragments[len(fragments) - 1].totalSize + 1}
				} else {
					fragment = &repeatStrFragment{i, i + 1, 1, 1}
				}
				fragments = append(fragments, fragment)
			}
			preIsStr = true
		}
		if len(fragments) > 0 && fragments[len(fragments) - 1].totalSize > k {
			for fi:= len(fragments) - 1;fi>=0;fi-- {
				fragment := fragments[fi]
				mod := k % (fragment.totalSize / fragment.repeat)
				if mod < fragment.totalSize / fragment.repeat - (fragment.strEnd - fragment.strStart) {
					k = mod
					continue
				} else {
					mod = mod - (fragment.totalSize / fragment.repeat - (fragment.strEnd - fragment.strStart)) + fragment.strStart
					return s[mod:mod + 1]
				}
			}
		}
	}
	return ""
}
