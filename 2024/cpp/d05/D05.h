#ifndef D05_H
#define D05_H

#include "SortRule.h"
#include "BookUpdates.h"
#include "../util/Day.h"

#include <vector>
#include <map>

class D05 : Day<int>
{
private:
    std::vector<BookUpdates> m_updates { };
    std::map<int, SortRule> m_rules { };

public:
    D05(): Day { "05" } { };
    void run() { Day::run(); };

protected:
    void initializeInput();
    int p1();
    int p2();
};

#endif
