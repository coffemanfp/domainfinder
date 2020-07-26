#!/bin/bash

echo Building domainfinder...
go build -o domainfinder

echo Building synonyms...
cd ../synonyms
go build -o ../domainfinder/lib/synonyms

echo Building available...
cd ../available
go build -o ../domainfinder/lib/available

echo Building sprinkle...
cd ../sprinkle
go build -o ../domainfinder/lib/sprinkle

echo Building coolify...
cd ../coolify
go build -o ../domainfinder/lib/coolify

echo Building domainify...
cd ../domainify
go build -o ../domainfinder/lib/domainify
