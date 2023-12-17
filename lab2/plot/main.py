import pandas as pd
import matplotlib
import matplotlib.pyplot as plt

matplotlib.rc("font", size=18)
plt.rcParams["figure.figsize"] = (20, 10)

data = pd.read_csv("data.csv")

_, ax = plt.subplots()


# ax.plot(data["task_count"][:10].to_numpy(), data["nfdh"][:10].to_numpy(), c="purple", marker="x")
# ax.plot(data["task_count"][:10].to_numpy(), data["ffdh"][:10].to_numpy(), c="blue", marker="*")
# ax.plot(data["task_count"][10:].to_numpy(), data["nfdh"][10:].to_numpy(), c="green", marker=".")
# ax.plot(data["task_count"][10:].to_numpy(), data["ffdh"][10:].to_numpy(), c="red", marker="P")

ax.plot(data["task_count"].to_numpy(), data["nfdh"].to_numpy(), c="purple", marker="x")
ax.plot(data["task_count"].to_numpy(), data["ffdh"].to_numpy(), c="blue", marker="*")

ax.grid()
ax.set_xticks(data["task_count"][:10].to_numpy())
ax.get_xaxis().set_major_formatter(matplotlib.ticker.ScalarFormatter())
# plt.title("Время выполнения алгоритмов в зависимости от количества m задач в наборе")
# plt.legend(["NFDH (n = 1024)", "FFDH (n = 1024)", "NFDH (n = 4096)", "FFDH (n = 4096)"])
# plt.ylabel("Время выполнения алгоритма (с)\n")
# plt.xlabel("Количество задач")

# plt.title("Значений целевой функции от расписаний")
# plt.legend(["NFDH (n = 1024)", "FFDH (n = 1024)", "NFDH (n = 4096)", "FFDH (n = 4096)"])
# plt.ylabel("Значение цел. функции\n")
# plt.xlabel("Количество задач")

plt.title("Значений целевой функции от расписаний ATLAS")
plt.legend(["NFDH (n = 1152)", "FFDH (n = 1152)"])
plt.ylabel("Значение цел. функции\n")
plt.xlabel("Количество задач")

plt.savefig('plot.png')