#ifndef SPACE_H
#define SPACE_H

#include <cstdint>
#include <iostream>

struct Space {
    enum SpaceType { 
        empty,
        file,
        max_types
    };

    SpaceType type { };
    int size { };
    size_t id { };
};

inline std::ostream& operator<<(std::ostream& out, Space space)
{
    for (int i { 0 }; i < space.size; ++i)
    {
        if (space.type == Space::empty)
            out << '.';
        else
            out << space.id;
    }
    return out;
}

#endif
