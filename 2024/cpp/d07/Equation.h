#ifndef EQUATION_H
#define EQUATION_H

#include <cstdint>
#include <vector>
#include <iostream>

class Equation
{
private:
    uint64_t m_result { };
    std::vector<uint64_t> m_values { };

public:
    Equation() = default;

    bool hasSolution();
    bool hasSolutionTwo();

    inline uint64_t getResult() { return m_result; };
    
private:
    uint64_t getLowerBound();
    uint64_t getUpperBound();
    uint64_t getUpperBoundP2();
    uint64_t concatInt(uint64_t x, uint64_t y);

public:
    friend std::ostream& operator<<(std::ostream& out, Equation val);
    friend std::istream& operator>>(std::istream& in, Equation& val);
};

#endif
