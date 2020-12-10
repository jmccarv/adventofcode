#!/usr/bin/env perl

use strict;
use warnings;
use Graph;
use Data::Dumper;

exit main();

sub main {
    my @input = sort { $a <=> $b } map{ chomp; $_ } <>;

    unshift @input, 0;
    push @input, $input[$#input]+3;

    my %dist;
    for (1..$#input) {
        $dist{abs($input[$_] - $input[$_-1])}++;
    }
    print Dumper(\%dist);
    print "part1: ".($dist{1} * $dist{3})."\n";
    print "part2: ".part2_redux(@input)."\n";
}

sub part2_redux {
    my @input = @_;
    my @graphs;

    my @prev;
    my $g;
    for my $n (@input) {
        shift @prev while @prev && $prev[0] < $n-3;
        if ((@prev && $prev[$#prev] + 3 == $n) || ($g && !@prev)) {
            push @graphs, $g if $g->edges;
            $g = undef;
        }

        $g = Graph->new unless $g;
        $g->add_edge($_, $n) for @prev;

        push @prev, $n;
    }
    push @graphs, $g if $g->edges;

    $_->expect_dag for @graphs;

    print "nr graphs: ".@graphs."\n";
    print "$_\n" for @graphs;
    print "\n";

    my $ret = 1;
    for $g (@graphs) {
        my $a = $g->APSP_Floyd_Warshall();
        my $start = ($g->source_vertices)[0];
        my $end = ($g->sink_vertices)[0];
        $ret *= scalar $a->all_paths($start, $end)
    }
    $ret;
}

# This works, but so, so slow
# Basically unusable for the real input set
sub part2 {
    my @input = @_;
    my $g = Graph->new;

    my @prev;
    for my $n (@input) {
        shift @prev while @prev && $prev[0] < $n-3;
        $g->add_edge($_, $n) for @prev;
        push @prev, $n;
    }

    $g->expect_dag;

    my $a = $g->APSP_Floyd_Warshall();
    return scalar $a->all_paths($input[0], $input[$#input]);
}
