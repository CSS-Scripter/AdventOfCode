#include "D02.h"
#include "Level.h"


void D02::initializeInput() {
    m_levels.clear();

    Level level { };
    while (m_input >> level) {
        m_levels.push_back(level);
    }
}

int D02::p1() {
    int safeCount { 0 };
    for (Level level : m_levels) {
        if (level.isSafe()) ++safeCount;
    }
    return safeCount;
}

int D02::p2() {
    int safeCount { 0 };
    for (Level level : m_levels)
    {
        if (level.isSafeWithDampener()) safeCount++;
    }
    return safeCount;
}

