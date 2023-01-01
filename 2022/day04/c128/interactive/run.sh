#!/bin/bash

cl65 -t c128 -o day04 day04.c && {
    petcat -text -w2 -o example.pet -- example_input
    petcat -text -w2 -o input.pet -- input

    c1541 -format mjmtest,jm d71 test.d71 \
        -write day04 day04 \
        -write example.pet example,s \
        -write input.pet input,s

    x128 test.d71
}
