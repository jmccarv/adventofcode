#include <conio.h>
#include <time.h>

#include "input.h"

// used by input.c
const char *infile_prefix = "d8.";

void solve(char *data);

// size of grid -- input is a square grid of cells
char size = 0;

void main(void) {
    clock_t t0, t1;
    unsigned char x,y;
    char *data;
    register char *p, *d;

    screensize(&x, &y);
    if (x == 80) fast();

    if (NULL == (data = load_input(get_input_file()))) return;

    t0 = clock();

    // Transform our input in-place
    // Remove newlines, shift data, and convert it from character to number
    for (d = p = data; p[0]; ++p) {
        if (p[0] == '\n') {
            ++size;
        } else {
            d[0] = p[0] - '0';
            ++d;
        }
    }
    cprintf("size: %d\r\n", size);

    solve(data);
    t1 = clock();
    cprintf("Runtime: %ld\n", t1-t0);
}

// #define cell(grid, r, c) grid[size*r+c]
void solve(char *data) {
    char v, nr, r;
    register char *cp, *p;
    register char *rp = data+size;
    unsigned int vis = 0;

    // longs (32 bit) are very slow, so we'll do as little as we can with them
    // using s1 and s2 instead of just using score sped this up from ~45 seconds
    // to ~30 seconds on the full input!
    unsigned int s1, s2;
    unsigned long score;
    unsigned long p2 = 0;

    for (r=1; rp < data+(size*(size-1)); rp += size, ++r) {
        cprintf("%d\r", r);
        for (cp = rp+1; cp < rp+size-1; ++cp) {
            v = 4;
            // check up
            for (nr = 0, p = cp-size; p >= data; p -= size) {
                ++nr;
                if (*cp <= *p) {
                    --v; // fail
                    break;
                }
            }
            s1 = nr;

            // check down
            for (nr = 0, p = cp+size; p < data+size*size; p += size) {
                ++nr;
                if (*cp <= *p) {
                    --v;
                    break;
                }
            }
            s1 *= nr;

            // left
            for (nr = 0, p = cp-1; p >= rp; --p) {
                ++nr;
                if (*cp <= *p) {
                    --v;
                    break;
                }
            }
            s2 = nr;

            // right
            for (nr = 0, p = cp+1; p < rp+size; ++p) {
                ++nr;
                if (*cp <= *p) {
                    --v;
                    break;
                }
            }
            s2 *= nr;
            score = (unsigned long)s1 * s2;

            if (v) ++vis;
            if (score > p2) p2 = score;
        }
    }
    cprintf("\r\n");
    cprintf("part 1: %d\r\n", (size-1)*4+vis);
    cprintf("part 2: %ld\r\n", p2);
}
