#include "BookUpdates.h"
#include "SortRule.h"

#include <math.h>

void BookUpdates::sort(const std::map<int, SortRule>& rules) {
    bool isSorted { false };
    while (!isSorted) {
        isSorted = true;
        for (size_t i { 0 }; i < m_updates.size()-1; ++i)
        {
            if (rules.find(m_updates.at(i)) == rules.end()) continue;
            SortRule r { rules.at(m_updates.at(i)) };

            if (r.isLower(m_updates.at(i+1))) {
                std::swap(m_updates[i], m_updates[i+1]);
                isSorted = false;
            }
        }
    }
};

bool BookUpdates::isSorted(const std::map<int, SortRule>& rules) {
    for (size_t i { 0 }; i < m_updates.size()-1; ++i)
    {
        if (rules.find(m_updates.at(i)) == rules.end()) continue;
        SortRule r { rules.at(m_updates.at(i)) };

        if (r.isLower(m_updates.at(i+1))) {
            return false;
            std::swap(m_updates[i], m_updates[i+1]);
        }
    }
    return true;
};

int BookUpdates::getMiddlePage() {
    int middle { static_cast<int>(ceil(double(m_updates.size()) / 2))-1 };
    return m_updates.at(static_cast<size_t>(middle));
}

std::istream& operator>>(std::istream& input, BookUpdates& val)
{
    val.m_updates.clear();
    std::string line;
    if (std::getline(input, line)) {
        std::istringstream iss(line);
        int x {};
        char c {};
        while (iss >> x) {
            val.m_updates.push_back(x);
            iss >> c;
        }
    }

    return input;
}
