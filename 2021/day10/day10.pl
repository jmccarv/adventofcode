#!/usr/bin/env perl

use strict;
use warnings;

my %closer_for = qw# [ ] ( ) { } < > #;
my %p1_points  = qw# ) 3 ] 57 } 1197 > 25137 #;
my %p2_points  = qw# ) 1 ] 2 } 3 > 4 #;

my (@stack, @incomplete, $p1, $good, $p2);
while (<>) {
    chomp; $good = 1; my @stack = ();
    for my $tok (split //) {
        if (my $c = $closer_for{$tok}) {
            push @stack, $c;
        } elsif ($tok ne pop @stack) {
            $p1 += $p1_points{$tok};
            $good = 0;
            last;
        }
    }
    push @incomplete, [@stack] if $good;
}
print "$p1\n";

my @scores = sort {$a <=> $b} 
             map{ my $score=0; map { $score = $score*5 + $p2_points{$_} } 
             reverse @$_; $score } @incomplete;
print $scores[@scores/2]."\n";
