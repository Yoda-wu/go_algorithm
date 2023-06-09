package cf_daily

import (
	"fmt"
	"io"
)

func CF414B(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, k, ans int
	ans = 1
	// 输入n和k，表示k个数，每个数最大值不超过n
	fmt.Fscan(in, &n, &k)
	// f[i][j] 表示前i个数尾数是j的数组个数。 枚举因子(除法）困难，所以枚举倍数
	// f[i][k] += f[i-1][j] k 为 j的倍数

	f := make([]int, n+1)
	f[1] = 1
	for i := 1; i <= k; i++ {
		for j := n; j > 0; j-- {
			for k := j * 2; k <= n; k += j { // 2倍 3倍，4倍...
				f[k] = (f[k] + f[j]) % mod
			}
		}
	}

	for _, v := range f {
		ans = (v + ans) % mod
	}
	fmt.Fprint(out, ans)
}
