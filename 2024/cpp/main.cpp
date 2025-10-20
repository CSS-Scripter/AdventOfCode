#include "d01/D01.h"
#include "d02/D02.h"
#include "d03/D03.h"
#include "d04/D04.h"
#include "d05/D05.h"
// #include "d06/D06.h"
// #include "d07/D07.h"
#include "d08/D08.h"
#include "d09/D09.h"
#include "d11/D11.h"

int main()
{
    D01().run();
    D02().run();
    D03().run();
    D04().run();
    D05().run();
    // D06().run(); // p2 800ms runtime, not great
    // D07().run(); // p2 8 second runtime, bad
    D08().run();
    D09().run();
    // D10().run(); // Doesn't exist yet
    D11().run();
    
    return 0;
}
