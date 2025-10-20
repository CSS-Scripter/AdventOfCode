#include "SortRule.h"

bool SortRule::isLower(int v) const {
    for (int l : m_lower) {
        if (l == v) return true;
    }
    return false;
}

void SortRule::add(int v) {
    m_lower.push_back(v);
}
