#!/usr/bin/env perl

use strict;
use warnings;
#use Data::Dumper;

my $bags = {};

while(<>) {
    next unless /^(.*)\sbags\scontain\s/;
    my $b = { color => $1, contains => {} };

    for (split /,/, $_) {
        next unless /(\d+)\s(.*?)\sbags?/;
        $b->{contains}->{$2} = $1;
    }
    $bags->{$b->{color}} = $b;
}

#print Dumper($bags);
print "part1 ", part1("shiny gold"), "\n";
print "part2 ", part2("shiny gold"), "\n";

sub part1 {
    my $color = shift;
    my $nr = 0;
    map { $nr++ if find1($color, $_) } values %$bags;
    $nr;
}

sub find1 {
    my ($color, $bag) = @_;
    grep { $color eq $_ || find1($color, $bags->{$_}) } keys %{$bag->{contains}}
}

sub part2 {
    my $color = shift;
    my $ret = 0;
    for (keys %{$bags->{$color}->{contains}}) {
        my $nr = $bags->{$color}->{contains}->{$_};
        $ret += $nr + $nr*part2($_)
    }
    $ret;
}
