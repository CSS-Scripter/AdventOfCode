#include "D03.h"

#include "../util/Optional.h"
#include "../util/Pair.h"

#include <vector>
#include <string>
#include <string_view>


void D03::initializeInput() {
    m_lines.clear();
    std::string line { };
    
    while (std::getline(m_input, line)) m_lines.push_back(line);
};

int D03::p1() {
    int total { 0 };
    for (std::string line : m_lines)
    {
        for (size_t i = 0; i < line.size(); i++)
        {
            if (line.at(i) != 'm') continue;
            Optional<Pair<int>> mul = findMul(line, i);
            if (mul.hasValue()) {
                int x { mul.getValue().first };
                int y { mul.getValue().second };

                total += (x * y);
            }
        }
    }
    return total;
};

int D03::p2() {
    bool enabled { true };
    int total { 0 };
    for (std::string line : m_lines)
    {
        for (size_t i = 0; i < line.size(); i++)
        {
            if (line.at(i) != 'm' && line.at(i) != 'd') continue;
            if (enabled) {
                Optional<Pair<int>> mul = findMul(line, i);
                if (mul.hasValue()) {
                    int x { mul.getValue().first };
                    int y { mul.getValue().second };

                    total += (x * y);
                }

                if (findSubstring(line, "don't()", i))
                    enabled = false;
            } else {
                if (findSubstring(line, "do()", i))
                    enabled = true;
            }
        }
    }
    return total;
};


bool D03::findSubstring(std::string_view line, std::string substr, size_t startAt)
{
    for (size_t j { 0 }; j < substr.size(); ++j)
    {
        if (substr.at(j) != line.at(j+startAt)) {
            return false;
        }
    }
    return true;
}

// Checks if the current position is "mul([0-999],[0-999])"
// Returns 2 digits multiplied
// Returns 0 when no mul found
Optional<Pair<int>> D03::findMul(std::string_view line, size_t i)
{
    // Check if line at pos i starts with "mul("
    std::vector<char> toFind { 'm', 'u', 'l', '(' };
    for (size_t j { 0 }; j < toFind.size(); ++j)
    {
        if (toFind.at(j) != line.at(j+i)) {
            return Optional<Pair<int>>{};
        }
    }

    if (!findSubstring(line, "mul(", i))
        return Optional<Pair<int>>{};

    // Position i + 3 will be (
    // Look for digits until ,
    // Look for digits until )
    // Encounter unexpected char? Return empty

    // Look for the first digit
    size_t startFrom { static_cast<size_t>(i + toFind.size()) };
    Optional<Pair<int>> x { findDigitUntilChar(line, startFrom, ',') };
    if (!x.hasValue()) return Optional<Pair<int>>{};

    startFrom = static_cast<size_t>(x.getValue().second + 1);
    Optional<Pair<int>> y { findDigitUntilChar(line, startFrom, ')') };
    if (!y.hasValue()) return Optional<Pair<int>>{};

    return Optional(Pair{ x.getValue().first, y.getValue().first });
}

// Return 2 ints : (value, ends-at)
Optional<Pair<int>> D03::findDigitUntilChar(std::string_view line, size_t i, char separator)
{
    std::vector<int> buff { };
    while(true) {
        char c { line.at(i) };
        if (c == separator) break;

        int x { parseInt(c) };
        if (x == EOF) return Optional<Pair<int>>();

        buff.push_back(x);
        ++i;
    }
    if (line.at(i) != separator) return Optional<Pair<int>>();

    int x { constructInt(buff) };
    if (x == EOF) return Optional<Pair<int>>();
    else return Optional<Pair<int>>(Pair<int>{ x, static_cast<int>(i) });
}

int D03::parseInt(char c) {
    if (c >= 48 && c <= 57) return (c - 48);
    return EOF;
}


int D03::constructInt(std::vector<int> buff)
{
    int total { 0 };
    for (int i : buff)
    {
        total *= 10;
        total += i;
    }

    return total;
}
