# gitodo
> List your readme todos on github.  

## Usage 
`$ go get https://github.com/akramsaouri/gitodo`

`$ GITHUB_ACCESS_TOKEN="xxxxxxxxxxxxxxxxxxxxxx" gitodo`

## Limitations
- Only look for readmes under the `master` branch.
- Is specifically looking for the `## TODO` markdown as title for todos.

## TODO
- Add concurrecy for parsing readmes.
