import csv
import matplotlib
import matplotlib.pyplot as plt

matplotlib.rc("font", size=18)
plt.rcParams["figure.figsize"] = (20, 10)

x = []
y = []

with open('./theta_tao_lambda.csv', newline='') as csvfile:
    for row in csv.DictReader(csvfile):
        x.append(float(row["n"]))
        y.append(float(row["tao"]))

plt.plot(x[:10], y[:10], c="purple", marker="x")
plt.plot(x[10:20], y[10:20], c="red", marker=".")
plt.plot(x[20:30], y[20:30], c="blue", marker="P")
# plt.plot(x[30:], y[30:], c="black", marker="d")

plt.grid()
# plt.yscale("log")
plt.xticks(x)


# plt.title("Зависимость Θ от λ (N = 65536, m = 2, µ = 1 1/hours)")
plt.title("Зависимость T от λ (N = 65536, m = 2, µ = 1 1/hours)")
plt.legend(["λ = 1e-5 1/hours", "λ = 1e-6 1/hours", "λ = 1e-7 1/hours"])

# plt.ylabel("Mean time between failures (hours)")
plt.ylabel("Mean time to recovery (hours)")
plt.xlabel("Number of n elementary machines")

plt.savefig('plot.png')