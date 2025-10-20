#ifndef OPTIONAL_H
#define OPTIONAL_H

#include <memory>

template<typename T>
class Optional {
private:
    std::unique_ptr<T> m_value {};

public:
    Optional() = default;
    Optional(T value): m_value(std::make_unique<T>(std::move(value))) {};

    inline bool hasValue() const { return m_value != nullptr; };
    inline T& getValue() { return *m_value; };
    inline const T& getValue() const { return *m_value; };
    inline void setValue(T val) { m_value = std::make_unique<T>(std::move(val)); }
    inline void reset() { m_value.reset(); };
};

#endif
