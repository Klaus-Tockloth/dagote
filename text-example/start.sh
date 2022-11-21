#!/bin/sh
# ------------------------------------
# Purpose: Start template execution.
# ------------------------------------

set -o verbose
dagote -templates='test.tmpl' -output=test.txt
