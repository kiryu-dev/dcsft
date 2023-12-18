import pandas as pd
import matplotlib
import matplotlib.pyplot as plt

matplotlib.rc("font", size=18)
plt.rcParams["figure.figsize"] = (20, 10)

data = pd.read_csv("results.csv")

_, ax = plt.subplots()


ax.plot(data["n"].to_numpy(), data["time"].to_numpy(), c="purple", marker="x")

ax.grid()
ax.set_xticks(data["n"].to_numpy())
ax.get_xaxis().set_major_formatter(matplotlib.ticker.ScalarFormatter())
plt.title("Зависимость вpeмeни paбoты aлгopитмa oт кoличecтвa n элeмeнтapныx мaшин в cиcтeмe")
plt.legend(["c1 = 3\nc2 = 5\nc3 = 6\neps = 0.001"])
plt.ylabel("Время работы алгоритма (с)\n")
plt.xlabel("Количество элементарных машин")

plt.savefig('plot.png')