#include "D08.h"

#include "../util/ListUtil.h"

#include <string>

void D08::initializeInput() {
    size_t y { 0 };
    std::string line { };

    while (std::getline(m_input, line))
    {
        m_boundsX = line.size();
        for (size_t x { 0 }; x < line.size(); ++x)
        {
            char c { line.at(x) };
            if (c == '.') continue;

            if (m_satelites.find(c) == m_satelites.end())
                m_satelites.insert(std::map<char, std::vector<Pair<size_t>>>::value_type(c, {}));

            m_satelites[c].push_back(Pair<size_t>{x, y});
        }
        ++y;
    }
    m_boundsY = y;
}

int D08::p1() {
    std::vector<Pair<size_t>> antinodes { };
    for (auto c : m_satelites) {
        std::vector<Pair<size_t>> sAntinodes { findAntinodes(c.second) };
        for (auto an : sAntinodes) {
            antinodes.push_back(an);
        }
    }

    return static_cast<int>(ListUtil::unique(antinodes).size());
}

int D08::p2() {
    std::vector<Pair<size_t>> antinodes { };
    for (auto c : m_satelites) {
        if (c.second.size() <= 0) continue;

        std::vector<Pair<size_t>> sAntinodes { findLinearAntinodes(c.second) };
        for (auto an : sAntinodes) {
            antinodes.push_back(an);
        }
        for (auto s : c.second) {
            antinodes.push_back(s);
        }
    }

    return static_cast<int>(ListUtil::unique(antinodes).size());
}

std::vector<Pair<size_t>> D08::findAntinodes(std::vector<Pair<size_t>> satelites)
{
    std::vector<Pair<size_t>> antinodes { };
    for (size_t i { }; i < satelites.size()-1; ++i)
    {
        Pair<size_t> s1 { satelites.at(i) };
        for (size_t j { i+1 }; j < satelites.size(); ++j)
        {
            Pair<size_t> s2 { satelites.at(j) };

            Optional<Pair<size_t>> an1 { findAntinode(s1, s2) };
            Optional<Pair<size_t>> an2 { findAntinode(s2, s1) };

            if (an1.hasValue()) antinodes.push_back(an1.getValue());
            if (an2.hasValue()) antinodes.push_back(an2.getValue());
        }
    }

    return antinodes;
}

Optional<Pair<size_t>> D08::findAntinode(Pair<size_t> s1, Pair<size_t> s2)
{
    int dx { static_cast<int>(s1.first) - static_cast<int>(s2.first) };
    int dy { static_cast<int>(s1.second) - static_cast<int>(s2.second) };

    int x { static_cast<int>(s1.first) + dx };
    int y { static_cast<int>(s1.second) + dy };

    if (x < 0 || static_cast<size_t>(x) >= m_boundsX) return Optional<Pair<size_t>>{};
    if (y < 0 || static_cast<size_t>(y) >= m_boundsY) return Optional<Pair<size_t>>{};

    return Optional<Pair<size_t>>{{
        static_cast<size_t>(x),
        static_cast<size_t>(y)
    }};
}


std::vector<Pair<size_t>> D08::findLinearAntinodes(std::vector<Pair<size_t>> satelites)
{
    std::vector<Pair<size_t>> antinodes { };
    for (size_t i { }; i < satelites.size()-1; ++i)
    {
        Pair<size_t> s1 { satelites.at(i) };

        for (size_t j { i+1 }; j < satelites.size(); ++j)
        {
            Pair<size_t> s2 { satelites.at(j) };

            Pair<size_t> an1 { s1 };
            Pair<size_t> an2 { s2 };
            Optional<Pair<size_t>> an { findAntinode(an1, an2) };
            while (an.hasValue()) {
                antinodes.push_back(an.getValue());
                an2 = an1;
                an1 = an.getValue();
                an = findAntinode(an1, an2);
            }

            an1 = s2;
            an2 = s1;
            an = findAntinode(an1, an2);
            while (an.hasValue()) {
                antinodes.push_back(an.getValue());
                an2 = an1;
                an1 = an.getValue();
                an = findAntinode(an1, an2);
            }
        }
    }

    return antinodes;
}


