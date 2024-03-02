import csv
import matplotlib
import matplotlib.pyplot as plt

matplotlib.rc("font", size=18)
plt.rcParams["figure.figsize"] = (20, 10)

x = []
y = [[], [], [], [], []]

with open('./result_6.csv', newline='') as csvfile:
    for row in csv.DictReader(csvfile):
        x.append(float(row["x"]))
        y[0].append(float(row["y1"]))
        y[1].append(float(row["y2"]))
        y[2].append(float(row["y3"]))
        y[3].append(float(row["y4"]))
        # y[4].append(float(row["y5"]))

plt.plot(x, y[0], c="purple", marker="x")
plt.plot(x, y[1], c="red", marker=".")
plt.plot(x, y[2], c="blue", marker="P")
plt.plot(x, y[3], c="green", marker="D")
# plt.plot(x, y[4], c="black", marker="*")

plt.grid()
# plt.yscale("log")
plt.xticks(x)

# plt.title("Mean time between failures (N = 65536, λ = 1e-5 1/hours, m = 1)")
# plt.legend(["µ = 1 1/hours", "µ = 10 1/hours", "µ = 100 1/hours", "µ = 1000 1/hours"])

# plt.title("Mean time between failures (N = 65536, µ = 1 1/hours, m = 1)")
# plt.legend(["λ = 1e-5 1/hours", "λ = 1e-6 1/hours", "λ = 1e-7 1/hours", "λ = 1e-8 1/hours", "λ = 1e-9 1/hours"])

# plt.title("Mean time between failures (N = 65536, µ = 1 1/hours, λ = 1e-5 1/hours)")
# plt.legend(["m = 1", "m = 2", "m = 3", "m = 4"])

# plt.title("Mean time to recovery (N = 1000, λ = 1e-3 1/hours, m = 1)")
# plt.legend(["µ = 1 1/hours", "µ = 2 1/hours", "µ = 4 1/hours", "µ = 6 1/hours"])

# plt.title("Mean time to recovery (N = 8192, µ = 1 1/hours, m = 1)")
# plt.legend(["λ = 1e-5 1/hours", "λ = 1e-6 1/hours", "λ = 1e-7 1/hours", "λ = 1e-8 1/hours", "λ = 1e-9 1/hours"])

plt.title("Mean time between failures (N = 8192, µ = 1 1/hours, λ = 1e-5 1/hours)")
plt.legend(["m = 1", "m = 2", "m = 3", "m = 4"])

# plt.ylabel("Mean time between failures (hours)")
plt.ylabel("Mean time to recovery (hours)")
plt.xlabel("Number of n elementary machines in base subsystem")

plt.savefig('plot.png')