#!/bin/bash

version=v$(< VERSION)
git add *; git add .*
git commit -S -m $version && git tag -a $version -m $version
