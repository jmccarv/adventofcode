// Day 4 with input set bundled in the executable

#include <conio.h>
#include <cbm.h>
#include <time.h>
#include <string.h>

#include "day05_data.h" // generated by elves

void disp_top(unsigned char *sp[NR_STACKS]);
void dump(unsigned char stacks [NR_STACKS][50], unsigned char *sp[NR_STACKS]);

void main(void) {
    unsigned char x,y;
    clock_t t0, t1;
    unsigned int i, j;
    unsigned char *from, *to;

    t0 = clock();
    screensize(&x, &y);
    if (x == 80) fast();

    //dump(p1stacks, p1sp);
    for (i = 0; i < nr_moves; ++i) {
        // part1
        for (j = 0; j < moves.nr[i]; ++j) {
            *p1sp[moves.to[i]]++ = *--p1sp[moves.from[i]];
        }

        // part2
        memcpy(p2sp[moves.to[i]], p2sp[moves.from[i]]-moves.nr[i], moves.nr[i]);
        p2sp[moves.from[i]] -= moves.nr[i];
        p2sp[moves.to[i]] += moves.nr[i];
    }

    cprintf("Part 1\r\n");
    disp_top(p1sp);

    cprintf("\r\nPart 2\r\n");
    disp_top(p2sp);

    t1 = clock();
    cprintf("Runtime: %ld\n", t1-t0);
}

void disp_top(unsigned char *sp[NR_STACKS]) {
    unsigned char i;
    for (i = 0; i < NR_STACKS; ++i) {
        cputc(*(sp[i]-1));
    }
    cprintf("\r\n");
}

void dump(unsigned char stacks [NR_STACKS][50], unsigned char *sp[NR_STACKS]) {
    unsigned char i, j;
    unsigned char nr[NR_STACKS];
    unsigned char max = 0;

    for (i = 0; i < NR_STACKS; i++) {
        nr[i] = sp[i] - stacks[i];
        if (nr[i] > max) max = nr[i];
    }

    for (j = max; j > 0; --j) {
        for (i = 0; i < NR_STACKS; ++i) {
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
