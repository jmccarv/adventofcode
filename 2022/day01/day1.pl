#!/usr/bin/env perl

use strict;
use warnings;
use List::Util qw( sum );

my $elf = 0;
my @elves;
while (<>) {
    chomp;
    if ($_) {
        $elves[$elf] += $_;
    } else {
        ++$elf;
    }
}

@elves = reverse sort { $a <=> $b } @elves;

print "part1: ".$elves[0]."\n";
print "part2: ". sum(@elves[0..2])."\n";
