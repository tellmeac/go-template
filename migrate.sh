#! /bin/bash

# TODO: Makefile would be nice

atlas migrate apply \
  --dir "file://migrations" \
  --url $DATABASE_URL