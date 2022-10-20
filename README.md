# A tool to generate pages in confluence(gpc)

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

1. create `profile`

```bash
vim ~/.gpc/profile
[<profile name>]
space_key = confluence space key to generate page  
domain = confluence domain. e.g. https://`domain`.atlassian.net/wiki/home 
parent = parent page ID  # optional
```

## usage

| param             | required | value                                |
| ----------------- | -------- | ------------------------------------ |
| -t, --title       | yes      | generated page title                 |
| -p, --profile     | yes      | profile name                         |
| -b, --body        |          | body of the page                     |
| -c, --credentials |          | chose credential. default: [default] |
| --debug           |          | debug flag                           |

```bash
% gpc -p 'profile' -t 'page title'
# page generation succeeded!!
# title: page title
# https://my-domain.atlassian.net/wiki/spaces/my-space/pages/xxxx
```
