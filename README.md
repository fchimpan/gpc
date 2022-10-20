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

## usage

| param             | required | value                                                            |
| ----------------- | -------- | ---------------------------------------------------------------- |
| -t, --title       | yes      | generated page title                                             |
| -s, --spaceKey    | yes      | confluence space key to generate page                            |
| -d, --domain      | yes      | confluence domain. e.g. https://`domain`.atlassian.net/wiki/home |
| -p, --parent      |          | parent page ID                                                   |
| -b, --body        |          | body of the page                                                 |
| -c, --credentials |          | chose credential. default: [default]                             |
| --debug           |          | debug flag                                                       |

```bash
% gpc -t 'page title' -s 'my-space' -d 'my-domain'                         
# page generation succeeded!!
# title: page title
# https://my-domain.atlassian.net/wiki/spaces/my-space/pages/xxxx
```
