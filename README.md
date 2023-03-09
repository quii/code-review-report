# code-review-report ![build status](https://github.com/quii/code-review-report/workflows/Test/badge.svg)

Generates a report for Friday's code review 

- How successful we were at integrating our work
- Areas of the code we committed to most frequently to spark conversation on the changes going on
- Commits that failed to integrate
- Files that featured in failed integrations

## usage

Download the latest binary from [releases](https://github.com/quii/code-review-report/releases)

`./code-review-report -owner=quii -repo=code-review-report -repo=learn-go-with-tests`

### notes

- You'll need a `$GITHUB_TOKEN` env var if you need this to work with private repos
- Put the binary somewhere in your $PATH for convienience

## some other useful commands

`git log --pretty="%h [%cn] (%ar)  %s" --since="1 hour ago"`

```
git log --since 6.months.ago --numstat |
  awk '/^[0-9-]+/{ print $NF}' |
  sort |
  uniq -c |
  sort -nr |
  head
```
