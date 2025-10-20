#include "D11.h"

#include "../util/Pair.h"

#include <cmath>

void D11::initializeInput() {
    uint64_t x { };
    while (m_input >> x) {
        if (m_stones.find(x) == m_stones.end()) {
            m_stones.insert(std::map<uint64_t, uint64_t>::value_type(x, 0));
        }

        m_stones[x]++;
    }
}

uint64_t D11::p1() {
    for (int i { 0 }; i < 25; ++i)
        blink();

    return countStones();
}

uint64_t D11::p2() {
    // We're already at blink 25 from p1
    for (int i { 25 }; i < 75; ++i)
        blink();

    return countStones();
};

void D11::blink() {
    std::map<uint64_t, uint64_t> newStones { };
    for (auto e : m_stones) {
        // If stone is 0, turn into 1
        if (e.first == 0) {
            uint64_t newStone { 1 };
            uint64_t newStoneCount { e.second };

            if (newStones.find(newStone) == newStones.end())
            {
                newStones.insert(std::map<uint64_t, uint64_t>::value_type(newStone, newStoneCount));
            } else
            {
                newStones[newStone] += newStoneCount;
            }
            continue;
        }

        // If stone is even, split into 2 stones
        uint64_t numCount { findNumCount(e.first) };
        if (numCount % 2 == 0) {
            Pair<uint64_t> stones { splitStone(e.first) };
            uint64_t count { e.second };
            if (newStones.find(stones.first) == newStones.end())
            {
                newStones.insert(std::map<uint64_t, uint64_t>::value_type(stones.first, count));
            } else
            {
                newStones[stones.first] += count;
            }

            if (newStones.find(stones.second) == newStones.end())
            {
                newStones.insert(std::map<uint64_t, uint64_t>::value_type(stones.second, count));
            } else
            {
                newStones[stones.second] += count;
            }
            continue;
        }

        // Otherwise, multiply by 2024
        uint64_t newStone { e.first * 2024 };
        if (newStones.find(newStone) == newStones.end())
        {
            newStones.insert(std::map<uint64_t, uint64_t>::value_type(newStone, e.second));
        } else
        {
            newStones[newStone] += e.second;
        }
    }

    m_stones.clear();
    m_stones = newStones;
}

uint64_t D11::findNumCount(uint64_t x) {
    uint64_t nums { 1 };
    while (x > 9) {
        nums += 1;
        x /= 10;
    }
    return nums;
}

Pair<uint64_t> D11::splitStone(uint64_t stone)
{
    uint64_t numCount { findNumCount(stone) };
    uint64_t base { static_cast<uint64_t>(pow(10, double(numCount/2))) };
    return { stone / base, stone % base };
}

uint64_t D11::countStones() {
    uint64_t total { };
    for (auto e : m_stones) {
        total += e.second;
    }
    return total;
}
