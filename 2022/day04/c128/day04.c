#include <stdio.h>
#include <errno.h>
#include <string.h>
#include <conio.h>
#include <cbm.h>
#include <fcntl.h>
#include <peekpoke.h>
#include <device.h>
#include <time.h>

const int seq_type = 16;

void solve(char *);
void get_input_file(char *);

unsigned char dev;
char fn[17];

void main(void) {
    unsigned char x,y;
    clock_t t0, t1;

    dev = getcurrentdevice();
    screensize(&x, &y);

    //cprintf("Start: DEV: %d  screen: %dx%d\r\n", dev, x,y);
    if (x == 80) fast();

    get_input_file(fn);

    t0 = clock();
    solve(fn);
    t1 = clock();
    cprintf("Runtime: %ld\n", t1-t0);
}

void get_input_file(char *fn) {
    unsigned char r = cbm_opendir(1, dev);
    struct cbm_dirent ent;
    unsigned char c, x, y;
    unsigned char len = 0;

    cprintf("Listing possible input files\r\n");
    cprintf("%-16s %s\r\n", "File Name", "Size");
    for (x = 0; x < 21; x++) {
        cputc(CH_HLINE);
    }
    cprintf("\r\n");

    if (r == 0) { 
        while (0 == (r = cbm_readdir(1, &ent))) {
            if (ent.type == seq_type) {
                cprintf("%-16s %4d\r\n", ent.name, ent.size);
                strcpy(fn, ent.name);
            }
        }
        cbm_closedir(1);
    } else {
        cprintf("Opendir failed: %d\r\n", r);
    }

    // CH_DEL 
    cprintf("\r\nFile to read: ");

    y = wherey();
    x = wherex();
    fn[0] = '\0';
    cursor(1);
    for (c = cgetc(); c != '\n'; c = cgetc()) {
        if (c == CH_DEL) {
            if (len > 0) fn[--len] = '\0';

        } else if (len < 16) {
            fn[len++] = c;
            fn[len] = '\0';
        }
        cputsxy(x, y, fn);
        cputc(' ');
        gotox(x+len);
    }
    cprintf("\r\n");
}

struct assignment {
    unsigned int min;
    unsigned int max;
};

void solve(char *fn) {
    FILE *fh;
    char buf[81];
    unsigned char nr;
    struct assignment a, b;
    unsigned int nr_contain = 0;
    unsigned int nr_overlap = 0;

    cprintf("Attempting to read %s\r\n", fn);

    if (fh = fopen(fn, "r")) {
        while (fgets(buf, 80, fh)) {
            buf[strlen(buf)-1] = '\0';
            nr = sscanf(buf, "%d-%d,%d-%d", &a.min, &a.max, &b.min, &b.max);
            cprintf("%2d-%-2d %2d-%-2d : ", a.min,a.max, b.min,b.max);

            // part1 -- does one contain the other?
            if ( (a.min <= b.min && a.max >= b.max) || (b.min <= a.min && b.max >= a.max) ) {
                cputs("*");
                nr_contain++;
            }

            // part2 -- does one overlap the other?
            if ( (a.min <= b.max && a.max >= b.min) || (b.min <= a.max && b.max >= a.min) ) {
                nr_overlap++;
                cputs("+");
            }
            cprintf("\r\n");
        }
        cprintf("Part1: %d\r\n", nr_contain);
        cprintf("Part2: %d\r\n", nr_overlap);
        
        fclose(fh);
    } else {
        cprintf("Open failed: %d %d\r\n", errno, __oserror);
        perror("Open failed :(");
    }

}
