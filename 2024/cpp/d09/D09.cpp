#include "D09.h"

void D09::initializeInput() {
    char x { };
    while (m_input >> x) {
        if (x < 48 || x > 57) {
            std::cerr << "input contains non-digit: " << x << '\n';
            continue;
        }

        // char - 48 for ascii to int conversion
        int v { static_cast<int>(x-48)};
        m_inputNumbers.push_back(v);
    }

    // If numbers are even, final number is space and can be omitted
    // guarantees an uneven amount of numbers
    if (m_inputNumbers.size() % 2 == 0) m_inputNumbers.pop_back();


    // Setup input for p2
    for (size_t i { }; i < m_inputNumbers.size(); ++i)
    {
        if (m_inputNumbers.at(i) == 0) continue; // skip empty spaces & files

        bool isSpace { i % 2 == 1 };
        if (isSpace)
            m_spaces.push_back(Space{
                Space::empty,
                m_inputNumbers.at(i),
            });
        else
            m_spaces.push_back(Space{
                Space::file,
                m_inputNumbers.at(i),
                (i/2),
            });
    }


    // Merge adjecant spaces
    std::vector<Space>::iterator i = m_spaces.begin();
    while (i != m_spaces.end()-1)
    {
        if (i->type == Space::empty && (i+1)->type == Space::empty)
        {
            (i+1)->size += i->size;
            m_spaces.erase(i++);
        } else {
            ++i;
        }
    }
};

uint64_t D09::p1() {
    // Approach: Shrinking window
    // - have indexes for values left and right
    // - when left is on a number, add to checksum from l indice
    // - when left is on a space, add to checksum form r indice

    uint64_t checksum { };
    size_t l { 0 };                         // Left indice
    size_t r { m_inputNumbers.size()-1 };   // Right indice
    size_t i { 0 };                         // File index

    // given input of 5, will have files 0, 1 and 2, max is 2
    // 5 / 2 -> 2.5 -> 2 (because int division)
    size_t rightFileID { (m_inputNumbers.size() / 2) };

    while (l <= r) // Continue untill shrinking window is empty
    {
        while (m_inputNumbers.at(l) == 0) ++l;
        while (m_inputNumbers.at(r) == 0) {
            r -= 2; // skip spaces
            --rightFileID;

            if (l > r) return checksum;
        }


        bool isSpace { l % 2 == 1 };
        size_t fileID { };
        if (!isSpace) {
            --m_inputNumbers[l];
            fileID = l/2;
        } else {
            --m_inputNumbers[r];
            --m_inputNumbers[l];
            fileID = rightFileID;
        }

        checksum += fileID * i;
        ++i;
    }

    return checksum;
};

uint64_t D09::p2() {
    // shit's fucked

    return 0;
};
