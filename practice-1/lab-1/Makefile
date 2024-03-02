CPP=mpicxx
CFLAGS=-O0 -Wall -Wextra -Werror -pedantic -std=c++17
SOURCES=src/main.cpp
EXEC=mpiexec
EXECUTABLE=bin/main

.PHONY: build
build:
		@$(CPP) $(CFLAGS) $(SOURCES) -o $(EXECUTABLE)

.PHONY: run
run: build
		@${EXEC} ./$(EXECUTABLE)

.DEFAULT_GOAL := run