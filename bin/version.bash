#!/bin/bash

# shellcheck shell=bash

main() {
  local latest_tag=$(git tag --list "v*" --sort=-v:refname | head -1)

  local major=$(echo $latest_tag | cut -d. -f1)
  local minor=$(echo $latest_tag | cut -d. -f2)
  local patch=$(echo $latest_tag | cut -d. -f3)

  local log_lines=$(git log $latest_tag..HEAD --pretty="format:- %s")

  if echo "$log_lines" | grep -q "BREAKING CHANGE"; then
    major=$((major + 1))
    minor=0
    patch=0
  elif echo "$log_lines" | grep -q "major"; then
    major=$((major + 1))
    minor=0
    patch=0
  elif echo "$log_lines" | grep -q "feat"; then
    minor=$((minor + 1))
    patch=0
  elif echo "$log_lines" | grep -q "fix"; then
    patch=$((patch + 1))
  elif echo "$log_lines" | grep -q "chore"; then
    patch=$((patch + 1))
  else
    true
    # ci
    # docs
    # major
    # perf
    # refactor
    # revert
    # style
    # test
  fi

  local new_tag="$major.$minor.$patch"
  npm version $new_tag --no-git-tag-version

  mkdir -p dist/
  echo $new_tag >dist/releasetag.txt
  echo "$log_lines" | grep -v "chore(release):" >dist/changelog.md
  echo "Semantic version tag: $latest_tag -> $new_tag"
  cat dist/changelog.md

  return 0
}

main "$@"
