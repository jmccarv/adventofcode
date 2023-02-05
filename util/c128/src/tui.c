#include <cbm.h>
#include <conio.h>

void box(unsigned char lx, unsigned char ty, unsigned char rx, unsigned char by) {
    unsigned char w = rx-lx-1;
    unsigned char h = by-ty-1;

    cputcxy(lx, ty, CH_ULCORNER);
    chline(w);
    cputc(CH_URCORNER);

    cvlinexy(lx, ty+1, h);
    cvlinexy(rx, ty+1, h);

    cputcxy(lx, by, CH_LLCORNER);
    chline(w);
    cputc(CH_LRCORNER);
}
