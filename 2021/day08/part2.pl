#!/usr/bin/env perl

use strict;
use warnings;

use Data::Dumper qw(Dumper);
use Array::Set qw(set_diff);

# length => digit
my %known = (2=>1, 4=>4, 3=>7, 7=>8);

my $ret;
while (<>) {
    chomp;
    my ($segments, $digits) = split /\|/;
    $ret += process([ split /\s/, $segments ], [ grep { length($_) } split /\s/, $digits ]);
}
print "$ret\n";

sub unknown { grep { !defined $_->{value} } @_ }
sub process {
    my ($segments, $digits) = @_;
    my (%k, $num);

    # Turn them into objects! yay!
    my @p = map { segment($_) } @$segments;

    # Let's start by finding our known digits, these are the ones that
    # use a unique number of segments in the 7 segment display: 1, 4, 7, and 8
    # This will be a map of { num => segment_object } i.e. { 7 => $obj }
    %k = map { $_->value => $_ } grep { $_->value } @p;

    # We can figure out which segment is 6: it has 6 segments and we can subtract
    # our known 7's segments from the possible 6s to find the one that must be 6
    # 7 - 6 => 4 segments
    $k{6} = (grep { $_->minus($k{7}) == 4 } grep { $_->len == 6 } unknown(@p))[0]->setval(6);

    # From 6 we can find 5, 5 has all the same segments as 6, minus one, so 6 - 5 => 1 segment
    $k{5} = (grep { $k{6}->minus($_) == 1 } grep { $_->len == 5 } unknown(@p))[0]->setval(5);
   
    # 3: 3 - 5 => 1
    $k{3} = (grep { $_->minus($k{5}) == 1 } grep { $_->len == 5 } unknown(@p))[0]->setval(3);

    # 2: 2 - 5 => 2
    $k{2} = (grep { $_->minus($k{5}) == 2 } grep { $_->len == 5 } unknown(@p))[0]->setval(2);

    # 9: 5 - 9 => 0
    $k{9} = (grep { $k{5}->minus($_) == 0 } grep { $_->len == 6 } unknown(@p))[0]->setval(9);

    # The last unknown is 0
    (unknown(@p))[0]->setval(0);

    # Now we can decode the digits
    for my $d (@$digits) {
        $num .= (grep { $_->equal($d) } @p)[0]->value;
    }
    $num;
}

sub segment {
    my @letters = sort split //, shift;
    my $segment = join('', @letters);
    bless {
        segment => $segment,
        letters => \@letters,
        len     => length($segment),
        value   => $known{length($segment)},
    };
}
sub letters  { @{shift->{letters}} }
sub value    { shift->{value} }
sub len      { shift->{len} }
sub minus    { my @r = sort @{set_diff(shift->{letters}, shift->{letters})}; @r }
sub equal    { shift->{segment} eq join('', sort split //, shift) }
sub setval   { my $self = shift; $self->{value} = shift; $self }
