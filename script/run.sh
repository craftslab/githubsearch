#!/bin/bash

# GET https://api.github.com/search/code?q={query}{&page,per_page,sort,order}
# GET https://api.github.com/search/repositories?q={query}{&page,per_page,sort,order}
# GET https://api.github.com/repos/{owner}/{repo}
#
# NOTES: max per_page: 100
#
# Example:
#
# curl https://api.github.com/search/code?q=runSearch+in:file+language:go+repo:githubsearch+user:craftslab&page=1&per_page=10&sort=stars&order=desc
# curl https://api.github.com/search/repositories?q=githubsearch+language:go+user:craftslab&page=1&per_page=10&sort=stars&order=desc

go run main.go --api "rest" --config "config/search.json" --output "output.json" --qualifier "in:file,language:go,repo:githubsearch,user:craftslab" --search "code:runSearch"
go run main.go --api "rest" --config "config/search.json" --output "output.json" --qualifier "language:go+user:craftslab" --search "repo:githubsearch"
