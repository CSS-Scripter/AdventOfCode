#include "Guard.h"
#include "../util/Pair.h"
#include "../util/Optional.h"
#include "../util/ListUtil.h"

Pair<int> Guard::peek() const {
    switch (m_direction)
    {
        case up:    return { m_x,   m_y-1 };
        case down:  return { m_x,   m_y+1 };
        case left:  return { m_x-1, m_y   };
        case right: return { m_x+1, m_y   };
        default:    return { m_x,   m_y   };
    }
}

Pair<size_t> Guard::getPosition() const {
    return { static_cast<size_t>(m_x), static_cast<size_t>(m_y) };
}

void Guard::move() {
    switch (m_direction)
    {
        case up:    m_y -= 1; break;
        case down:  m_y += 1; break;
        case left:  m_x -= 1; break;
        case right: m_x += 1; break;
        default: break;
    }
}

void Guard::turnRight() {
    m_direction = getTurnRightDirection();

    Pair<Pair<size_t>, Direction> turn { getPosition(), m_direction };
    m_isInLoop = m_isInLoop || ListUtil::hasElement(m_turnHistory, turn);
    m_turnHistory.push_back(turn);
}

Guard::Direction Guard::getTurnRightDirection() const {
    switch (m_direction)
    {
        case up:    return right;
        case right: return down;
        case down:  return left;
        case left:  return up;
        default:    return up;
    }
}

bool Guard::isInLoop() const {
    return m_isInLoop;
}

void Guard::reset() {
    m_turnHistory.clear();
    m_isInLoop = false;

    m_x = m_origin_x;
    m_y = m_origin_y;
    m_direction = up;
}
