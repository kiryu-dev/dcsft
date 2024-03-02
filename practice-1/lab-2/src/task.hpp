#pragma once

#include <algorithm>
#include <vector>

struct Task {
    int time;
    int rank;
};

using Tasks = std::vector<Task>;

struct Package {
    Tasks tasks;
    int time;
    int rank;

    bool operator==(const Task &other) const {
        return this->rank == other.rank;
    }
};

using Packages = std::vector<Package>;

Packages make_packages(const Tasks &tasks) {
    Packages result;
    const auto mx_task = std::max_element(
        tasks.begin(), tasks.end(), [](const auto &lhs, const auto &rhs) {
            return lhs.time < rhs.time;
        });
    const auto theta = 10 * mx_task->time;
    for (const auto &task : tasks) {
        const auto iter = std::find(result.begin(), result.end(), task);
        if (iter == result.end()) {
            Tasks t;
            t.push_back(task);
            result.push_back({t, theta, task.rank});
        } else if (iter->time + task.time <= theta) {
            result[std::distance(result.begin(), iter)].tasks.push_back(task);
            // result[std::distance(result.begin(), iter)].time = theta;
        }
    }
    return result;
}
