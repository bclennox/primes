#!/bin/bash

make clean primes

if diff verified.txt <(./build/primes 1000000 | sort -n) ; then
  echo "Passed"
else
  echo
  echo "Failed"
fi
