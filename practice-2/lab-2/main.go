package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var tasks = [...]string{
	"\t1. Вычислить значение функции R*(t) оперативной надежности для следующих значений параметров: N = 10; n ∈ {8, 9, 10}; λ = 0,024 1/ч; μ = 0,71 1/ч; m = 1; t = 0, 2, 4, ..., 24 ч.",
	"\t2. Вычислить значение функции U*(t) оперативной восстановимости для следующих значений параметров: N = 16; n ∈ {10, 11, ..., 16}; λ = 0,024 1/ч; μ =0,711/ч; m = 1; t = 0, 2, 4, ..., 24 ч.",
	"\t3. Вычислить значение показателя S для следующих значений параметров: N = 16; λ = 0,024 1/ч; μ = 0,71 1/ч.",
}

type callable func(w io.Writer)

var taskNumToFunc = [...]callable{foo, bar, baz}

func main() {
	fmt.Println("Выбери вариант задачи:")
	for _, task := range tasks {
		fmt.Println(task)
	}
	selected := 0
	fmt.Print(">")
	_, _ = fmt.Scan(&selected)
	for selected < 1 || selected > len(tasks) {
		fmt.Print("Введен некорректный вариант задачи. Попробуйте еще раз\n>")
		_, _ = fmt.Scan(&selected)
	}
	fmt.Print("Введи путь до csv файла для записи результатов: ")
	filepath := ""
	_, _ = fmt.Scan(&filepath)
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatalf("open file '%s': %v", filepath, err)
	}
	taskNumToFunc[selected-1](file)
}

func foo(w io.Writer) {
	var (
		N      = 10
		m      = 1
		lambda = 0.024
		mu     = 0.71
	)
	if _, err := w.Write([]byte("x,y1,y2,y3\n")); err != nil {
		log.Fatalf("write header: %v", err)
	}
	builder := strings.Builder{}
	for t := 0; t < 25; t += 2 {
		builder.WriteString(strconv.Itoa(t))
		for n := 8; n < 11; n++ {
			rStar := rStar(n, N, m, t, lambda, mu)
			builder.WriteByte(',')
			builder.WriteString(strconv.FormatFloat(rStar, 'f', 6, 64))
		}
		builder.WriteByte('\n')
	}
	if _, err := w.Write([]byte(builder.String())); err != nil {
		log.Fatalf("write row: %v", err)
	}
}

func bar(w io.Writer) {
	var (
		N      = 16
		m      = 1
		lambda = 0.024
		mu     = 0.71
	)
	if _, err := w.Write([]byte("x,y1,y2,y3,y4,y5,y6,y7\n")); err != nil {
		log.Fatalf("write header: %v", err)
	}
	builder := strings.Builder{}
	for t := 0; t < 25; t += 2 {
		builder.WriteString(strconv.Itoa(t))
		for n := 10; n < 17; n++ {
			uStar := uStar(n, N, m, t, lambda, mu)
			builder.WriteByte(',')
			builder.WriteString(strconv.FormatFloat(uStar, 'f', 6, 64))
		}
		builder.WriteByte('\n')
	}
	if _, err := w.Write([]byte(builder.String())); err != nil {
		log.Fatalf("write row: %v", err)
	}
}

func baz(w io.Writer) {
	var (
		N      = 16
		m      = []int{1, 16}
		lambda = 0.024
		mu     = 0.71
	)
	if _, err := w.Write([]byte("n,m1,m16\n")); err != nil {
		log.Fatalf("write header: %v", err)
	}
	builder := strings.Builder{}
	for n := 11; n < 17; n++ {
		builder.WriteString(strconv.Itoa(n))
		for i := range m {
			s := S(n, N, m[i], lambda, mu)
			builder.WriteByte(',')
			builder.WriteString(strconv.FormatFloat(s, 'f', 6, 64))
		}
		builder.WriteByte('\n')
	}
	if _, err := w.Write([]byte(builder.String())); err != nil {
		log.Fatalf("write row: %v", err)
	}
}

func rStar(n int, N int, m int, t int, lambda float64, mu float64) float64 {
	if t == 0 {
		return s(n, N, m, lambda, mu)
	}
	result := 0.0
	for i := n; i <= N; i++ {
		v := 0.0
		for j := 0; j <= i-n; j++ {
			v += math.Pow(float64(i)*lambda*float64(t), float64(j)) /
				fact(j) * math.Exp(float64(-i)*lambda*float64(t))
		}
		result += p(i, N, lambda, mu) * v
	}
	return result
}

func uStar(n int, N int, m int, t int, lambda float64, mu float64) float64 {
	result := 0.0
	for i := 0; i < n; i++ {
		v := 0.0
		for j := 0; j < n-i; j++ {
			v += math.Pow(mu*float64(t), float64(j)) / fact(j) *
				(dt(N-i-m)*math.Pow(float64(m), float64(j))*math.Exp(-1*float64(i)*mu*float64(t)) +
					dt(m+i-N)*math.Pow(float64(N-i), float64(j))*math.Exp(float64(i-N)*mu*float64(t)))
		}
		result += p(i, N, lambda, mu) * v
	}
	return 1 - result
}

func S(n int, N int, m int, lambda float64, mu float64) float64 {
	if m == N {
		return 1 - math.Pow(lambda, float64(N-n+1))*math.Pow(lambda+mu, float64(n-N-1))
	}
	result := 0.0
	for i := 0; i < n; i++ {
		result += p(i, N, lambda, mu)
	}
	return 1 - result
}

func s(n int, N int, m int, lambda float64, mu float64) float64 {
	if m == N {
		return 1 - math.Pow(lambda, float64(N-n+1))*math.Pow(lambda+mu, float64(n-N-1))
	}
	result := 0.0
	for i := 0; i < n; i++ {
		result += p(i, N, lambda, mu)
	}
	return 1 - result
}

func p(i int, N int, lambda float64, mu float64) float64 {
	v := 0.0
	for j := 0; j <= N; j++ {
		v += math.Pow(mu/lambda, float64(j)) / fact(j)
	}
	v *= fact(i)
	return math.Pow(mu/lambda, float64(i)) / v
}

func dt(v int) float64 {
	if v < 0 {
		return 0.0
	}
	return 1.0
}

func fact(n int) float64 {
	if n < 2 {
		return 1
	}
	v := float64(n)
	return math.Sqrt(2*math.Pi*v) * math.Pow(v/math.E, v)
}
