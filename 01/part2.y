%{
#include <stdio.h>
#include <string.h>
#include "lex.yy.h"
#include "parser.h"

static void yyerror(current_line_t *__ret, const char *s)
{
    printf("Parse error: %s\n", s);
}

void handle_number(current_line_t *line, int number)
{
    line->last = number;
    if (line->first == 0) {
        line->first = number;
    }
    printf("number: \t%d fist: \t%d last: \t%d total: \t%d\n", number, line->first, line->last, line->total);
}

void handle_end_of_input(current_line_t *line)
{
    line->total += (line->first * 10) + line->last;
    line->first = line->last = 0;
}
%}

%code requires {
    #include "parser.h"
}

%define parse.error verbose
%parse-param {current_line_t *line}

%left DIGIT ONE TWO THREE FOUR FIVE SIX SEVEN EIGHT NINE
%token OTHER NEW_LINE

%start INPUT

%%
INPUT : INPUT NEW_LINE LINE {handle_end_of_input(line);}
      | INPUT NEW_LINE
      | LINE {handle_end_of_input(line);}
      | %empty;

VALID_DIGIT: DIGIT {handle_number(line, yytext[0] - '0');} 
           | ONE {handle_number(line, 1);} 
           | TWO {handle_number(line, 2);} 
           | THREE {handle_number(line, 3);}
           | FOUR {handle_number(line, 4);}
           | FIVE {handle_number(line, 5);}
           | SIX {handle_number(line, 6);}
           | SEVEN {handle_number(line, 7);}
           | EIGHT {handle_number(line, 8);}
           | NINE {handle_number(line, 9);};

LETTER: VALID_DIGIT
      | OTHER;

LINE: LINE LETTER | LETTER;
%%

int parse_input_string(const char* input_string)
{
    current_line_t line;
    memset(&line, 0, sizeof(line));

    YY_BUFFER_STATE input_buffer = yy_scan_string(input_string);
    int result = yyparse(&line);
    yy_delete_buffer(input_buffer);

    if (result == 0) {
        printf("Total: %d\n", line.total);
    } else {
        puts("Cannot parse input");
    }
    return result;
}
