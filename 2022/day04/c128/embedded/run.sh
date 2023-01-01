#!/bin/bash

./gen_input.pl > d4b_data.h

cl65 -t c128 -o day04 d4b.c && x128 day04
