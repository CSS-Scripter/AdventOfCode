#ifndef CLOCK_H
#define CLOCK_H

#include <ctime>

class Clock
{
private:
    clock_t m_start { };

public:
    Clock() = default;

    inline void start() {
        m_start = clock();
    }

    inline double elapsed() {
        return static_cast<double>((clock() - m_start)) / double(CLOCKS_PER_SEC) * 1000;
    }
};

#endif
