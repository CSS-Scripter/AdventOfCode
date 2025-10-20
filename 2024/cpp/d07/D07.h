#ifndef D07_H
#define D07_H

#include "../util/Day.h"
#include "../util/Pair.h"

#include "Equation.h"

#include <vector>

class D07 : Day<uint64_t>
{
private:
    std::vector<Equation> m_equations { };

public:
    D07(): Day { "07" } { };
    void run() { Day::run(); };

protected:
    void initializeInput();
    uint64_t p1();
    uint64_t p2();
};

#endif
