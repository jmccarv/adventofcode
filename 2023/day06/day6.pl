#!/usr/bin/env perl

use strict;
use warnings;
use v5.10;

my @times = <> =~ /(\d+)/g;
my @dists = <> =~ /(\d+)/g;

part1(\@times, \@dists);
part2(\@times, \@dists);

sub part1 {
    my @times = @{$_[0]};
    my @dists = @{$_[1]};
    my $p1 = 1;
    while (@times) {
        $p1 *= p1calc(shift @times, shift @dists);
    }

    say "Part1 ",$p1;
}

sub p1calc {
    my ($time, $dist) = @_;

    my $nr = 0;
    for (my $t = 1; $t < $time; $t++) {
        my $move = $t * ($time - $t);
        ++$nr if $move > $dist;
        return $nr if $nr && $move <= $dist;
    }
    return $nr;
}


sub calc {
    my ($t, $time) = @_;
    return $t*($time-$t);
}

sub part2 {
    use integer;
    my $time = join('',@{$_[0]});
    my $dist = join('',@{$_[1]});

    my $m, my $l = 0;
    my $r = $time/2;
    while ($l <= $r) {
        $m = $l + ($r - $l)/2;
        if (calc($m, $time) > $dist) {
            if (calc($m-1, $time) <= $dist) {
                last
            }
            $r = $m-1;
        } else {
            $l = $m+1;
        }
    }
    my $left = $m;

    $l = $time / 2+1;
    $r = $time;
    while ($l <= $r) {
        $m = $l + ($r - $l)/2;
        if (calc($m, $time) > $dist) {
            if (calc($m+1, $time) <= $dist) {
                last
            }
            $l = $m+1;
        } else {
            $r = $m-1;
        }
    }
    my $right = $m;

    say "Part2 $left - $right = ",$right-$left+1;
}
