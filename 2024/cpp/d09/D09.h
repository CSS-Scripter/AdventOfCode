#ifndef D09_H
#define D09_h

#include "../util/Day.h"

#include "Space.h"

#include <vector>


class D09 : Day<uint64_t>
{
private:
    std::vector<int> m_inputNumbers { };
    std::vector<Space> m_spaces { };

public:
    D09() : Day{ "09" } { };
    void run() { Day::run(); };

protected:
    void initializeInput();
    uint64_t p1();
    uint64_t p2();
};

#endif
