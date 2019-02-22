#!/bin/bash

go-bindata \
  -o initializer/bindata.go \
  -pkg initializer \
  -prefix templates \
  templates templates/server templates/src