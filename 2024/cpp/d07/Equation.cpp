#include "Equation.h"

#include "../util/ListUtil.h"

#include <iostream>
#include <vector>
#include <string>
#include <sstream>
#include <cstdint>


bool Equation::hasSolution()
{
    // Lowest possible number is sum
    uint64_t lower { getLowerBound() };
    if (lower >= m_result) return lower == m_result;

    // Higher possible number is product
    uint64_t upper { getUpperBound() };
    if (upper <= m_result) return upper == m_result;

    if (m_values.size() == 2) return false; // should have been found at this point

    std::vector<std::string> operations {
        ListUtil::createCombinations(std::vector<char>{'+', '*'}, m_values.size()-1)
    };

    for (auto operators : operations)
    {
        uint64_t total { m_values.at(0) };
        for (size_t i { 0 }; i < m_values.size()-1; ++i)
        {
            char op { operators.at(i) };
            if (op == '+')
                total += m_values.at(i+1);
            else if (op == '*')
                total *= m_values.at(i+1);
            else
                std::cerr << "unknown operator found: " << op << '\n';

            if (total > m_result) break;
        }

        if (total == m_result) {
            return true;
        }
    }

    return false;
}


bool Equation::hasSolutionTwo() {
    // Lowest possible number is sum
    uint64_t lower { getLowerBound() };
    if (lower >= m_result) return lower == m_result;

    // Higher possible number is product
    uint64_t upper { getUpperBoundP2() };
    if (upper <= m_result) return upper == m_result;

    std::vector<std::string> operations {
        ListUtil::createCombinations(std::vector<char>{'+', '*', '|'}, m_values.size()-1)
    };

    for (auto operators : operations)
    {
        uint64_t total { m_values.at(0) };
        for (size_t i { 0 }; i < m_values.size()-1; ++i)
        {
            char op { operators.at(i) };
            if (op == '+')
                total += m_values.at(i+1);
            else if (op == '*')
                total *= m_values.at(i+1);
            else if (op == '|')
                total = concatInt(total, m_values.at(i+1));
            else
                std::cerr << "unknown operator found: " << op << '\n';

            if (total > m_result) break;
        }

        if (total == m_result) {
            return true;
        }
    }

    return false;
}


uint64_t Equation::concatInt(uint64_t x, uint64_t y)
{
    uint64_t base { 10 };
    uint64_t acc { y };
    while (acc > 9) {
        base *= 10;
        acc /= 10;
    }

    return (x * base) + y;
}


uint64_t Equation::getLowerBound()
{
    uint64_t total { 0 };
    for (uint64_t v : m_values) {
        total = v <= 1 ? total * v : total + v;
    }

    return total;
}

uint64_t Equation::getUpperBound()
{
    uint64_t total { 1 };
    for (uint64_t v : m_values)
            total = v <= 1 ? total + v : total * v;

    return total;
}

uint64_t Equation::getUpperBoundP2()
{
    uint64_t total { 1 };
    for (uint64_t v : m_values)
        total = concatInt(total, v);

    return total;
}


std::ostream& operator<<(std::ostream& out, Equation val)
{
    out << val.m_result << ": ";
    for (auto v : val.m_values)
        out << v << " ";

    return out;
}


std::istream& operator>>(std::istream& in, Equation& val)
{
    val.m_values.clear();

    std::string line { };
    if (std::getline(in, line)) {
        std::istringstream iss(line);
        
        char s { }; // separator (:)
        iss >> val.m_result >> s;

        uint64_t x { };
        while (iss >> x) val.m_values.push_back(x);
    }

    return in;
}
