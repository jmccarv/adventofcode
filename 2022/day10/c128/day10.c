#include <conio.h>
#include <time.h>
#include <stdio.h>
#include <cbm.h>

#include "input.h"
#include "tui.h"

#define INFILE_PREFIX "d10."

void part1(void);
void crt(void);
void cpu(register char *p);
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

    screensize(&x, &y);
    if (x == 80) fast();

    if (NULL == (data = load_input(get_input_file(INFILE_PREFIX)))) return;

    t0 = clock();
    init();
    for (p = data; *p; ++p) {
        tick();
        crt();
        cpu(p);
        for (p += 4; *p && *p != '\n'; ++p);
    }
    t1 = clock();

    gotoxy(0, 12);
    textcolor(COLOR_CYAN);
    cprintf("  Part1: %d\r\n", m.p1);
    cprintf("Runtime: %ld\n", t1-t0);
}

void cpu(register char *p) {
    int arg;

    part1();
    if (*p != 'a') return; // noop

    // addx takes two cycles to complete
    tick();
    crt();

    // move past 'addx ' to get to the argument
    sscanf(p+5, "%d\n", &arg);

    part1();
    m.x += arg;

    return;
}

void init(void) {
    unsigned char i = 0;
    bgcolor(COLOR_BLACK);
    bordercolor(COLOR_BLACK);
    textcolor(COLOR_RED);

    clrscr();
    gotoxy(0,0);
    box(10, 0, 30, 3);
    textcolor(COLOR_LIGHTRED);
    cputsxy(12, 1, "CLK");
    cputsxy(22, 1, "X");

    textcolor(COLOR_ORANGE);
    cputsxy(14, 2, "S");
    cputsxy(22, 2, "T");

    textcolor(COLOR_LIGHTGREEN);
}

void tick(void) {
    ++m.clock;
    gotoxy(16, 1);
    cprintf("%03d", m.clock);
    gotoxy(24,1);
    cprintf("%-3d", m.x);
}

void crt(void) {
    unsigned char y = (m.clock - 1) / 40;
    unsigned char x = (m.clock - 1) % 40;

    // X register is the midpoint of our 3px wide cursor
    if (x >= m.x-1 && x <= m.x+1) {
        // The current pixel is lit
        //textcolor(COLOR_LIGHTGREEN);
        cputcxy(x, y+5, 230);
    }
}

void part1(void) {
    int s;
    if (m.clock == 20 || (m.clock-20)%40 == 0) {
        s = m.x * m.clock;
        m.p1 += s;

        gotoxy(16, 2);
        cprintf("%4d", s); 

        gotoxy(24, 2);
        cprintf("%-5d", m.p1);
    }
}
