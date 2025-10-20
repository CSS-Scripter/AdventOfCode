#ifndef BOOK_UPDATES_H
#define BOOK_UPDATES_H

#include "SortRule.h"

#include <vector>
#include <iostream>
#include <sstream>
#include <map>

class BookUpdates {
private:
    std::vector<int> m_updates {};

public:
    BookUpdates() = default;

    void sort(const std::map<int, SortRule>& rules);
    bool isSorted(const std::map<int, SortRule>& rules);
    int getMiddlePage();

    inline bool isEmpty() { return m_updates.size() == 0; };

    friend std::istream& operator>>(std::istream& input, BookUpdates& val);
};

#endif
