#ifndef D03_H
#define D03_H

#include "../util/Day.h"
#include "../util/Optional.h"
#include "../util/Pair.h"

#include <vector>
#include <string>
#include <string_view>

class D03 : Day<int>
{
private:
    std::vector<std::string> m_lines {};

public:
    D03(): Day { "03" } {};
    void run() { Day::run(); };

protected:
    int p1();
    int p2();
    void initializeInput();

private:
    bool findSubstring(std::string_view line, std::string substr, size_t startAt);
    Optional<Pair<int>> findMul(std::string_view line, size_t i);
    Optional<Pair<int>> findDigitUntilChar(std::string_view line, size_t i, char separator);
    int parseInt(char c);
    int constructInt(std::vector<int> buff);
};

#endif
