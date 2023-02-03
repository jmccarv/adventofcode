// Day 4 with input set bundled in the executable

#include <conio.h>
#include <time.h>
#include <string.h>
#include <stdio.h>

#include "input.h"

#define MAX_STACKS 9

// used by input.c
const char *infile_prefix = "d5.";

char *parse_initial_state(char *data);
void dump(char stacks[MAX_STACKS][50], char *sp[MAX_STACKS]);
void solve(char *moves);
void disp_top(unsigned char *sp[MAX_STACKS]);

char nr_stacks = 0;
char p1stacks[MAX_STACKS][50];
char p2stacks[MAX_STACKS][50];
char *p1sp[MAX_STACKS];
char *p2sp[MAX_STACKS];

void main(void) {
    unsigned char x,y;

    screensize(&x, &y);
    if (x == 80) fast();

    solve(parse_initial_state(load_input(get_input_file())));
    /*
    cprintf("nr_stacks: %d\r\n", nr_stacks);
    cprintf("%c\r\n", *data);
    dump(p1stacks, p1sp);
    */
}

// returns a pointer to where the move instructions start
char *parse_initial_state(char *p) {
    char r = 0;
    char s, i;
    char *q = p;

    if (!p) return NULL;

    for (i = 0; i < MAX_STACKS; i++) {
        p1sp[i] = 0;
        p2sp[i] = 0;
    }

    while (q = strchr(q+1, '\n')) {
        if (q[1] == '\n') {
            break;
        }
        ++r;
    }
    --r;
    // r is now the height of the tallest stack, starting from 0

    // We need to know our maximum height (max stack depth)
    while (p[0] && p[0] != 'm') {
        for (i = 0; p[i]; ++i) {
            if (p[i] >= 'A' && p[i] <= 'Z') {
                s = (i-1)/4;
                p1stacks[s][r] = p[i];
                p2stacks[s][r] = p[i];
                if (! p1sp[s]) {
                    //cprintf("r=%d i=%d s=%d %c\r\n", r, i, s+1, p1stacks[s][r]);
                    p1sp[s] = p1stacks[s] + r+1;
                    p2sp[s] = p2stacks[s] + r+1;
                }

            } else if (p[i] >= '1' && p[i] <= '9') {
                nr_stacks = p[i] - '1' + 1;

            } else if (p[i] == '\n') {
                ++i;
                break;
            }
        }
        p += i;
        --r;
    }
    return p;
}


void solve(char *moves) {
    clock_t t0, t1;
    unsigned int i, j;
    unsigned int nr, from, to;

    if (!moves) return;

    t0 = clock();

    while(*moves) {
        if (3 != sscanf(moves, "move %d from %d to %d", &nr, &from, &to)) {
            cprintf("Error in input :(\r\n");
            return;
        }

        //cprintf("%d from %d to %d\r\n", nr, from, to);

        --from;
        --to;

        // part 1
        for (j = 0; j < nr; ++j) {
            *p1sp[to]++ = *--p1sp[from];
        }

        // part2
        memcpy(p2sp[to], p2sp[from]-nr, nr);
        p2sp[from] -= nr;
        p2sp[to] += nr;

        for (moves += 18; *moves && *moves != '\n'; ++moves);
        if (*moves) ++moves;
        putchar('.');
    }
    cprintf("\r\n");

    cprintf("Part 1\r\n");
    disp_top(p1sp);

    cprintf("\r\nPart 2\r\n");
    disp_top(p2sp);

    t1 = clock();
    cprintf("Runtime: %ld\n", t1-t0);
}

void disp_top(unsigned char *sp[MAX_STACKS]) {
    unsigned char i;
    for (i = 0; i < nr_stacks; ++i) {
        cputc(*(sp[i]-1));
    }
    cprintf("\r\n");
}

void dump(char stacks[MAX_STACKS][50], char *sp[MAX_STACKS]) {
    char i, j;
    char nr[MAX_STACKS];
    char max = 0;

    for (i = 0; i < nr_stacks; i++) {
        nr[i] = sp[i] - stacks[i];
        if (nr[i] > max) max = nr[i];
    }

    for (j = max; j > 0; --j) {
        for (i = 0; i < nr_stacks; ++i) {
            if (j <= nr[i]) {
                cputc(stacks[i][j-1]);
            } else {
                cputc(' ');
            }
        }
        cprintf("\r\n");
    }
    cprintf("\r\n");
}
