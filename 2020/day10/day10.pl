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

#
# The key here is to break the input down in to multiple small(er)
# graphs instead of one giant graph. This makes the solution run
# much MUCH faster.
# 
# We only need a graph when there is more than one path possible.
# Consider this part of the example_input (0 added to front):
#   0 1 4 5 6 7 10 11 12
#
# This can be broken into 2 graphs:
#   1-4 4-5 4-6 4-7 5-6 5-7 6-7
#   10-11 10-12 11-12
# 
# This is possible because, for instance, there is only one
# path to get from 7 to 10, so we can split the graph there.
#
# Then we can compute the number of possible paths by taking
# the product of possible paths through each graph having more than one path
#
sub part2_redux {
    my @input = @_;
    my @graphs;

    my @prev;
    my $g = Graph->new;
    for my $n (@input) {
        # Remove any previous elements that are not reachable from $n
        shift @prev while @prev && $prev[0] < $n-3;

        if (!@prev || $prev[$#prev]+3 == $n) {  # safe to split here
            # Save the current graph (if we have one) and start a new one
            push @graphs, $g if $g->edges > 1;
            $g = Graph->new;
            @prev = ()
        }

        $g->add_edge($_, $n) for @prev;
        push @prev, $n;
    }
    push @graphs, $g if $g->edges > 1;

    # Sanity check
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

# This works, but so, so slow (and consumes much memory for large graphs)
# Basically unusable for the real input set
# Left here for posterity. Don't try this at home.
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
