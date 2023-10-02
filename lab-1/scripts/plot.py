import pandas as pd
import matplotlib
import matplotlib.pyplot as plt

matplotlib.rc("font", size=18)
plt.rcParams["figure.figsize"] = (20, 10)

data = pd.read_csv("results.csv")


plt.plot(data["message_length(mb)"][:7].to_numpy(), data["time"][:7].to_numpy(), c="purple", marker="x")
plt.plot(data["message_length(mb)"][7:14].to_numpy(), data["time"][7:14].to_numpy(), c="red", marker=".")
plt.plot(data["message_length(mb)"][14:].to_numpy(), data["time"][14:].to_numpy(), c="blue", marker="P")

plt.grid()
plt.xticks(data["message_length(mb)"][:7].to_numpy())
plt.title("Зависимость времени передачи сообщения от его размера\nна разных уровнях коммуникационной среды")
plt.legend(["Уровень оперативной памяти NUMA узла", "Уровень внутрисистемной шины Intel QPI,\nобъединяющей процессоры NUMA-узлов", "Уровень сети связи между ЭМ (адаптер InfiniBand QDR)"])
plt.ylabel("Время передачи сообщения (с)\n")
plt.xlabel("Размер сообщения (Мб)")

plt.savefig('plot.png')