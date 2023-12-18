package strategymatrix

type strategyMatrix struct {
	c [][]int
	a []int
	b []int
	p []int
	q []int
	i int
}

func New(c []int, n uint64) *strategyMatrix {
	matrix := genPaymentMatrix(c, int(n))
	return &strategyMatrix{
		c: matrix,
		a: make([]int, n),
		b: make([]int, n),
		p: make([]int, n),
		q: make([]int, n),
	}
}

func (s *strategyMatrix) Run(eps float64) float64 {
	s.iterate()
	dif := s.Max()-s.Min()
	for dif > eps {
		s.iterate()
		dif = s.Max()-s.Min()
	}
	return dif
}

func (s *strategyMatrix) iterate() {
	s.i++
	i := argmax(s.a)
	for j, val := range s.c[i] {
		s.b[j] += val
	}
	s.q[i]++
	i = argmin(s.b)
	for row := range s.c {
		s.a[row] += s.c[row][i]
	}
	s.p[i]++
}

func (s *strategyMatrix) Min() float64 {
	min := s.b[0]
	for _, v := range s.b {
		if min > v {
			min = v
		}
	}
	return float64(min) / float64(s.i)
}

func (s *strategyMatrix) Max() float64 {
	max := s.a[0]
	for _, v := range s.a {
		if max < v {
			max = v
		}
	}
	return float64(max) / float64(s.i)
}

func genPaymentMatrix(c []int, n int) [][]int {
	var (
		res = make([][]int, n)
		max = 0
	)
	for i := range res {
		res[i] = make([]int, n)
		j := 0
		for ; j <= i; j++ {
			res[i][j] = j*c[0] + (i-j)*c[1]
			if res[i][j] > max {
				max = res[i][j]
			}
		}
		for ; j < n; j++ {
			res[i][j] = i*c[1] + (j-i)*c[2]
			if res[i][j] > max {
				max = res[i][j]
			}
		}
	}
	// fmt.Println(res)
	return res
}

func argmin(arr []int) int {
	idx := 0
	for i := range arr {
		if arr[i] < arr[idx] {
			idx = i
		}
	}
	return idx
}

func argmax(arr []int) int {
	idx := 0
	for i := range arr {
		if arr[i] > arr[idx] {
			idx = i
		}
	}
	return idx
}

func (s *strategyMatrix) P() []float64 {
	p := make([]float64, 0, len(s.p))
	for _, v := range s.p {
		p = append(p, float64(v)/float64(s.i))
	}
	return p
}

func (s *strategyMatrix) Q() []float64 {
	q := make([]float64, 0, len(s.q))
	for _, v := range s.q {
		q = append(q, float64(v)/float64(s.i))
	}
	return q
}
