#include <iostream>
#include <mpi.h>
#include <vector>

int main(int argc, char *argv[]) {
    int commsize, rank;
    MPI_Init(&argc, &argv);
    MPI_Comm_size(MPI_COMM_WORLD, &commsize);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);
    const auto n = 1024*1024;
    std::vector<std::uint8_t> read_buf(n);
    std::vector<std::uint8_t> send_buf(n, 88);
    std::vector<MPI_Request> requests(2*commsize);
    std::vector<MPI_Status> stats(2*commsize);
    auto i = rank;
    auto t = MPI_Wtime();
    do {
        MPI_Isend(send_buf.data(), n, MPI_UINT8_T, i, 0, MPI_COMM_WORLD, &requests[i]);
        MPI_Irecv(read_buf.data(), n, MPI_UINT8_T, i, 0, MPI_COMM_WORLD, &requests[commsize + i]);
        i = (i + 1) % commsize;
    } while (i != rank);
    MPI_Waitall(2*commsize, requests.data(), stats.data());
    t = MPI_Wtime() - t;
    double result_time;
    MPI_Reduce(&t, &result_time, 1, MPI_DOUBLE, MPI_MAX, 0, MPI_COMM_WORLD);
    if (rank == 0) {
        std::cout << "elapsed time: " << result_time;
    }
    MPI_Finalize();
    return 0;
}