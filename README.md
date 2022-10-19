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

```bash
export CONFLUENCE_API_TOKEN=<your api token>
export CONFLUENCE_USER_EMAIL=<confluence user email>
```

## usage

| param          | required | value                                                                         | 
| -------------- | -------- | ----------------------------------------------------------------------------- | 
| -t, --title    | yes      | generated page title                                                          | 
| -s, --spaceKey | yes      | confluence space key to generate page                                         | 
| -d, --domain   | yes      | confluence domain. e.g. https://`domain`.atlassian.net/wiki/home              | 
| -p, --parent   |          | parent page ID. If not set, a page will be generated directly under the space | 
| -b, --body     |          | body of the page                                                              | 
| --debug        |          | debug flag                                                                    | 

```bash
% gpc -t 'page title' -s 'my-space' -d 'my-domain'                         
# page generation succeeded!!
# title: page title
# https://my-domain.atlassian.net/wiki/spaces/my-space/pages/xxxx
```
