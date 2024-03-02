#pragma once

#include "task.hpp"

Packages counting_sort(const Packages &packages) {
    int M = 0;
    for (const auto &package : packages) {
        M = std::max(M, package.time);
    }
    std::vector<int> count(M + 1, 0);
    for (const auto &package : packages) {
        ++count[package.time];
    }
    for (int i = 1; i <= M; ++i) {
        count[i] += count[i - 1];
    }
    Packages sorted(packages.size());
    for (int i = (int)packages.size() - 1; i >= 0; --i) {
        sorted[count[packages[i].time] - 1] = packages[i];
        --count[packages[i].time];
    }
    return sorted;
}

double nfdh(Packages *packages, int n) {
    double Ts = 0.0;
    int level = 0;
    std::vector<int> h;
    h.push_back(0);
    std::vector<int> w;
    w.push_back(0);
    auto sorted = counting_sort(*packages);
    std::reverse(sorted.begin(), sorted.end());
    h[level] = sorted[0].time;
    w[level] = sorted[0].rank;
    for (std::size_t i = 1; i < sorted.size(); ++i) {
        if (n - w[level] >= sorted[i].rank) {
            w[level] += sorted[i].rank;
        } else {
            ++level;
            w.push_back(sorted[i].rank);
            h.push_back(h[level - 1] + sorted[i].time);
        }
    }
    Ts = h.at(level);
    *packages = sorted;
    return Ts;
}