#! /bin/bash

atlas migrate apply \
  --dir "file://migrations" \
  --url $DATABASE_URL