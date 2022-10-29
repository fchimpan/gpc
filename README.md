# A tool to generate pages in confluence(gpc)

![gpc-demo](https://user-images.githubusercontent.com/52129983/198816064-3f11fb58-9d25-4e1a-a2e9-4eaded223eab.gif)

## install

### brew

```bash
brew tap fchimpan/gpc
brew install fchimpan/gpc/gpc
```

### go install
```bash
go install github.com/fchimpan/gpc
```

## setup

1. make `$HOME/.gpc` directiry

```bash
mkdir ~/.gpc
```

2. create `credentials` file

```bash
vim ~/.gpc/credentials

[default]
confluence_api_token = <your api token>
confluence_email = <confluence user email>
```

3. create `config` file

```bash
vim ~/.gpc/config

[<profile name>]
space_key = confluence space key to generate page  
domain = confluence domain. e.g. https://`domain`.atlassian.net/wiki/home 
parent = parent page ID  # optional. If parent is not set, the page is generated directly under the space
```

## usage

| param             | type   | value                                         |
| ----------------- | ------ | --------------------------------------------- |
| -b, --body        | bool   | If this flag is set, page body can be entered |
| -c, --credentials | string | chose credential. default: [default]          |
| --debug           | bool   | debug flag                                    |

```bash
% gpc 

✔ dev                  # select parent
Page title: page title # input page title
# page generation succeeded!!
# title: page title
# https://my-domain.atlassian.net/wiki/spaces/my-space/pages/xxxx
```
