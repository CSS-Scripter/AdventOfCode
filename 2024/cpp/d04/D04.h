#ifndef D04_H
#define D04_H

#include "../util/Day.h"
#include "../util/Pair.h"

#include <vector>
#include <string>

class D04 : Day<int>
{
private:
    std::vector<std::string> m_grid {};

public:
    D04(): Day { "04" } {};
    void run() { Day::run(); };

protected:
    void initializeInput();
    int p1();
    int p2();

private:
    bool checkXMAS(Pair<int> from, Pair<int> direction);
    bool checkX_MAS(Pair<size_t> from);
};

#endif
