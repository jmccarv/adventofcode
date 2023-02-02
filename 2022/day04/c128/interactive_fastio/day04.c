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

void solve(char *);
struct cbm_dirent *get_input_file(void);

unsigned char dev;

const char infile_prefix[] = "d4.";

void main(void) {
    unsigned char x,y;
    clock_t t0, t1;
    struct cbm_dirent *ent;
    char *data;
    unsigned int nr;

    dev = getcurrentdevice();
    screensize(&x, &y);

    //cprintf("Start: DEV: %d  screen: %dx%d\r\n", dev, x,y);
    if (x == 80) fast();

    ent = get_input_file();

    t0 = clock();
    cprintf("Reading: %s\r\n", ent->name);

    // The size from the directory is the number of 256 byte blocks on disk
    // used for the file. Each block has two bytes of overhead, so 254 bytes
    // per block are used to store the file. The last block may not be full
    // but there is no way to know until you actually read the second byte
    // of the last block.
    // see: https://www.lemon64.com/forum/viewtopic.php?t=8485
    // and  https://retrocomputing.stackexchange.com/questions/17178/how-do-i-get-the-size-of-a-file-on-disk-on-the-commodore-64
    if (NULL == (data = malloc(254*ent->size))) {
        // Since the first two bytes of the file are skipped, but we want to add
        // a trailing null to this data, we'll do size-1 (instead of -2) so we have room.
        cprintf("Failed to malloc %d bytes for input\r\n", 254*ent->size-1);
        return;
    }

    // This is much faster than reading the file a byte at a time but
    // it requires the file to have an extra two bytes at the beginning
    // that are ignored for our purposes -- they would be the load address
    // if we didn't pass a pointer to where we want the data loaded.
    nr = cbm_load(ent->name, dev, data);
    if (nr == 0) {
        cprintf("Failed to load input: %d", __oserror);
        return;
    }
    cprintf("Loaded %d bytes to %p\r\n", nr, data);
    data[nr] = '\0';

    solve(data);
    t1 = clock();

    cprintf("Solution runtime: %ld\r\n", t1-t0);
}

struct cbm_dirent *get_input_file(void) {
    unsigned char r = cbm_opendir(1, dev);
    register struct cbm_dirent *ent;
    unsigned char x;
    char c;
    unsigned char len = 0;
    unsigned char nr_ents = 0;
    static struct cbm_dirent ents[4];

    cprintf("Listing possible input files\r\n");
    cprintf("Nr %-16s %s\r\n", "File Name", "Size");
    for (x = 0; x < 24; x++) {
        cputc(CH_HLINE);
    }
    cprintf("\r\n");

    ent = &ents[0];
    if (r == 0) { 
        while (nr_ents < 4 && 0 == (r = cbm_readdir(1, ent))) {
            if (ent->type == CBM_T_PRG && strlen(ent->name) > 3 && 0 == strncmp(ent->name, infile_prefix, strlen(infile_prefix))) {
                cprintf("%2d %-16s %4d\r\n", nr_ents+1, ent->name, ent->size);
                ent = &ents[++nr_ents];
            }
        }
        cbm_closedir(1);
    } else {
        cprintf("Opendir failed: %d\r\n", r);
    }

    // CH_DEL 
    cprintf("\r\nEnter number of file to read: ");
    while(1) {
        c = cgetc() - '1';
        if (c < nr_ents) {
            cprintf("\r\n");
            return &ents[c];
        }
    }
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
