package cf_daily

import (
	"fmt"
	"io"
)

func CF480C(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var n, a, b, k int
	fmt.Fscan(in, &n, &a, &b, &k)
	if a > b {
		a = n + 1 - a
		b = n + 1 - b
	}
	f := make([]int, b)
	f[a] = 1
	s := make([]int, b+1)
	for ; k > 0; k-- {
		for i, v := range f {
			s[i+1] = (s[i] + v) % mod
		}
		for y := 1; y < b; y++ {
			f[y] = (s[y+(b-y-1)/2+1] - f[y]) % mod
		}
	}
	ans := int64(0)
	for _, v := range f {
		ans += int64(v)
	}
	fmt.Fprint(out, (ans%mod+mod)%mod)
}
