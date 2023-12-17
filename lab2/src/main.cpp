#include "nfdh.hpp"
#include "task.hpp"
#include "ffdh.hpp"

#include <fstream>
#include <iostream>
#include <sys/time.h>
#include <unistd.h>

double wtime() {
    struct timeval t;
    gettimeofday(&t, NULL);
    return (double)t.tv_sec + (double)t.tv_usec * 1E-6;
}

const std::size_t ITER_COUNT = 1000;

int main(int argc, char *argv[]) {
    if (argc < 2) {
        std::cout << "Enter filename with tasks" << std::endl;
        return -1;
    }
    std::size_t n = 0; //количество эм в системе
    std::size_t m = 0; //количество параллельных задач
    Tasks tasks; //задачи для упаковки
    double Tsh = 0.0; //нижняя граница значения целевой функции
    std::ifstream fin(argv[1]);
    fin >> n >> m;
    for (std::size_t i = 0; i < m; ++i) {
        Task task;
        fin >> task.time >> task.rank;
        Tsh += task.rank * task.time;
        tasks.push_back(task);
    }
    Tsh /= n;
    auto packages = make_packages(tasks);
    double nfdh_time = 0;
    [[maybe_unused]]double Ts_nfdh = 0;
    for (std::size_t i = 0; i < ITER_COUNT; ++i) {
        const auto t = wtime();
        /* вычисляем знач. цел. функции */
        Ts_nfdh = nfdh(&packages, n);
        nfdh_time += wtime() - t;
    }
    nfdh_time /= ITER_COUNT;
    ExtendedTasks extended_tasks;
    for (auto package : packages) {
        extended_tasks.push_back({package, 0, 0});
    }
    double ffdh_time = 0;
    [[maybe_unused]]double Ts_ffdh = 0;
    for (std::size_t i = 0; i < ITER_COUNT; ++i) {
        const auto t = wtime();
        /* вычисляем знач. цел. функции */
        Ts_ffdh = ffdh(&extended_tasks, n);
        ffdh_time += wtime() - t;
    }
    ffdh_time /= ITER_COUNT;
    std::ofstream fout("data.csv", std::ios::app);
    fout << m << "," << nfdh_time << "," << ffdh_time << "\n";
    // fout << m << "," << Ts_nfdh << "," << Ts_ffdh << "\n";
    // fout << m << "," << (Ts_nfdh - Tsh) / Tsh << "," << (Ts_ffdh - Tsh) / Tsh << "\n";
    fout.close();
    std::cout << "results were written to the file date.csv\n";
}