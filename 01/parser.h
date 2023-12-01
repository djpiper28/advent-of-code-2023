#pragma once

int parse_input_string(const char *input);
typedef struct current_line_t {
    int first;
    int last;
    int total;
} current_line_t;
