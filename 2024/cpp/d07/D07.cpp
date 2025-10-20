#include "D07.h"
#include "Equation.h"

void D07::initializeInput() {
    Equation e { };
    while (m_input >> e)
        m_equations.push_back(e);
};

uint64_t D07::p1() {
    uint64_t total { };
    for (auto e : m_equations) {
        if (e.hasSolution())
            total += e.getResult();
    }
    return total;
};

uint64_t D07::p2() {
    uint64_t total { };
    for (auto e : m_equations) {
        if (e.hasSolutionTwo())
            total += e.getResult();
    }
    return total;
};
