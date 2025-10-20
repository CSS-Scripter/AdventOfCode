#ifndef PAIR_H
#define PAIR_H

#include <iostream>

template<typename T, typename R = T>
struct Pair {
    T first {};
    R second {};
};

template<typename T, typename R = T>
std::ostream& operator<<(std::ostream& out, const Pair<T, R>& p)
{
    out << "(" << p.first << ", " << p.second << ")";
    return out;
}

template<typename T, typename R = T>
bool operator== (const Pair<T, R>& x, const Pair<T, R>& y)
{
    return (x.first == y.first) && (x.second == y.second);
}

template<typename T, typename R = T>
bool operator!= (const Pair<T, R>& x, const Pair<T, R>& y)
{
    return !(x == y);
}


#endif
