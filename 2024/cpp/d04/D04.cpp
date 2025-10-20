#include "D04.h"

#include "../util/Pair.h"

#include <vector>
#include <string>
#include <string_view>


void D04::initializeInput()
{
    m_grid.clear();

    std::string line {};
    while (std::getline(m_input, line))
    {
        m_grid.push_back(line);
    }
}

int D04::p1()
{
    std::vector<Pair<int>> directions {
        {  0,  1 },
        {  1,  1 },
        {  1,  0 },
        {  1, -1 },
        {  0, -1 },
        { -1, -1 },
        { -1,  0 },
        { -1,  1 },
    };

    int total { 0 };
    for (size_t y { 0 }; y < m_grid.size(); ++y)
    {
        std::string row { m_grid.at(y) };
        for (size_t x { 0 }; x < row.size(); ++x)
        {
            if (row.at(x) != 'X') continue;

            Pair<int> from { (int)x, (int)y };
            for ( Pair<int> dir : directions )
                if (checkXMAS(from, dir)) ++total;
        }
    }
    return total;
}

int D04::p2()
{
    int total { 0 };
    for (size_t y { 0 }; y < m_grid.size(); ++y)
    {
        std::string row { m_grid.at(y) };
        for (size_t x { 0 }; x < row.size(); ++x)
        {
            if (row.at(x) != 'A') continue;
            if (checkX_MAS({ x, y })) ++total;
        }
    }
    return total;
}

bool D04::checkXMAS(Pair<int> from, Pair<int> direction)
{
    const int querySize { 4 };

    // Check if query will go out of bounds
    int maxX { from.first  + ((querySize-1) * direction.first)  };
    int maxY { from.second + ((querySize-1) * direction.second) };

    if (
        maxY < 0 ||
        maxX < 0 ||
        maxY >= static_cast<int>(m_grid.size()) ||
        maxX >= static_cast<int>(m_grid.at(0).size())
    ) return false;

    std::string_view query { "XMAS" };
    for (size_t i { 0 }; i < query.size(); ++i)
    {
        size_t x { static_cast<size_t>(from.first  + (static_cast<int>(i) * direction.first )) };
        size_t y { static_cast<size_t>(from.second + (static_cast<int>(i) * direction.second)) };

        if (m_grid.at(y).at(x) != query.at(i)) return false;
    }
    return true;
}


bool D04::checkX_MAS(Pair<size_t> from)
{
    // Check if query goes out of bounds
    if (
        from.first == 0 ||
        from.first == m_grid.size()-1 ||
        from.second == 0 ||
        from.second == m_grid.at(0).size()-1
    ) return false;

    // From needs to be an A
    if (m_grid.at(from.second).at(from.first) != 'A') return false;

    // Abbreviations : t = top, l = left, r = right, b = bottom
    const Pair<size_t> tl { from.first-1, from.second-1 };
    const Pair<size_t> tr { from.first+1, from.second-1 };
    const Pair<size_t> bl { from.first-1, from.second+1 };
    const Pair<size_t> br { from.first+1, from.second+1 };

    // Abberviation : c = char
    const char tlc { m_grid.at(tl.second).at(tl.first) };
    const char trc { m_grid.at(tr.second).at(tr.first) };
    const char blc { m_grid.at(bl.second).at(bl.first) };
    const char brc { m_grid.at(br.second).at(br.first) };

    bool posDiag { (tlc == 'M' || tlc == 'S') && (brc == 'M' || brc == 'S') && (tlc != brc) };
    bool negDiag { (trc == 'M' || trc == 'S') && (blc == 'M' || blc == 'S') && (trc != blc) };

    return posDiag && negDiag;
}
