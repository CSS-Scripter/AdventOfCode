#ifndef SORT_RULE_ENTRY_H
#define SORT_RULE_ENTRY_H

#include <iostream>
#include <sstream>

struct SortRuleEntry {
    int before { };
    int after  { };
};

inline std::istream& operator>>(std::istream& input, SortRuleEntry& val)
{
    char s { };
    input >> val.before >> s >> val.after;
    return input;
}

#endif
