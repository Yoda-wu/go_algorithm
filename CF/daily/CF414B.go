package cf_daily

import (
	"fmt"
	"io"
)

func CF414B(in io.Reader, out io.Writer) {
	var n , k int
	// 输入n和k，表示k个数，每个数最大值不超过n
	fmt.Fscan(in, &n , &k)
	
}
