#!/usr/bin/env bash

set -ex

: ${DOWNSTREAM=git@github.com:tonistiigi/gvisor.git}
: ${UPSTREAM=git://github.com/google/gvisor.git}


tmp=$(mktemp -u)

git clone $DOWNSTREAM $tmp
cd $tmp

git fetch origin upstream:prev filter-state:filter-state
git fetch $UPSTREAM master:upstream
git checkout -B master upstream

if [ "$(git rev-parse upstream)" = "$(git rev-parse prev)" ]; then
  exit 0;
fi

git filter-branch -f --tree-filter unbazel-gvisor --msg-filter "sed '\${/^\$/d;p;}' | sed '\$d'; echo Upstream-commit: \$GIT_COMMIT" --state-branch refs/heads/filter-state prev..master

git push origin master:master upstream:upstream filter-state:filter-state
