#!/usr/bin/env perl

use strict;
use warnings;
use v5.10;

my $p1;
my $p2;
while (<>) {
    my @seq = /([-\d]+)/g;
    $p1 += get_next_val(@seq);
    $p2 += get_first_val(@seq);
}
say "Part1 $p1";
say "Part2 $p2";

sub get_next_val {
    my @seq = @_;
    return 0 if @seq < 2;

    my @new;
    my $a = shift @seq;
    while (@seq) {
        push @new, $seq[0] - $a;
        $a = shift @seq;
    }
    return $a + get_next_val(@new);
}

sub get_first_val {
    my @seq = @_;
    return 0 if @seq < 2;

    my @new;
    my $f = $a = shift @seq;
    while (@seq) {
        push @new, $seq[0] - $a;
        $a = shift @seq;
    }
    return $f - get_first_val(@new);
}
