%option noyywrap noinput nounput case-insensitive

%{
#include "part1.tab.h"
%}

%%
\n { return NEW_LINE; }
[ \t]+ { return WHITESPACE; }
[0-9]+ { return NUMBER; }

"game" { return GAME; }
"blue" { return BLUE; }
"red" { return RED; }
"green" { return GREEN; }

"," { return COMMA; }
";" { return SEMI_COLON; }
":" { return COLON; }
%%
