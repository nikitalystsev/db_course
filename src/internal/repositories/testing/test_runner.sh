#!/bin/bash

docker run --rm --network="host" pgtap-tester:latest -h 0.0.0.0 -p 5432 -u postgres -w postgres -d smartshopper