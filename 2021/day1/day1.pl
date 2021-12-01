#!/usr/bin/env perl

use strict;
use warnings;
use List::Util qw(reduce);

my @m;  # measurements
while (<>) {
    chomp;
    push @m, $_;
}

part1(@m);
part2(@m);

sub part1 {
    my $last = 0;
    my $nr = -1;
    for (@_) {
        $nr++ if $_ > $last;
        $last = $_;
    }
    print "$nr\n";
}

sub part2 {
    my $nr = 0;
    while (@_ >= 4) {
        my $s1 = reduce { $a + $b } @_[0..2];
        shift;
        my $s2 = reduce { $a + $b } @_[0..2];
        $nr++ if $s2 > $s1;
    }
    print "$nr\n";
}
