#include <stdlib.h>
#include <stdio.h>
#include <string.h>

typedef struct _state_t {
    long regs[4];
    long min;
    long max;
} state_t;

typedef struct _state_list_t {
    long nr;
    state_t *s;
} state_list_t;

static inline void *xmalloc(size_t len) {
    void *p = malloc(len);
    if (!p) {
        fprintf(stderr, "malloc failed\n");
        exit(1);
    }
    return p;
}

static inline void *xzmalloc(size_t len) {
    void *p = xmalloc(len);
    memset(p, 0, len);
    return p;
}

static inline void *xrealloc(void *p, size_t len) {
    p = realloc(p, len);
    if (!p) {
        fprintf(stderr, "realloc failed\n");
        exit(1);
    }
    return p;
}

void add_state(state_list_t *l, state_t s) {
    l->nr++;
    l->s = xrealloc(l->s, sizeof(state_t) * l->nr);
    l->s[l->nr-1] = s;
    return;
}

static inline int state_cmp(const void *p1, const void *p2) {
    const state_t *a = p1;
    const state_t *b = p2;
    for (int i = 0; i < 4; i++) {
        if (a->regs[i] < b->regs[i]) {
            return -1;
        } else if (a->regs[i] > b->regs[i]) {
            return 1;
        }
    }
    return 0;
}

void free_state_list(state_list_t *l) {
    free(l->s);
    l->s = NULL;
    l->nr = 0;
}

static inline long argval(long regs[4], char *arg) {
    switch (arg[0]) {
        case 'w': return regs[0]; break;
        case 'x': return regs[1]; break;
        case 'y': return regs[2]; break;
        case 'z': return regs[3]; break;
    }
    return atol(arg);
}

void dump_state(state_t *s) {
    printf("w:%ld x:%ld y:%ld z:%ld min:%ld max:%ld\n", 
           s->regs[0], s->regs[1], s->regs[2], s->regs[3], s->min, s->max);
}

void dump_states(state_list_t l) {
    printf("Dumping %ld\n", l.nr);
    for (state_t *s = l.s; s < l.s+l.nr; s++) {
        dump_state(s);
    }
}

#define MIN(a,b) ((a) < (b) ? (a) : (b))
#define MAX(a,b) ((a) > (b) ? (a) : (b))

int main(void) {
    char line[20];
    char op[4];
    char reg[2];
    char arg[10];
    int i, nr, nrinp = 0;

    setlinebuf(stdout);

    state_list_t states = {0, NULL};
    long new_regs[4] = {0,0,0,0};
    state_t empty_state = {{0,0,0,0},0,0};
    add_state(&states, empty_state);

    dump_states(states);

    while (fgets(line, 20, stdin)) {
        op[0] = reg[0] = arg[0] = '\0';
        if ((nr = sscanf(line, "%s %s %s", op, reg, arg)) < 2) {
            continue;
        }

        int ra;
        switch (reg[0]) {
            case 'w': ra = 0; break;
            case 'x': ra = 1; break;
            case 'y': ra = 2; break;
            case 'z': ra = 3; break;
            default:
                fprintf(stderr, "Invalid register %s\n", reg);
                exit(1);
        }

        if (!strcmp(op, "inp")) {
            printf("inp %d - %ld\n", ++nrinp, states.nr);

            for (int i = 0; i < states.nr; i++) {
                states.s[i].regs[ra] = 0;
            } 

            printf("  sorting...\n");
            qsort(states.s, states.nr, sizeof(state_t), state_cmp);
            printf("  sort done, compacting...\n");
            i = 0;
            for (int j = 1; j < states.nr; j++) {
                if (0 == state_cmp(&states.s[i], &states.s[j])) {
                    states.s[i].min = MIN(states.s[i].min, states.s[j].min);
                    states.s[i].max = MAX(states.s[i].max, states.s[j].max);
                    continue;
                }
                states.s[++i] = states.s[j];
            }
            states.nr = i+1;

            printf("  %ld after compaction\n", states.nr);
            state_list_t new_states;
            new_states.nr = states.nr * 9;
            new_states.s = xmalloc(sizeof(state_t) * new_states.nr);
            for (i = 0; i < states.nr; i++) {
                for (int j = 1; j <= 9; j++) {
                    state_t *p = &new_states.s[i*9+j-1];
                    *p = states.s[i];
                    p->regs[ra] = j;
                    p->max = p->max*10+j;
                    p->min = p->min*10+j;
                }
            }
            free_state_list(&states);
            states = new_states;
        } else if (!strcmp(op, "add")) {
            for (i=0; i < states.nr; i++) {
                states.s[i].regs[ra] += argval(states.s[i].regs, arg);
            }
        } else if (!strcmp(op, "mul")) {
            for (i=0; i < states.nr; i++) {
                states.s[i].regs[ra] *= argval(states.s[i].regs, arg);
            }
        } else if (!strcmp(op, "div")) {
            for (i=0; i < states.nr; i++) {
                states.s[i].regs[ra] /= argval(states.s[i].regs, arg);
            }
        } else if (!strcmp(op, "mod")) {
            for (i=0; i < states.nr; i++) {
                states.s[i].regs[ra] %= argval(states.s[i].regs, arg);
            }
        } else if (!strcmp(op, "eql")) {
            for (i=0; i < states.nr; i++) {
                states.s[i].regs[ra] = states.s[i].regs[ra] == argval(states.s[i].regs, arg) ? 1 : 0;
            }
        }
    }

    long p1 = 0;
    long p2 = 99999999999999;
    for (state_t *s = states.s; s < states.s+states.nr; s++) {
        if (s->regs[3] != 0) continue;
        dump_state(s);
        p1 = MAX(p1, s->max);
        p2 = MIN(p2, s->min);
    }
    printf("p1: %ld\n", p1);
    printf("p2: %ld\n", p2);
    return 0;
}

