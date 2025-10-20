#include "Level.h"

#include <iostream>
#include <vector>
#include <string>
#include <sstream>


/**
 * 
 * Is Safe returns a boolean, indicating if
 * (you guessed id) a level is safe.
 * 
 * A level is safe when it abides by 2 rules:
 * 
 * 1. The numbers of a level must be sorted
 *      (either asc or desc). E.g. steps are always
 *      incrementing, or decreasing, no mix&match.
 * 
 * 2. Distance between steps is 1, 2 or 3.
 * 
 */
bool Level::isSafe()
{
    if (m_steps.at(0) == m_steps.at(1)) return false;

    bool isIncreasing { m_steps.at(0) < m_steps.at(1) };

    for (size_t i { 0 }; i < m_steps.size()-1; ++i) {
        int i1 { m_steps.at(i)   };
        int i2 { m_steps.at(i+1) };

        if (!isStepSafe(i1, i2, isIncreasing)) return false;
    }

    return true;
}

/**
 * A level can be dampened, meaning a single number
 * may be left out in order to achieve a safe level.
 */
bool Level::isSafeWithDampener()
{
    for (size_t j { 0 }; j < m_steps.size(); ++j) {
        Level clone { m_steps };
        clone.m_steps.erase(clone.m_steps.begin() + static_cast<int>(j));
        if (clone.isSafe()) return true;
    }
    return false;
}

/**
 * Singular check done by isLevelSafe.
 * Checks if step is in right direction, and small enough.
 */
bool Level::isStepSafe(int x, int y, bool isIncreasing) 
{
    int diff { isIncreasing ? y - x : x - y };
    if (diff <= 0 || diff > 3) return false;
    return true;
}


std::ifstream& operator>>(std::ifstream& input, Level& val)
{
    val.m_steps.clear();

    std::string line;
    if (std::getline(input, line)) {
        std::istringstream iss(line);
        int x {};
        while (iss >> x) val.m_steps.push_back(x);
    }

    return input;
};

std::ostream& operator<<(std::ostream& out, const Level& val)
{
    for (size_t i { 0 }; i < val.m_steps.size(); ++i)
    {
        out << val.m_steps[i];
        if (i < val.m_steps.size()-1) out << ", ";
    }
    return out;
}
