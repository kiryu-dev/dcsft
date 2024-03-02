package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var tasks = [...]string{
	"\t1. Вычислить математическое ожидание времени Θ безотказной работы (средней наработки до отказа). Параметры: N = 65536; λ = 1e-5; m = 1; n = 65527, 65528, …, 65536; µ ∈ {1, 10, 100, 1000}",
	"\t2. Вычислить математическое ожидание времени Θ безотказной работы (средней наработки до отказа). Параметры: N = 65536; µ = 1; m = 1; n = 65527, 65528, …, 65536; λ ∈ {1e-5, 1e-6, 1e-7, 1e-8, 1e-9}",
	"\t3. Вычислить математическое ожидание времени Θ безотказной работы (средней наработки до отказа). Параметры: N = 65536; µ = 1; λ = 1e-5; n = 65527, 65528, …, 65536; m ∈ {1, 2, 3, 4}",
	"\t4. Вычислить среднее время T восстановления ВС со структурной избыточностью. Параметры: N = 1000; λ = 1e-3; m = 1; n = 900, 910, …, 1000; µ ∈ {1, 2, 4, 6}",
	"\t5. Вычислить среднее время T восстановления ВС со структурной избыточностью. Параметры: N = 8192; µ = 1; m = 1; n = 8092, 8102, …, 8192; λ ∈ {1e-5, 1e-6, 1e-7, 1e-8, 1e-9}",
	"\t6. Вычислить среднее время T восстановления ВС со структурной избыточностью. Параметры: N = 8192; µ = 1; λ = 1e-5; n = 8092, 8102, …, 8192; m ∈ {1, 2, 3, 4}",
}

type callable func(w io.Writer)

var taskNumToFunc = [...]callable{foo, baz, bar, qux, corge, grault}

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

func calcTheta(N, n, m, mu int, lambda float64) float64 {
	sum := 0.0
	mul := 1.0
	for i := n + 1; i <= N; i++ {
		v := 0
		if i >= N-m+1 && i <= N+1 {
			v = (N - i + 1) * mu
		} else {
			v = m * mu
		}
		mul *= float64(v) / (float64(i-1) * lambda)
		sum += mul / (float64(i) * lambda)
	}
	return sum + 1/(float64(n)*lambda)
}

func calcTao(N, n, m, mu int, lambda float64) float64 {
	if n == 1 {
		return float64(m * mu)
	}
	mul := 1.0
	for i := 1; i < n; i++ {
		mul *= float64(i) * lambda / float64(mu*i)
	}
	sum := 0.0
	for i := 1; i < n; i++ {
		v := 1.0
		for j := i; j < n; j++ {
			if j >= N-m && j <= N {
				v *= float64(j) * lambda / float64((N-j)*mu)
			} else {
				v *= float64(j) * lambda / float64(m*mu)
			}
		}
		sum += v / (float64(i) * lambda)
	}
	return mul + sum
}

func foo(w io.Writer) {
	var (
		N      = 65536
		lambda = 1e-5
		m      = 1
	)
	if _, err := w.Write([]byte("x,y1,y2,y3,y4\n")); err != nil {
		log.Fatalf("write header: %v", err)
	}
	builder := strings.Builder{}
	for n := 65527; n <= N; n++ {
		builder.WriteString(strconv.Itoa(n))
		for mu := 1; mu <= 1000; mu *= 10 {
			theta := calcTheta(N, n, m, mu, lambda)
			builder.WriteByte(',')
			builder.WriteString(strconv.FormatFloat(theta, 'f', 6, 64))
		}
		builder.WriteByte('\n')
		if _, err := w.Write([]byte(builder.String())); err != nil {
			log.Fatalf("write row: %v", err)
		}
		builder.Reset()
	}
}

