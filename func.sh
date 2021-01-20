#!/bin/bash

go build -o cli .

echo "  TEST QUERY: ./cli host port abracadabra test"
./cli host port abracadabra test
echo
echo "  TEST QUERY: ./cli host port abracadabra xxx"
./cli host port abracadabra xxx
echo
echo "  TEST QUERY: ./cli host port some thing"
./cli host port some thing
echo
echo "  TEST QUERY: ./cli"
./cli
echo "END OF FUNC TEST"

rm -rf cli
