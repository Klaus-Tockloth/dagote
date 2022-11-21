#!/bin/sh

# ------------------------------------
# Purpose:
# - Builds assets (tar.gz or zip).
#
# Releases:
# - v1.0.0 - 2022-11-10: initial release
# ------------------------------------

# set -o xtrace
set -o verbose

# recreate directory
rm -r ./assets
mkdir ./assets

# asset 'aix'
tar -cvzf ./assets/aix-ppc64_dagote.tar.gz ./binaries/aix-ppc64/dagote

# assets 'darwin'
tar -cvzf ./assets/darwin-amd64_dagote.tar.gz ./binaries/darwin-amd64/dagote
tar -cvzf ./assets/darwin-arm64_dagote.tar.gz ./binaries/darwin-arm64/dagote

# assets 'dragonfly'
tar -cvzf ./assets/dragonfly-amd64_dagote.tar.gz ./binaries/dragonfly-amd64/dagote

# assets 'freebsd'
tar -cvzf ./assets/freebsd-amd64_dagote.tar.gz/freebsd-amd64/dagote
tar -cvzf ./assets/freebsd-arm64_dagote.tar.gz ./binaries/freebsd-arm64/dagote

# asset 'illumos'
tar -cvzf ./assets/illumos-amd64_dagote.tar.gz ./binaries/illumos-amd64/dagote

# assets 'linux'
tar -cvzf ./assets/linux-amd64_dagote.tar.gz ./binaries/linux-amd64/dagote
tar -cvzf ./assets/linux-arm64_dagote.tar.gz ./binaries/linux-arm64/dagote
tar -cvzf ./assets/linux-mips64_dagote.tar.gz ./binaries/linux-mips64/dagote
tar -cvzf ./assets/linux-mips64le_dagote.tar.gz ./binaries/linux-mips64le/dagote
tar -cvzf ./assets/linux-ppc64_dagote.tar.gz ./binaries/linux-ppc64/dagote
tar -cvzf ./assets/linux-ppc64le_dagote.tar.gz ./binaries/linux-ppc64le/dagote
tar -cvzf ./assets/linux-riscv64_dagote.tar.gz ./binaries/linux-riscv64/dagote
tar -cvzf ./assets/linux-s390x_dagote.tar.gz ./binaries/linux-s390x/dagote

# assets 'netbsd'
tar -cvzf ./assets/netbsd-amd64_dagote.tar.gz ./binaries/netbsd-amd64/dagote
tar -cvzf ./assets/netbsd-arm64_dagote.tar.gz ./binaries/netbsd-arm64/dagote

# assets 'openbsd'
tar -cvzf ./assets/openbsd-amd64_dagote.tar.gz ./binaries/openbsd-amd64/dagote
tar -cvzf ./assets/openbsd-arm64_dagote.tar.gz ./binaries/openbsd-arm64/dagote
tar -cvzf ./assets/openbsd-mips64_dagote.tar.gz ./binaries/openbsd-mips64/dagote

# asset 'solaris'
tar -cvzf ./assets/solaris-amd64_dagote.tar.gz ./binaries/solaris-amd64/dagote

# assets 'windows'
zip ./assets/windows-amd64_dagote.zip ./binaries/windows-amd64/dagote.exe
zip ./assets/windows-386_dagote.zip ./binaries/windows-386/dagote.exe
zip ./assets/windows-arm_dagote.zip ./binaries/windows-arm/dagote.exe
