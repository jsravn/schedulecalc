# schedulecalc

Simple calculator to calculate probability of a scheduling conflict, assuming event length is uniform.

Usage:
    
    go get github.com/jsravn/schedulecalc
    schedulecalc -events 6 -event-length 5m -window 24h

This gives the probability of a conflict for 6 events, with length of 5m, in a 24h window.
