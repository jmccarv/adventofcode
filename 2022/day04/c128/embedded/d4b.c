// Day 4 with input set bundled in the executable

#include <conio.h>
#include <cbm.h>
#include <time.h>

#include "d4b_data.h" // generated by elves

void main(void) {
    unsigned char x,y;
    clock_t t0, t1;
    int i;
    unsigned int nr_contain = 0;
    unsigned int nr_overlap = 0;

    screensize(&x, &y);
    if (x == 80) fast();

    t0 = clock();
    for (i = 0; i < assn.nr; i++) {
        // part1 -- does one contain the other?
        if (   (assn.min1[i] <= assn.min2[i] && assn.max1[i] >= assn.max2[i])
            || (assn.min2[i] <= assn.min1[i] && assn.max2[i] >= assn.max1[i]) ) { 
            nr_contain++;

            // if one contains the other, they also overlap (part 2)
            nr_overlap++;

        } else if (   (assn.min1[i] <= assn.max2[i] && assn.max1[i] >= assn.min2[i])
                   || (assn.min2[i] <= assn.max1[i] && assn.max2[i] >= assn.min1[i]) ) {
            // part2 -- does one overlap the other?
            nr_overlap++;
        }
    }
    t1 = clock();

    cprintf("Part1: %d\r\n", nr_contain);
    cprintf("Part2: %d\r\n", nr_overlap);
    cprintf("Runtime: %ld\n", t1-t0);
}
