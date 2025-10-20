#ifndef D06_H
#define D06_H

#include "../util/Day.h"
#include "Map.h"

class D06 : Day<int>
{
private:
    Map m_map {};

public:
    D06(): Day { "06" } { };
    void run() { Day::run(); };

protected:
    void initializeInput();
    int p1();
    int p2();

private:
    bool isMapInLoop();
};

#endif
