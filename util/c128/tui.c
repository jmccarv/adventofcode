#include <cbm.h>
#include <conio.h>

void box(unsigned char lx, unsigned char ty, unsigned char rx, unsigned char by) {
    unsigned char i;
    cputcxy(lx, ty, CH_ULCORNER);
    for (i = lx+1; i < rx; ++i) {
        cputc(CH_HLINE);
    }
    cputc(CH_URCORNER);

    for (i = ty+1; i < by; ++i) {
        cputcxy(lx, i, CH_VLINE);
        cputcxy(rx, i, CH_VLINE);
    }

    cputcxy(lx, by, CH_LLCORNER);
    for (i = lx+1; i < rx; ++i) {
        cputc(CH_HLINE);
    }
    cputc(CH_LRCORNER);
}
