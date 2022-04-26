#!/bin/bash

nohup ./mobileAddress > run.log 2>&1 &

tail -f run.log
