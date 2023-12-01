%option noyywrap noinput nounput case-insensitive

%{
#include "part2.tab.h"
%}

%%
\n { return NEW_LINE; }

"one" { return ONE; }
"two" { return TWO; }
"three" { return THREE; }
"four" { return FOUR; }
"five" { return FIVE; }
"six" { return SIX; }
"seven" { return SEVEN; }
"eight" { return EIGHT; }
"nine" { return NINE; }
[0-9] { return DIGIT; }
[a-z] { return OTHER; }
%%
