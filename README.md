# gitodo
> Lists your readme todos on github.  

## Usage 
* Create [Github token](https://github.com/settings/tokens/new) with `repo` scope

`$ go get github.com/akramsaouri/gitodo`

`$ GITHUB_ACCESS_TOKEN="xxxxxxxxxxxxxxxxxxxxxx" gitodo`

## Limitations
- Only look for readmes under the `master` branch.
- Is specifically looking for the `## TODO` markdown as title for todos.

## TODO
- Add concurrecy for parsing readmes.
- Skip archived repos.
