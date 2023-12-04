#include <stdio.h>
#include <stdlib.h>
#include "parser.h"

// returns:
// 0 on success
// -1 on error (lol imagine having more than one error code)
int try_parse(FILE * restrict f)
{
    fseek(f, 0, SEEK_END);
    size_t fsize = ftell(f);
    rewind(f);

    char *fcontent = (char*) malloc(sizeof(char) * (fsize + 1));
    fread(fcontent, 1, fsize, f);
    fcontent[fsize] = 0;

    parse_input_string(fcontent);
    free(fcontent);
    return 0;
}

int main()
{
    FILE *f = fopen("input.txt", "r");
    if (f == NULL) {
        puts("Cannot open the file");
        return EXIT_FAILURE;
    }

    if (try_parse(f)) {
        puts("Cannot process input, this is quite sad");
        return EXIT_FAILURE;
    }
    fclose(f);
    return EXIT_SUCCESS;
}
