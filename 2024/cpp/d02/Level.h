#ifndef LEVEL_H
#define LEVEL_H

#include <iostream>
#include <fstream>
#include <vector>

class Level
{
private:
    std::vector<int> m_steps { };

public:
    Level(std::vector<int> steps = { }): m_steps { steps } { };
    Level(const Level& level): m_steps { level.m_steps } { };

    bool isSafe();
    bool isSafeWithDampener();
    bool isStepSafe(int x, int y, bool isIncreasing);

    friend std::ifstream& operator>>(std::ifstream& input, Level& val);
    friend std::ostream& operator<<(std::ostream& output, const Level& val);
};



#endif
