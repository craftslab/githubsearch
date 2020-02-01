# githubsearch

[![Build Status](https://travis-ci.com/craftslab/githubsearch.svg?branch=master)](https://travis-ci.com/craftslab/githubsearch)
[![Coverage Status](https://coveralls.io/repos/github/craftslab/githubsearch/badge.svg?branch=master)](https://coveralls.io/github/craftslab/githubsearch?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/craftslab/githubsearch)](https://goreportcard.com/report/github.com/craftslab/githubsearch)
[![License](https://img.shields.io/github/license/craftslab/githubsearch.svg?color=brightgreen)](https://github.com/craftslab/githubsearch/blob/master/LICENSE)
[![Tag](https://img.shields.io/github/tag/craftslab/githubsearch.svg?color=brightgreen)](https://github.com/craftslab/githubsearch/tags)



## Introduction

*GitHub Search* is a GitHub search tool written in Go.



## Features

*GitHub Search* supports:

- [GitHub GraphQL API v4](https://developer.github.com/v4/)
- [GitHub REST API v3](https://developer.github.com/v3/)



## Examples

```
githubsearch \
  --api "graphql" \
  --config "config/search.json" \
  --output "output.json" \
  --query "code:runSearch,owner:craftslab,repo:githubsearch"

githubsearch \
  --api "rest" \
  --config "config/search.json" \
  --output "output.json" \
  --query "code:runSearch,owner:craftslab,repo:githubsearch"
```



## License

[Apache 2.0](LICENSE)