func baz(w io.Writer) {
	var (
		N  = 65536
		mu = 1
		m  = 1
	)
	if _, err := w.Write([]byte("x,y1,y2,y3,y4,y5\n")); err != nil {
		log.Fatalf("write header: %v", err)
	}
	builder := strings.Builder{}
	for n := 65527; n <= N; n++ {
		builder.WriteString(strconv.Itoa(n))
		for lambda := 1e-5; lambda >= 1e-9; lambda /= 10 {
			theta := calcTheta(N, n, m, mu, lambda)
			builder.WriteByte(',')
			builder.WriteString(strconv.FormatFloat(theta, 'f', 6, 64))
		}
		builder.WriteByte('\n')
		if _, err := w.Write([]byte(builder.String())); err != nil {
			log.Fatalf("write row: %v", err)
		}
		builder.Reset()
	}
}

func bar(w io.Writer) {
	var (
		N      = 65536
		mu     = 1
		lambda = 1e-5
	)
	if _, err := w.Write([]byte("x,y1,y2,y3,y4\n")); err != nil {
		log.Fatalf("write header: %v", err)
	}
	builder := strings.Builder{}
	for n := 65527; n <= N; n++ {
		builder.WriteString(strconv.Itoa(n))
		for m := 1; m <= 4; m++ {
			theta := calcTheta(N, n, m, mu, lambda)
			builder.WriteByte(',')
			builder.WriteString(strconv.FormatFloat(theta, 'f', 6, 64))
		}
		builder.WriteByte('\n')
		if _, err := w.Write([]byte(builder.String())); err != nil {
			log.Fatalf("write row: %v", err)
		}
		builder.Reset()
	}
}

func qux(w io.Writer) {
	var (
		N      = 1000
		lambda = 1e-3
		m      = 1
	)
	if _, err := w.Write([]byte("x,y1,y2,y3,y4\n")); err != nil {
		log.Fatalf("write header: %v", err)
	}
	builder := strings.Builder{}
	for n := 900; n <= N; n += 10 {
		builder.WriteString(strconv.Itoa(n))
		for _, mu := range []int{1, 2, 4, 6} {
			tao := calcTao(N, n, m, mu, lambda)
			builder.WriteByte(',')
			builder.WriteString(strconv.FormatFloat(tao, 'f', 6, 64))
		}
		builder.WriteByte('\n')
		if _, err := w.Write([]byte(builder.String())); err != nil {
			log.Fatalf("write row: %v", err)
		}
		builder.Reset()
	}
}

func corge(w io.Writer) {
	var (
		N  = 8192
		mu = 1
		m  = 1
	)
	if _, err := w.Write([]byte("x,y1,y2,y3,y4,y5\n")); err != nil {
		log.Fatalf("write header: %v", err)
	}
	builder := strings.Builder{}
	for n := 8092; n <= N; n += 10 {
		builder.WriteString(strconv.Itoa(n))
		for lambda := 1e-5; lambda >= 1e-9; lambda /= 10 {
			tao := calcTao(N, n, m, mu, lambda)
			builder.WriteByte(',')
			builder.WriteString(strconv.FormatFloat(tao, 'f', 6, 64))
		}
		builder.WriteByte('\n')
		if _, err := w.Write([]byte(builder.String())); err != nil {
			log.Fatalf("write row: %v", err)
		}
		builder.Reset()
	}
}

func grault(w io.Writer) {
	var (
		N      = 8192
		mu     = 1
		lambda = 1e-5
	)
	if _, err := w.Write([]byte("x,y1,y2,y3,y4\n")); err != nil {
		log.Fatalf("write header: %v", err)
	}
	builder := strings.Builder{}
	for n := 8092; n <= N; n += 10 {
		builder.WriteString(strconv.Itoa(n))
		for m := 1; m <= 4; m++ {
			tao := calcTao(N, n, m, mu, lambda)
			builder.WriteByte(',')
			builder.WriteString(strconv.FormatFloat(tao, 'f', 6, 64))
		}
		builder.WriteByte('\n')
		if _, err := w.Write([]byte(builder.String())); err != nil {
			log.Fatalf("write row: %v", err)
		}
		builder.Reset()
	}
}
