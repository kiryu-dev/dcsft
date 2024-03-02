#!/bin/bash

bold=$(tput bold)
basic=$(tput sgr0)

taskjob1="#!/bin/bash\n\n#SBATCH --nodes=1\n#SBATCH --job-name=lab1\n\ncd \$SLURM_SUBMIT_DIR\n\nsrun --ntasks=2 --cpus-per-task=2 --cpu-bind=v,sockets --mpi=pmi2 mpiexec --report-bindings ./bin/main --msg-size "
taskjob2="#!/bin/bash\n\n#SBATCH --nodes=1\n#SBATCH --job-name=lab1\n\ncd \$SLURM_SUBMIT_DIR\n\nsrun --ntasks=2 --ntasks-per-node=2 --cpus-per-task=1 --cpu-bind=v,cores --mpi=pmi2 mpiexec --report-bindings ./bin/main --msg-size "
taskjob3="#!/bin/bash\n\n#SBATCH --nodes=2 --ntasks-per-node=1\n#SBATCH --job-name=lab1\n\ncd \$SLURM_SUBMIT_DIR\n\nsrun mpiexec --report-bindings ./bin/main --msg-size "


function select_task() {
    >&2 echo -e "${bold}Выбери задачу, которую хочешь решить 🤓✨🧙‍♂️🚀🔥🌟💻*ੈ✩‧₊˚\n\n${basic}"
    while true; do
        >&2 echo "${bold}1.${basic} Измерить время обмена сообщениями на разных уровнях коммуникационной сети"
        >&2 echo "${bold}2.${basic} Построить расписание решения параллельных задач на распределенной ВС"
        >&2 echo "${bold}3.${basic} Решить задачу 'Диcпeтчeр-вычиcлитeльный цeнтр'"
        >&2 echo -e "\nЖми q, чтобы выйти\n"
        read -n1 -r res
        case $res in
            [123q] ) break;;
            * ) >&2 echo -e "${bold}\nВыбери задачу по ее номеру! 🤬🤬🤬\n${basic}";;
        esac
    done
    echo $res
}

function launch_first_task() {
    clear
    echo -e "${bold}Заходим в ./lab-1/...\n"
    cd ./lab-1
    echo -e "Собираем программу...\n"
    make build
    read -p "Введи размер сообщения в Мб, которое хочешь отправлять: ${basic}" msgsize
    echo -e "${bold}Выбери уровень коммуникационной среды\n\n${basic}"
    while true; do
        echo "${bold}1.${basic} Оперативная память NUMA/SMP узлов"
        echo "${bold}2.${basic} Внутрисистемная шина Intel QPI, объединяющую процессоры NUMA-узлов"
        echo "${bold}3.${basic} Сеть связи между ЭМ"
        read -n1 -r mode
        case $mode in
            1)
                echo -e $taskjob1 $msgsize > task.job
                sbatch task.job
                break;;
            2)
                echo -e $taskjob2 $msgsize > task.job
                sbatch task.job
                break;;
            3)
                echo -e $taskjob3 $msgsize > task.job
                sbatch task.job
                break;;
            *) echo -e "${bold}\nВыбери уровень коммуникационной среды по его номеру! 🤬🤬🤬\n${basic}";;
        esac
    done
}

function launch_second_task() {
    clear
    echo -e "${bold}Заходим в ./lab-2/...\n"
    cd ./lab-2
    echo -e "Собираем программу...\n"
    make
    read -p "Введи путь до файла с задачами: ${basic}" filepath
    ./main ${filepath}
}

function launch_third_task() {
    clear
    echo -e "${bold}Заходим в ./lab-3/...\n"
    cd ./lab-3
    read -p "Введи путь до конфигурационного файла со значениями n, c1, c2, c3, eps: ${basic}" cfgpath
    go run main.go -cfg ${cfgpath}
}

res=$(select_task)
case $res in
    q)
        exit;;
    1)
        launch_first_task;;
    2)
        launch_second_task;;
    3)
        launch_third_task;;    
esac