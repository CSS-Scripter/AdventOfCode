#ifndef GUARD_H
#define GUARD_H

#include "../util/Pair.h"
#include "../util/Optional.h"

#include <vector>

class Guard {
public:
    enum Direction {
        up = 0,
        right,
        down,
        left,
        max_directions,
    };

private:
    std::vector<Pair<Pair<size_t>, Direction>> m_turnHistory { };

    int       m_x           { };
    int       m_y           { };
    int       m_origin_x    { };
    int       m_origin_y    { };
    bool      m_isInLoop    { };
    Direction m_direction   { };

public:
    Guard(int x, int y)
    : m_x           { x }
    , m_y           { y }
    , m_origin_x    { x }
    , m_origin_y    { y }
    { };

    void move();
    void turnRight();
    void reset();
    bool isInLoop() const;
    Pair<int> peek() const;
    Pair<size_t> getPosition() const;

private:
    Direction getTurnRightDirection() const;
};

#endif
