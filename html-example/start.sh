#!/bin/sh
# ------------------------------------
# Purpose: Start template execution.
# ------------------------------------

set -o verbose
dagote -templates='category.tmpl' -output=category.html -format=html -dotfile='category-67.json' -dottype='json'
