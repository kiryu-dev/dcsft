import csv
import matplotlib
import matplotlib.pyplot as plt

matplotlib.rc("font", size=18)
plt.rcParams["figure.figsize"] = (20, 10)

x = []
y = [[], [], [], [], [], [], []]

with open('./result_2.csv', newline='') as csvfile:
    for row in csv.DictReader(csvfile):
        x.append(float(row["x"]))
        y[0].append(float(row["y1"]))
        y[1].append(float(row["y2"]))
        y[2].append(float(row["y3"]))
        y[3].append(float(row["y4"]))
        y[4].append(float(row["y5"]))
        y[5].append(float(row["y6"]))
        y[6].append(float(row["y7"]))

plt.plot(x, y[0], c="purple", marker="x")
plt.plot(x, y[1], c="red", marker=".")
plt.plot(x, y[2], c="blue", marker="P")
plt.plot(x, y[3], c="green", marker="d")
plt.plot(x, y[4], c="black", marker="p")
plt.plot(x, y[5], c="cyan", marker="*")
plt.plot(x, y[6], c="orange", marker="8")

plt.grid()
plt.xticks(x)

# plt.title("График зависимости функции оперативной надежности от времени\n(N = 10, λ = 0,024 1/ч, μ = 0,71 1/ч, m = 1)")
# plt.legend(["n = 8", "n = 9", "n = 10"])

plt.title("График зависимости функции оперативной восстановимости от времени\n(N = 16, λ = 0,024 1/ч, μ = 0,71 1/ч, m = 1)")
plt.legend(["n = 10", "n = 11", "n = 12", "n = 13", "n = 14", "n = 15", "n = 16"])

# plt.ylabel("Значение функции оперативной надежности")
plt.ylabel("Значение функции оперативной восстановимости")
plt.xlabel("Время, ч")

plt.savefig('plot.png')