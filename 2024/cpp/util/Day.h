#ifndef DAY_H
#define DAY_H

#include <iostream>
#include <string>
#include <fstream>

template <typename T>
class Day {
private:
    std::string   m_inputFolder { "./src/inputs/" };

protected:
    std::ifstream m_input {};
    std::string   m_day {};

public:
    Day(std::string day)
    : m_day { day } {
        m_input.open(m_inputFolder + day + ".txt");
        if (!m_input) {
            std::cerr << "failed to open file " << m_inputFolder << day << ".txt";
        }
    };

    virtual ~Day() {
        m_input.close();
    };

    void run();

protected:
    virtual void initializeInput() = 0;
    virtual T p1() = 0;
    virtual T p2() = 0;

    inline std::ifstream& getInput() {
        return m_input;
    }
};

#endif
