#!/usr/bin/env perl

use strict;
use warnings;
use Data::Dumper;

my @lines;

while (<>) {
    next unless /^\[/;
    push @lines, eval $_;
}

print "Part1: ".part1(@lines)."\n";
print "Part2: ".part2(@lines)."\n";

sub part1 {
    my $p = 0;
    my $sum = 0;
    while (@_) {
        $p++;
        $sum += $p if -1 == compare(shift, shift);
    }
    return $sum;
}

sub part2 {
    my @packets = sort { compare($a, $b) } @_, [[2]], [[6]];

    my $p1 = 0;
    for (my $i = 0; $i < @packets; $i++) {
        if (0 == compare([[2]], $packets[$i])) {
            $p1 = $i+1;
        
        } elsif (0 == compare([[6]], $packets[$i])) {
            return $p1 * ($i+1)
        }
    }
}

sub compare {
    my $l1 = [ @{$_[0]} ];
    my $l2 = [ @{$_[1]} ];

    while (@$l1 && @$l2) {
        my $l = shift @$l1;
        my $r = shift @$l2;

        my $ret;
        if (ref($1) && ref($2)) {
            # Both are lists
            $ret = compare($l, $r);

        } elsif (!ref($l) && !ref($r)) {
            # Both sides are integers
            $ret = $l <=> $r;

        } else {
            # One's a list, one's an integer
            $l = [$l] unless ref($l);
            $r = [$r] unless ref($r);

            $ret = compare($l, $r);
        }
        return $ret if $ret;
    }

    # Both lists are empty, continue checking
    return 0 unless @$l1 || @$l2;

    # left list ran out first, correct order
    return -1 unless @$l1;

    return 1 unless @$l2;
}
