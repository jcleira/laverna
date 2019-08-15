# Laverna

Laverna is a tool that will expose an API feeded from a YAML file.

![Imgur](https://i.imgur.com/GL6e58c.jpg)

## Purpose

Larverna has been created to quick implement (mock?) an API.

## API definition

Laverna needs a YAML file to provide an API, the following example show all Lavernas's features:

```yaml
---
- endpoint:
  url: /steps
  status_code: 200
  response: 
  - step:
    description: Good or Evil?
    options:
    - option:
      value: Good 
    - option: Good
      value: Evil 
```

That YAML file will create the following JSON endpoint response:

```json
// Endpoint: /steps
// StatusCode: 200
// Response:
[{
  "description": "Good or Evil?",
  "options": [{
    "value": "Good"
  },{
    "value": "Evil"
  }]
}]
```
