#include "Day.h"
#include "Clock.h"

#include <iomanip>

template <typename T>
void Day<T>::run() {
    std::cout << std::setprecision(3) << std::fixed;
    std::cout << "=== Day " << m_day << " ================= ";
    initializeInput();

    Clock c { };
    c.start();
    T p1Result { p1() };
    double p1MS { c.elapsed() };

    T p2Result { p2() };
    double p2MS { c.elapsed() - p1MS};
    std::cout << "(" << c.elapsed() << "ms)\n";

    std::cout << "[" << p1MS << "ms] Part 1: " << p1Result << '\n';
    std::cout << "[" << p2MS << "ms] Part 2: " << p2Result << "\n\n";
}

template void Day<int>::run();
template void Day<uint64_t>::run();
