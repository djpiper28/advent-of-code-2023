#pragma once

int parse_input_string(const char *input);

struct line_data_t {
    int max_red;
    int max_green;
    int max_blue;
    int game_id;
};

typedef struct current_line_t {
    // Alias to allow for easy memset of each line's data
    struct line_data_t data;
    int total;
    // Tempory store for the last read number
    int num;
} current_line_t;
