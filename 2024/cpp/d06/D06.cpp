#include "D06.h"
#include "../util/Clock.h"

void D06::initializeInput() {
    m_input >> m_map;
};

int D06::p1() {
    while (m_map.step()) { };
    return m_map.countVisited();
};

int D06::p2() {
    std::vector<Pair<size_t>> steps { m_map.getSteps() };

    int total { 0 };
    for (auto step : steps) {
        m_map.reset();
        m_map.setTmpObstacle(step);
        if (isMapInLoop()) {
            ++total;
        }
    }

    return total;
};

bool D06::isMapInLoop() {
    while (m_map.step()) {
        if (m_map.isInLoop()) return true;
    }
    return false;
}
