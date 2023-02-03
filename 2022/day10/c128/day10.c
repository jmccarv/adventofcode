#include <conio.h>
#include <time.h>
#include <stdio.h>
#include <cbm.h>

#include "util/input.h"
#include "util/tui.h"

// used by input.c
const char *infile_prefix = "d10.";

void part1(void);
void crt(void);
void tick(void);
void init(void);

struct machine {
    int clock;
    int x;  // X register
    int p1;
};

struct machine m = { 0, 1, 0 };

void main(void) {
    clock_t t0, t1;
    unsigned char x,y;
    char *data;
    register char *p;
    int arg;

    screensize(&x, &y);
    if (x == 80) fast();

    if (NULL == (data = load_input(get_input_file()))) return;

    t0 = clock();

    init();
    for (p = data; *p; ++p) {
        tick();
        //cprintf("C: %d  X: %d  op: %c  ", m.clock, m.x, *p);
        crt();

        // CPU
        part1();
        switch (*p) {
            case 'a': //addx
                tick();
                crt();
                p += 5;  // move past 'addx ' to get to the argument
                sscanf(p, "%d\n", &arg);

                // part 1
                //cprintf("%d",m.clock);
                part1();
                
                m.x += arg;
                break;
        }
        for (; *p && *p != '\n'; ++p);
        //cprintf("\r\n");
    }
    t1 = clock();

    gotoxy(0, 12);
    textcolor(COLOR_CYAN);
    cprintf("part1: %d\r\n", m.p1);

    cprintf("Runtime: %ld\n", t1-t0);
}

void init(void) {
    char i = 0;
    bgcolor(COLOR_BLACK);
    bordercolor(COLOR_BLACK);
    textcolor(COLOR_RED);

    clrscr();
    gotoxy(0,0);
    box(0, 0, 20, 2);
    textcolor(COLOR_LIGHTRED);
    cputsxy(2, 1, "CLK");
    cputsxy(12, 1, "X");

    textcolor(COLOR_LIGHTGREEN);
}

void tick(void) {
    ++m.clock;
    gotoxy(6, 1);
    cprintf("%03d", m.clock);
    gotoxy(14,1);
    cprintf("%03d", m.x);
}

void part1(void) {
    if (m.clock == 20 || (m.clock-20)%40 == 0) {
        m.p1 += m.x * m.clock;
    }
}

void crt(void) {
    unsigned char y = (m.clock - 1) / 40;
    unsigned char x = (m.clock - 1) % 40;

    // X register is the midpoint of our 3px wide cursor
    if (x >= m.x-1 && x <= m.x+1) {
        // The current pixel is lit
        //textcolor(COLOR_LIGHTGREEN);
        cputcxy(x, y+4, 230);
    }
}
