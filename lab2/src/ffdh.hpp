#pragma once

#include <cmath>

#include "task.hpp"

struct Node {
    Node* left;
    Node* right;
    Node* top;
    int type;
    int start_time;
    int lvl_size;
};

Node* create_node(Node* top, int i, int depth, int n) {
    auto *ptr = new Node;
	if (i != 0) {
        ptr->top = top;
    }
	ptr->lvl_size = n;
	if (i < depth) {
		ptr->left = create_node(ptr, i+1, depth, n);
		ptr->right = create_node(ptr, i+1, depth, n);
		ptr->type = 0;
	} else {
		ptr->start_time = -1;
		ptr->type = 1;
	}
	return ptr;
}

Node* find_node(Node* root, int lvl_size, int depth) {
    Node *ptr = root, *b_ptr;
	do {
		if (ptr->left->lvl_size < lvl_size) {
            ptr = ptr->right;
        } else {
            ptr = ptr->left;
        }
	} while (--depth);
	ptr->lvl_size -= lvl_size;
	b_ptr = ptr;
	while (ptr != root) {
		if (ptr->top->left->lvl_size >= ptr->top->right->lvl_size) {
            ptr->top->lvl_size = ptr->top->left->lvl_size;
        } else {
            ptr->top->lvl_size = ptr->top->right->lvl_size;
        }
		ptr = ptr->top;
	}
	return b_ptr;
}

struct ExtendedTask {
    Package job;
    double start_time;
    int position;
};

using ExtendedTasks = std::vector<ExtendedTask>;

double ffdh(ExtendedTasks *task_ptr, int n) {
    double Ts = 0.0;
	int depth;
	depth = std::ceil(log((*task_ptr).size()) / log(2));
	auto root = create_node(NULL, 0, depth, n);
	for (std::size_t i = 0; i < (*task_ptr).size(); ++i) {
		auto ptr = find_node(root, (*task_ptr)[i].job.rank, depth);
		if (ptr->start_time == -1) {
			ptr->start_time = Ts;
			Ts += (*task_ptr)[i].job.time;
		}
		(*task_ptr)[i].start_time = ptr->start_time;
		(*task_ptr)[i].position = (n - ptr->lvl_size - (*task_ptr)[i].job.rank);
	}
	return Ts;
}