#ifndef D02_H
#define D02_H

#include "../util/Day.h"
#include "Level.h"

#include <vector>

class D02 : Day<int>
{
private:
    std::vector<Level> m_levels { };

public:
    D02() : Day{ "02" } { };
    inline void run() { Day::run(); };

protected:
    void initializeInput();
    int p1();
    int p2();
};

#endif
