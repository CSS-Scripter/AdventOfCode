#ifndef SORT_RULE_H
#define SORT_RULE_H

#include <vector>

class SortRule {
private:
    std::vector<int> m_lower { };

public:
    SortRule() = default;

    bool isLower(int v) const;
    void add(int v);
};

#endif
