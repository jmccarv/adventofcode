#!/usr/bin/env perl

use v5.10;
use List::Util qw(max all reduce);

my %p1max = ( red => 12, green => 13, blue => 14 );
while ($game = <>) {
    $id++;
    $m{$_} = max($game =~ /(\d+)\s$_/g) for keys %p1max;
    $p1 += $id if all { $m{$_} <= $p1max{$_} } keys %m;
    $p2 += reduce { $a * $b } values %m;
}
say "p1 $p1\np2 $p2"
