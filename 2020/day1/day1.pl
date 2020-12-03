#!/usr/bin/env perl

# https://adventofcode.com/2020/day/1

use strict;
use warnings;

my @nums;
while (<>) {
    chomp;
    push @nums, $_;
}

# Sorting allows us to take a shortcut later
@nums = sort { $a <=> $b } @nums;

my ($a, $b, $c);

($a, $b) = find(@nums);
print "$a * $b = ".$a*$b."\n" if $a && $b;

($a, $b, $c) = find3(@nums);
print "$a * $b * $c = ".$a*$b*$c."\n" if $a && $b && $c;

# Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.
sub find {
    my @nums = @_;
    my $a;
    my $b;

    while (@nums > 1) {
        $a = pop @nums;
        for $b (@nums) {
            return ($a, $b) if $a+$b == 2020;
            last if $a+$b > 2020;
        }
    }

    return undef,undef
}

# Find 3 entries that sum to 2020
# ugly brute force ... meh whatever it's q&d and works
sub find3 {
    my @nums = @_;
    my ($a, $b, $c);

    while (@nums > 2) {
        $a = pop @nums;

        my @nums2 = @nums;
        while (@nums2 > 1) {
            $b = pop @nums2;
            for $c (@nums2) {
                return ($a, $b, $c) if $a+$b+$c == 2020;
                last if $a+$b+$c > 2020;
            }
        }
    }

    return undef,undef,undef
}
