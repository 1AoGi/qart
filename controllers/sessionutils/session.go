package sessionutils

import "fmt"

func SessionKey(sha string, keys ...string) string {
	ret := sha
	for _, key := range keys {
		ret += fmt.Sprintf(":%v", key)
	}
	return ret
}
