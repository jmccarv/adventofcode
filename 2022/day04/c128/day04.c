#include <stdio.h>
#include <errno.h>
#include <string.h>
#include <conio.h>
#include <cbm.h>
#include <fcntl.h>
#include <peekpoke.h>
#include <device.h>
#include <time.h>
#include <cbm_filetype.h>
#include <stdlib.h>

#include "input.h"

void solve(char *);

void main(void) {
    unsigned char x,y;
    clock_t t0, t1;
    char *data;
    unsigned int nr;

    screensize(&x, &y);

    //cprintf("Start: DEV: %d  screen: %dx%d\r\n", dev, x,y);
    if (x == 80) fast();

    if (NULL == (data = load_input(get_input_file(INFILE_PREFIX)))) return;

    t0 = clock();
    solve(data);
    t1 = clock();

    cprintf("Solution runtime: %ld\r\n", t1-t0);
}

struct assignment {
    unsigned int min;
    unsigned int max;
};

void solve(char *data) {
    struct assignment a, b;
    int nr;
    unsigned int nr_contain = 0;
    unsigned int nr_overlap = 0;
    register char *p = data;

    while (*p) {
        // data is multiples lines of input of the form
        // aa-bb,cc-dd
        nr = sscanf(p, "%d-%d,%d-%d", &a.min, &a.max, &b.min, &b.max);
        if (nr != 4) {
            cprintf("Error in input :(\r\n");
            return;
        }

        // part1 -- does one contain the other?
        if ( (a.min <= b.min && a.max >= b.max) || (b.min <= a.min && b.max >= a.max) ) {
            putchar('*');
            ++nr_contain;

            // if one contains the other, they also overlap (part 2)
            ++nr_overlap;

        } else if ( (a.min <= b.max && a.max >= b.min) || (b.min <= a.max && b.max >= a.min) ) {
            // part2 -- does one overlap the other?
            ++nr_overlap;
            putchar('+');

        } else {
            putchar(' ');
        }

        // Find the next dataset (new line of data)
        for (p += 7; *p && *p != '\n'; ++p);
        if (*p) ++p;
    }

    cprintf("\r\n");
    cprintf("Part1: %d\r\n", nr_contain);
    cprintf("Part2: %d\r\n", nr_overlap);
}
