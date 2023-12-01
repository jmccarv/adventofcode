#include <cbm.h>
#include <conio.h>
#include <string.h>

void box(unsigned char lx, unsigned char ty, unsigned char rx, unsigned char by, char *title) {
    unsigned char w = rx-lx-1;
    unsigned char h = by-ty-1;
    unsigned char c = 0;
    unsigned char tofs = 0;
    unsigned char tlen, wm2;

    cputcxy(lx, ty, CH_ULCORNER);
    chline(w);
    cputc(CH_URCORNER);

    if (title && w > 2) {
        wm2 = w-2;
        tlen = strlen(title);
        if (tlen > wm2) {
            c = title[wm2];
            title[wm2] = '\0';
        } else {
            tofs = ((wm2)-tlen)/2;
        }

        cputcxy(lx+1+tofs, ty, CH_RTEE);
        cputs(title);
        cputc(CH_LTEE);

        if (c) {
            title[wm2] = c;
        }
    }

    cvlinexy(lx, ty+1, h);
    cvlinexy(rx, ty+1, h);

    cputcxy(lx, by, CH_LLCORNER);
    chline(w);
    cputc(CH_LRCORNER);
}
