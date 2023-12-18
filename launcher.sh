#!/bin/bash

bold=$(tput bold)
normal=$(tput sgr0)

function select_task() {
    >&2 echo -e "${bold}Выбери задачу, которую хочешь решить 🤓✨🧙‍♂️🚀🔥🌟💻*ੈ✩‧₊˚\n\n${normal}"
    while true; do
        >&2 echo "${bold}1.${normal} Измерить время обмена сообщениями на разных уровнях коммуникационной сети"
        >&2 echo "${bold}2.${normal} Построить расписание решения параллельных задач на распределенной ВС"
        >&2 echo "${bold}3.${normal} Решить задачу 'Диcпeтчeр-вычиcлитeльный цeнтр'"
        >&2 echo -e "\nЖми q, чтобы выйти\n"
        read -n1 -r res
        case $res in
            [123q] ) break;;
            * ) >&2 echo -e "${bold}\nВыбери задачу по ее номеру! 🤬🤬🤬\n${normal}";;
        esac
    done
    >&2 echo
    echo $res
}

function launch_second_task() {
    clear
    echo -e "${bold}Заходим в ./lab-2/...\n"
    cd ./lab-2
    echo -e "Собираем программу...\n"
    make
    read -p "Введи путь до файла с задачами: ${normal}" filepath
    >&2 echo
    ./main ${filepath}
}

function launch_third_task() {
    clear
    echo -e "${bold}Заходим в ./lab-3/...\n"
    cd ./lab-3
    read -p "Введи путь до конфигурационного файла со значениями n, c1, c2, c3, eps: ${normal}" cfgpath
    >&2 echo
    go run main.go -cfg ${cfgpath}
}

res=$(select_task)
case $res in
    q)
        exit;;
    1)
        >&2 echo "добавить первую лабу";;
    2)
        launch_second_task;;
    3)
        launch_third_task;;    
esac