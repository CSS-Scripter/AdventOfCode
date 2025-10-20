#ifndef D01_H
#define D01_H

#include "../util/Day.h"

#include <vector>

class D01 : Day<int>
{
private:
    std::vector<int> l1 {};
    std::vector<int> l2 {};

public:
    D01() : Day{ "01" } { };
    inline void run() { Day::run(); };

protected:
    void initializeInput();
    int p1();
    int p2();
};

#endif
