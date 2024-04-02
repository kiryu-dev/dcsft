package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Введи путь до csv файла для записи результатов: ")
	filepath := ""
	_, _ = fmt.Scan(&filepath)
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatalf("open file '%s': %v", filepath, err)
	}
	var (
		N         = 65536
		lambdaArr = []float64{1e-5, 1e-6, 1e-7}
	)
	_, _ = file.WriteString("lambda,mu,m,n,theta,tao\n")
	for _, lambda := range lambdaArr {
		for mu := 1; mu <= 1000; mu *= 10 {
			for m := 1; m < 4; m++ {
				for n := 65527; n <= N; n++ {
					theta := calcTheta(N, n, m, mu, lambda)
					tao := calcTao(N, n, m, mu, lambda)
					_, _ = fmt.Fprintf(file, "%f,%d,%d,%d,%f,%f\n", lambda, mu, m, n, theta, tao)
				}
			}
		}
	}
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
