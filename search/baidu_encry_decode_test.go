package search

import (
	"fmt"
	"testing"
)

func TestDecodeBaiduEncURL(t *testing.T) {
	fmt.Println(DecodeBaiduEncURL("https://www.baidu.com/link?url=w__LfN75x1OTuz4HqnR9gwDbdWg0Er5qqKQF5rEZmteoyLlkwT74kF_EJvsMRvtu&wd=&eqid=affb93eb0001153f000000065c74ca55"))
}
