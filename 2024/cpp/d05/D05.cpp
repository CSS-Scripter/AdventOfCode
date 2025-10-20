#include "D05.h"
#include "SortRuleEntry.h"
#include "BookUpdates.h"

#include <sstream>

void D05::initializeInput() {
    std::vector<SortRuleEntry> ruleEntries {};
    std::string line {};
    while (std::getline(m_input, line) && !line.empty())
    {
        std::istringstream iss(line);
        SortRuleEntry entry { };
        if (iss >> entry) {
            ruleEntries.push_back(entry);
        }
    }

    for (SortRuleEntry rule : ruleEntries)
    {
        if (m_rules.find(rule.after) == m_rules.end())
            m_rules.insert(std::map<int, SortRule>::value_type(rule.after, {}));
        
        m_rules.at(rule.after).add(rule.before);
    }

    BookUpdates update {};
    while (m_input >> update && !update.isEmpty()) {
        m_updates.push_back(update);
    }
}

int D05::p1() {
    int total { };
    for (BookUpdates update : m_updates) {
        if (update.isSorted(m_rules)) total += update.getMiddlePage();
    }
    return total;
}

int D05::p2() {
    int total { };
    for (BookUpdates update : m_updates) {
        if (!update.isSorted(m_rules)) {
            update.sort(m_rules);
            total += update.getMiddlePage();
        }
    }
    return total;
}
