#include <stdio.h>
#include <stdlib.h>

typedef struct {
    int len;
    int cap;
    int *lines;
} list_t;

void part1(list_t);
void part2(list_t);

void main() {
    int i;
    list_t input = { 0, 0, NULL };
    while(EOF != scanf("%d\n", &i)) {
        if (input.cap <= input.len) {
            input.cap += 100;
            input.lines = realloc(input.lines, sizeof(int) * input.cap);
        }
        input.lines[input.len] = i;
        input.len++;
    }

    part1(input);
    part2(input);
}

void part1(list_t input) {
    int last = 0;
    int nr = -1;
    for (int i = 0; i < input.len; i++) {
        nr  += input.lines[i] > last;
        last = input.lines[i];
    }
    printf("%d\n", nr);
}

void part2(list_t input) {
    int p1, p2;
    int s1, s2, nr;
    s1 = s2 = nr = 0;

    p1 = input.lines[0];
    p2 = input.lines[1];

    for (int i = 0; i < 3; i++) {
        s1 += input.lines[i];
        s2 += input.lines[i+1];
    }

    for (int i = 4; i < input.len; i++) {
        nr += s2 > s1;

        s1 -= p1;
        s1 += input.lines[i-1];
        p1 = p2;

        s2 -= p2;
        s2 += input.lines[i];
        p2 = input.lines[i-2];
    }
    nr += s2 > s1;

    printf("%d\n", nr);
}
