%{
#include <stdio.h>
#include <string.h>
#include "lex.yy.h"
#include "parser.h"
#include <sys/param.h> // MAX

static void yyerror(current_line_t *__ret, const char *s)
{
    printf("Parse error: %s\n", s);
}

#define MAX_RED 12
#define MAX_GREEN 13
#define MAX_BLUE 14
%}

%code requires {
    #include "parser.h"
}

%define parse.error verbose
%parse-param {current_line_t *line}

%left WHITESPACE NUMBER RED BLUE GREEN COLON SEMI_COLON COMMA GAME
%token OTHER NEW_LINE

%start INPUT

%%
INPUT : INPUT NEW_LINE LINE
      | INPUT NEW_LINE
      | LINE
      | %empty;

NUM: NUMBER {
       line->num = atoi(yytext);
     };

CUBE : NUM WHITESPACE GREEN {
        line->data.max_green = MAX(line->data.max_green, line->num);
     }
     | NUM WHITESPACE RED {
        line->data.max_red = MAX(line->data.max_red, line->num);
     }
     | NUM WHITESPACE BLUE {
        line->data.max_blue = MAX(line->data.max_blue, line->num);
     } ;

PICK : CUBE 
     | CUBE COMMA WHITESPACE PICK;

PICKS : PICK
      | PICK SEMI_COLON WHITESPACE PICKS;

LINE : GAME WHITESPACE NUMBER {
       memset(&line->data, 0, sizeof(line->data));
       line->data.game_id = atoi(yytext);
     } COLON WHITESPACE PICKS {
       int valid = 1;
       printf("Game %d; ", line->data.game_id);
       if (line->data.max_red > MAX_RED) {
           printf("Exceeded max red; ");
           valid = 0;
       }

       if (line->data.max_green > MAX_GREEN) {
           printf("Exceeded max green;");
           valid = 0;
       }

       if (line->data.max_blue > MAX_BLUE) {
           printf("Exceeded max blue;");
           valid = 0;
       }

       printf("Valid %d; ", valid);
       if (valid) {
           line->total += line->data.game_id;
       }

       printf("Total %d\n", line->total);
     };
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
