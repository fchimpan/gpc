# gpc

A tool to generate pages in confluence

## install

```bash
go install
```

## setup

```bash
export CONFLUENCE_API_TOKEN=<your api token>
export CONFLUENCE_USER_EMAIL=<confluence user email>
```

## usage

| param          | required | value                                                            | 
| -------------- | -------- | ---------------------------------------------------------------- | 
| -t, --title    | yes      | generated page title                                             | 
| -s, --spaceKey | yes      | confluence space key to generate page                            | 
| -d, --domain   | yes      | confluence domain. e.g. https://`domain`.atlassian.net/wiki/home | 
| -p, --parent   |          | parent page ID                                                   | 
| -b, --body     |          | body of the page                                                 | 
| --debug        |          | debug flag                                                       | 

```bash
% gpc main.go -t 'page titlr' -s 'my-space' -d 'my-domain' -p '12345' -b 'hoge'                              
# page generation succeeded!!
# title: page title
# https://my-domain.atlassian.net/wiki/spaces/my-space/pages/xxxx
```
