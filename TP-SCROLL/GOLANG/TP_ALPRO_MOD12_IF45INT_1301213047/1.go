package main

type trec struct {
	v1, v4 int
	vx     struct {
		v2, v3 int
	}
}

func manyVal(rec *trec) {
	var n []int
	n = []int{rec.v1, rec.vx.v2, rec.vx.v3, rec.v4}
	for i, min := 0, n[0]; i < len(n); i++ {
		if min > n[i] {
			// Replace
			min = n[i]
		}
		rec.v1 = min
	}
	rec.vx.v2 = 0
	for _, v := range n {
		rec.vx.v2 += v
	}
	rec.vx.v3 = rec.vx.v2 / len(n)
	for i, max := 0, n[0]; i < len(n); i++ {
		if max < n[i] {
			max = n[i]
		}
		rec.v4 = max
	}
}

func main() {

}
