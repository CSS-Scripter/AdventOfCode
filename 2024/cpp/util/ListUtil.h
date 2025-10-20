#ifndef LIST_UTIL_H
#define LIST_UTIL_H

#include <vector>

namespace ListUtil {
    template <typename T>
    inline int count(std::vector<T> arr, T el)
    {
        int count { 0 };
        for (size_t i { 0 }; i < arr.size(); ++i)
        {
            if (arr[i] == el) ++count;
        }
        
        return count;
    }


    template <typename T>
    inline void selectionSort(std::vector<T> arr)
    {
        for (size_t i{ 0 }; i < arr.size(); ++i)
        {
            size_t smallestIndex { i };
            for (size_t j { i+1 }; j < arr.size(); ++j)
            {
                if (arr[j] < arr[smallestIndex]) smallestIndex = j;
            }

            std::swap(arr[i], arr[smallestIndex]);
        }
    }

    template <typename T>
    inline bool hasElement(const std::vector<T>& arr, const T& toFind)
    {
        for (T el : arr)
            if (el == toFind) return true;
        
        return false;
    }

    template <typename T>
    inline std::vector<T> unique(const std::vector<T>& arr)
    {
        std::vector<T> uniqueItems {};
        for (T el : arr) {
            if (!ListUtil::hasElement(uniqueItems, el)) uniqueItems.push_back(el);
        }
        return uniqueItems;
    }


    inline void generateCombinations(const std::vector<char>& options, size_t n, std::string current, std::vector<std::string>& results) {
        if (current.size() == n) {
            results.push_back(current);
            return;
        }

        for (char option : options) {
            generateCombinations(options, n, current + option, results);
        }
    }


    inline std::vector<std::string> createCombinations(const std::vector<char>& options, size_t n) {
        std::vector<std::string> results;
        generateCombinations(options, n, "", results);
        return results;
    }


}

#endif
