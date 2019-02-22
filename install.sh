#!/bin/bash

go-bindata \
  -o initializer/bindata.go \
  -pkg initializer \
  -prefix templates \
  templates \
  templates/public \
  templates/server \
  templates/server/api \
  templates/src \
  templates/src/actions \
  templates/src/components \
  templates/src/components/navigation \
  templates/src/constants \
  templates/src/modules \
  templates/src/reducers \
  templates/src/store \
  templates/src/views

go install