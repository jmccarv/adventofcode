// Day 4 with input set bundled in the executable

#include <conio.h>
#include <time.h>
#include <string.h>
#include <stdio.h>

#include "input.h"

#define MAX_STACKS 9

// used by input.c
const char *infile_prefix = "d6.";

int solve(char *line, int nr_uniq, int start);

void main(void) {
    clock_t t0, t1;
    unsigned char x,y;
    char *data;

    screensize(&x, &y);
    if (x == 80) fast();

    t0 = clock();

    data = strtok(load_input(get_input_file()), "\n");
    while (data) {
        solve(data, 14, solve(data, 4, 0)-4);
        data = strtok(NULL, "\n");
    }

    t1 = clock();
    cprintf("Runtime: %ld\n", t1-t0);
}

int solve(char *line, int nr_uniq, int start) {
    int j, i;
    int len = strlen(line);
    int nr = 0;
    char seen['z'+1];


    for (i = start + nr_uniq - 1; i < len && nr == 0; i = j + nr_uniq) {
        bzero(seen + 'a', 26);
        nr = i;
        for (j = i; j >= i-(nr_uniq-1); --j) {
            if (seen[line[j]]) {
                nr = 0;
                break;
            }
            seen[line[j]] = 1;
        }
    }

    cprintf("%2d %4d\r\n", nr_uniq, nr+1);
    return nr+1;
}
