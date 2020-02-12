# githubsearch

[![Build Status](https://travis-ci.com/craftslab/githubsearch.svg?branch=master)](https://travis-ci.com/craftslab/githubsearch)
[![Coverage Status](https://coveralls.io/repos/github/craftslab/githubsearch/badge.svg?branch=master)](https://coveralls.io/github/craftslab/githubsearch?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/craftslab/githubsearch)](https://goreportcard.com/report/github.com/craftslab/githubsearch)
[![License](https://img.shields.io/github/license/craftslab/githubsearch.svg?color=brightgreen)](https://github.com/craftslab/githubsearch/blob/master/LICENSE)
[![Tag](https://img.shields.io/github/tag/craftslab/githubsearch.svg?color=brightgreen)](https://github.com/craftslab/githubsearch/tags)



## Introduction

*GitHub Search* is a GitHub search tool written in Go.



## Feature

*GitHub Search* supports:

- [~~GitHub GraphQL API v4~~](https://developer.github.com/v4/)

- [GitHub REST API v3](https://developer.github.com/v3/)

  - [Search on GitHub](https://help.github.com/en/github/searching-for-information-on-github)
  - [Search Query](https://developer.github.com/v3/search/#constructing-a-search-query)
  - [License Type](https://help.github.com/en/github/creating-cloning-and-archiving-repositories/licensing-a-repository#searching-github-by-license-type)



## Usage

```bash
usage: githubsearch --api=API --config=CONFIG --output=OUTPUT --qualifier=QUALIFIER --search=SEARCH [<flags>]

GitHub Search

Flags:
      --help                 Show context-sensitive help (also try --help-long
                             and --help-man).
      --version              Show application version.
  -a, --api=API              API type, format: graphql or rest
  -c, --config=CONFIG        Config file, format: .json
  -o, --output=OUTPUT        Output file, format: .json
  -q, --qualifier=QUALIFIER  Qualifier list, format:
                             {qualifier}:{query},{qualifier}:{query},...
  -s, --search=SEARCH        Search list, format: code:{text} or repo:{text}
```

`-q/--qualifier`: See ["Searching on GitHub"](https://help.github.com/articles/searching-on-github/) for a complete list of available qualifiers, their format,
 and an example of how to use them.



## Setting

*GitHub Search* parameters can be set in the directory config.

An example of configuration in [search.json](https://github.com/craftslab/githubsearch/blob/master/config/search.json):

```bash
{
  "rest": {
    "order": "desc",
    "page": 2,
    "per_page": 10,
    "sort": "stars"
  }
}
```



## Running

```
githubsearch \
  --api "rest" \
  --config "config/search.json" \
  --output "output.json" \
  --qualifier "in:file,language:go,repo:githubsearch,user:craftslab" \
  --search "code:runSearch"
```

```
githubsearch \
  --api "rest" \
  --config "config/search.json" \
  --output "output.json" \
  --qualifier "language:go,user:craftslab" \
  --search "repo:githubsearch"
```



## License

Project License can be found [here](LICENSE).
