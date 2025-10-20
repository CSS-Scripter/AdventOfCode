#ifndef D08_H
#define D08_H

#include "../util/Day.h"
#include "../util/Pair.h"
#include "../util/Optional.h"

#include <map>
#include <vector>

class D08 : Day<int>
{
private:
    std::map<char, std::vector<Pair<size_t>>> m_satelites { };
    size_t m_boundsX { };
    size_t m_boundsY { };

public:
    D08() : Day { "08" } { };
    void run() { Day::run(); };

protected:
    void initializeInput();
    int p1();
    int p2();

private:
    std::vector<Pair<size_t>> findAntinodes(std::vector<Pair<size_t>> satelites);
    Optional<Pair<size_t>> findAntinode(Pair<size_t> s1, Pair<size_t> s2);

    std::vector<Pair<size_t>> findLinearAntinodes(std::vector<Pair<size_t>> satelites);
};

#endif
