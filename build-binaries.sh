#!/bin/sh

# ------------------------------------
# Purpose:
# - Builds executables / binaries.
#
# Releases:
# - v1.0.0 - 2022-11-18: initial release
#
# Remarks:
# - go tool dist list
# ------------------------------------

# set -o xtrace
set -o verbose

# compile 'aix'
env GOOS=aix GOARCH=ppc64 go build -o binaries/aix-ppc64/dagote

# compile 'darwin'
env GOOS=darwin GOARCH=amd64 go build -o binaries/darwin-amd64/dagote
env GOOS=darwin GOARCH=arm64 go build -o binaries/darwin-arm64/dagote

# compile 'dragonfly'
env GOOS=dragonfly GOARCH=amd64 go build -o binaries/dragonfly-amd64/dagote

# compile 'freebsd'
env GOOS=freebsd GOARCH=amd64 go build -o binaries/freebsd-amd64/dagote
env GOOS=freebsd GOARCH=arm64 go build -o binaries/freebsd-arm64/dagote

# compile 'illumos'
env GOOS=illumos GOARCH=amd64 go build -o binaries/illumos-amd64/dagote

# compile 'linux'
env GOOS=linux GOARCH=amd64 go build -o binaries/linux-amd64/dagote
env GOOS=linux GOARCH=arm64 go build -o binaries/linux-arm64/dagote
env GOOS=linux GOARCH=mips64 go build -o binaries/linux-mips64/dagote
env GOOS=linux GOARCH=mips64le go build -o binaries/linux-mips64le/dagote
env GOOS=linux GOARCH=ppc64 go build -o binaries/linux-ppc64/dagote
env GOOS=linux GOARCH=ppc64le go build -o binaries/linux-ppc64le/dagote
env GOOS=linux GOARCH=riscv64 go build -o binaries/linux-riscv64/dagote
env GOOS=linux GOARCH=s390x go build -o binaries/linux-s390x/dagote

# compile 'netbsd'
env GOOS=netbsd GOARCH=amd64 go build -o binaries/netbsd-amd64/dagote
env GOOS=netbsd GOARCH=arm64 go build -o binaries/netbsd-arm64/dagote

# compile 'openbsd'
env GOOS=openbsd GOARCH=amd64 go build -o binaries/openbsd-amd64/dagote
env GOOS=openbsd GOARCH=arm64 go build -o binaries/openbsd-arm64/dagote
env GOOS=openbsd GOARCH=mips64 go build -o binaries/openbsd-mips64/dagote

# compile 'solaris'
env GOOS=solaris GOARCH=amd64 go build -o binaries/solaris-amd64/dagote

# compile 'windows'
env GOOS=windows GOARCH=amd64 go build -o binaries/windows-amd64/dagote.exe
env GOOS=windows GOARCH=386 go build -o binaries/windows-386/dagote.exe
env GOOS=windows GOARCH=arm go build -o binaries/windows-arm/dagote.exe
