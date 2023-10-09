#include <iostream>
#include <mpi.h>
#include <vector>

int main(int argc, char *argv[]) {
    std::size_t msg_size = 1024 * 1024;
    if (argc > 2 && std::string_view(argv[1]) == "--msg-size") {
        msg_size *= std::atoi(argv[2]);
    }
    int commsize, rank;
    MPI_Init(&argc, &argv);
    MPI_Comm_size(MPI_COMM_WORLD, &commsize);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);
    const std::size_t nruns = 10;
    std::vector<std::uint8_t> read_buf(msg_size);
    std::vector<std::uint8_t> send_buf(msg_size, 88);
    auto t = MPI_Wtime();
    std::vector<MPI_Request> requests(2);
    std::vector<MPI_Status> stats(2);
    for (std::size_t i = 0; i < nruns; ++i) {
        if (rank == 0){
            MPI_Isend(send_buf.data(), msg_size, MPI_UINT8_T, 1, 0, MPI_COMM_WORLD, &requests[0]);
            MPI_Irecv(read_buf.data(), msg_size, MPI_UINT8_T, 1, 0, MPI_COMM_WORLD, &requests[1]);
        } else if (rank == 1) {
            MPI_Irecv(read_buf.data(), msg_size, MPI_UINT8_T, 0, 0, MPI_COMM_WORLD, &requests[0]);
            MPI_Isend(send_buf.data(), msg_size, MPI_UINT8_T, 0, 0, MPI_COMM_WORLD, &requests[1]);
        }
        MPI_Waitall(2, requests.data(), stats.data());
    }
    t = (MPI_Wtime() - t) / nruns;
    if (rank == 0) {
        std::cout << msg_size / (1024*1024) << "," << t;
    }
    MPI_Finalize();
    return 0;
}
