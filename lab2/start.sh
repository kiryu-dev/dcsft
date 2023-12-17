#!/bin/bash

touch data.csv
echo task_count,nfdh,ffdh > data.csv

# ./main data/task0n1024
# ./main data/task1n1024
# ./main data/task2n1024
# ./main data/task3n1024
# ./main data/task4n1024
# ./main data/task5n1024
# ./main data/task6n1024
# ./main data/task7n1024
# ./main data/task8n1024
# ./main data/task9n1024

# ./main data/task0n4096
# ./main data/task1n4096
# ./main data/task2n4096
# ./main data/task3n4096
# ./main data/task4n4096
# ./main data/task5n4096
# ./main data/task6n4096
# ./main data/task7n4096
# ./main data/task8n4096
# ./main data/task9n4096

./main data/parse0n1152
./main data/parse1n1152
./main data/parse2n1152