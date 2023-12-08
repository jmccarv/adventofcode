#!/usr/bin/env perl

use strict;
use warnings;
use Data::Dumper qw(Dumper);
use List::Util qw(max sum);
use v5.10;

my %sort_map = (A => 'z', K => 'y', Q => 'x', J => '/', T => 'v');
my @data;
while (<>) {
    my @d = split /\s/;
    push @data, { classify(hand => $d[0], bid => $d[1]) };
}
@data = sort hand_cmp @data;
# say Dumper(\@data);

my $rank = 1;
my $p2 = sum map { $_->{bid} * $rank++ } @data;
say "Part2 ", $p2;

sub hand_sort_key { join ('', map { $sort_map{$_} } split //, shift) }
sub sort_map { $sort_map{$_[0]} // $_[0] }

sub classify {
    my %d = @_;
    my %cards;

    $cards{$_}++ for split //, $d{hand};

    if ($cards{J}) {
        my $j = $cards{J};
        delete $cards{J};
        my @cards = reverse sort card_cmp map { { card => $_, count => $cards{$_} } } keys %cards;
        $cards[0]->{count} += $j;
        $cards{$cards[0]->{card} // 'J'} += $j;
    }

    my $maxdup = max values %cards;
    my $unique = values %cards;
    
    my $type = $maxdup;  # type: 7=5 of a kind, 1=high card
    $type+=2  if $maxdup > 3;                   # 5 or 4 of a kind
    $type++   if $maxdup == 3 && $unique == 2;  # full house
    $type++   if $maxdup == 3;                  # 3 of a kind
    $type++   if $maxdup == 2 && $unique == 3;  # two pair
    $d{type} = $type;
    $d{sort_key} = join('', map { sort_map($_) } split //, $d{hand});

    return %d;
}

sub card_cmp {
    $a->{count} <=> $b->{count} ||
    sort_map($a->{card}) cmp sort_map($b->{card});
}

sub hand_cmp {
    return  1 if $a->{type} > $b->{type};
    return -1 if $a->{type} < $b->{type};
    $a->{sort_key} cmp $b->{sort_key};
}
