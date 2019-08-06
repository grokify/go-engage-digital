# RingCentral Engage SDK for Go

[![Build Status][build-status-svg]][build-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
![API Coverage][api-coverage-svg]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]
[![Stack Overflow][stackoverflow-svg]][stackoverflow-link]
[![Twitter][twitter-svg]][twitter-link]

 [api-coverage-svg]: https://img.shields.io/badge/api%20coverage-101%2F121%20%3D%2083%25-yellow.svg
 [build-status-svg]: https://api.travis-ci.org/grokify/go-ringcentral-engage.svg?branch=master
 [build-status-link]: https://travis-ci.org/grokify/go-ringcentral-engage
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-ringcentral-engage
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/go-ringcentral-engage
 [docs-godoc-svg]: https://img.shields.io/badge/docs-godoc-blue.svg
 [docs-godoc-link]: https://godoc.org/github.com/grokify/go-ringcentral-engage/engagedigital
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/go-ringcentral-engage/blob/master/LICENSE
 [stackoverflow-svg]: https://img.shields.io/badge/Stack%20Overflow-ringcentral-orange.svg
 [stackoverflow-link]: https://stackoverflow.com/questions/tagged/ringcentral
 [twitter-svg]: https://img.shields.io/twitter/follow/ringcentraldevs.svg?style=social&label=follow
 [twitter-link]: https://twitter.com/RingCentralDevs

## Overview

This currently provides a minimal Go SDK for RingCentral Engage Digital, formerly Dimelo:

https://www.dimelo.com

## Pre-requisites

You need the following:

1. RingCentral Engage Digital / Dimelo Account and the URL subdomain, e.g. `my-subdomain` in the following URL: `https://my-subdomain.engagement.dimelo.com`
1. An API Access Token. You can retrive this at the following URL, using your own domain in place of `my-subdomain`: `https://my-subdomain.engagement.dimelo.com/admin/api_access_tokens`

## Installation

```
$ go get github.com/grokify/go-ringcental-engage/...
```

## Usage

The main SDK is in the `engagedigital` folder. The `utils` folder provides a helper function to instantiate the SDK:

```go
import("github.com/grokify/go-ringcentral-engage/utils")

func main() {
    apiClient := utils.NewApiClient("my-subdomain", "my-access-token")

    // helper function to get access to raw `*http.Client`
    httpClient := apiClient.HTTPClient()
}
```

For information on how to use the `apiClient` object, see:

1. [`examples`](examples) folder for usage. [`examples/simple_get`](examples/simple_get) includes a simple functions for a lot of the `GET` APIs.
1. [engagedigital/README.md](https://github.com/grokify/go-ringcentral-engage/blob/master/engagedigital/README.md) for SDK documentation.

## Documentation

[engagedigital/README.md](https://github.com/grokify/go-ringcentral-engage/blob/master/engagedigital/README.md)

## Coverage

101/121 APIs - 83%

- [x] Communities
  - [x] GET /1.0/communities
  - [x] GET /1.0/communities/:id

- [x] Sources
  - [x] GET /1.0/content_sources
  - [x] GET /1.0/content_sources/:id
  - [x] PUT /1.0/content_sources/:id

- [x] Folders
  - [x] GET /1.0/folders
  - [x] GET /1.0/folders/:id
  - [x] POST /1.0/folders
  - [x] PUT /1.0/folders/:id
  - [x] DELETE /1.0/folders/:id

- [x] Roles
  - [x] GET /1.0/roles
  - [x] GET /1.0/roles/:id
  - [x] POST /1.0/roles
  - [x] PUT /1.0/roles/:id

- [x] Categories
  - [x] GET /1.0/categories
  - [x] GET /1.0/categories/:id
  - [x] POST /1.0/categories
  - [x] PUT /1.0/categories/:id
  - [x] DELETE /1.0/categories/:id

- [x] Tags
  - [x] GET /1.0/tags
  - [x] GET /1.0/tags/:id
  - [x] POST /1.0/tags
  - [x] PUT /1.0/tags/:id
  - [x] DELETE /1.0/tags/:id

- [x] Teams
  - [x] GET /1.0/teams
  - [x] GET /1.0/teams/:id
  - [x] POST /1.0/teams
  - [x] PUT /1.0/teams/:id
  - [x] DELETE /1.0/teams/:id

- [x] Users
  - [x] GET /1.0/users/me
  - [x] GET /1.0/users
  - [x] GET /1.0/users/:id
  - [x] POST /1.0/users
  - [x] PUT /1.0/users/:id
  - [x] POST /1.0/users/invite
  - [x] DELETE /1.0/users/:id

- [ ] User Source Permissions
  - [ ] GET /1.0/users/:id/sources_permissions
  - [ ] PUT /1.0/users/:id/sources_permissions

- [x] Identities
  - [x] GET /1.0/identities
  - [x] GET /1.0/identities/:id

- [x] Identity Groups
  - [x] GET /1.0/identity_groups
  - [x] GET /1.0/identity_groups/:id
  - [x] PUT /1.0/identity_groups/:id

- [x] Custom Fields
  - [x] GET /1.0/custom_fields
  - [x] GET /1.0/custom_fields/:id
  - [x] POST /1.0/custom_fields
  - [x] PUT /1.0/custom_fields/:id
  - [x] DELETE /1.0/custom_fields/:id

- [x] Threads
  - [x] GET /1.0/content_threads
  - [x] GET /1.0/content_threads/:id
  - [x] PUT /1.0/content_threads/:id/update_categories
  - [x] PUT /1.0/content_threads/:id/ignore
  - [x] PUT /1.0/content_threads/:id/close
  - [x] GET /1.0/content_threads/:id/open

- [x] Contents
  - [x] GET /1.0/contents
  - [x] GET /1.0/contents/:id
  - [x] POST /1.0/contents
  - [x] PUT /1.0/contents/:id/update_categories
  - [x] PUT /1.0/contents/:id/ignore

- [x] Attachments
  - [x] GET /1.0/attachments
  - [x] GET /1.0/attachments/:id
  - [x] POST /1.0/attachments

- [x] Events
  - [x] GET /1.0/events
  - [x] GET /1.0/events/:id

- [ ] Interventions
  - [x] GET /1.0/interventions
  - [x] GET /1.0/interventions/:id
  - [x] POST /1.0/interventions
  - [ ] PUT /1.0/interventions/:id
  - [x] PUT /1.0/interventions/:id/close
  - [x] PUT /1.0/interventions/:id/reassign
  - [x] PUT /1.0/interventions/:id/update_categories
  - [x] DELETE /1.0/interventions/:id/cancel

- [x] Intervention comments
  - [x] GET /1.0/intervention_comments
  - [x] GET /1.0/intervention_comments/:id
  - [x] POST /1.0/intervention_comments
  - [x] DELETE /1.0/intervention_comments/:id

- [x] Agent Status (task view)
  - [x] GET /1.0/status
  - [x] GET /1.0/status/:agent_id
  - [ ] PUT /1.0/status/:agent_id

- [x] Webhook
  - [x] GET /1.0/webhooks
  - [x] GET /1.0/webhooks/:id
  - [x] POST /1.0/webhooks
  - [x] PUT /1.0/webhooks/:id
  - [x] DELETE /1.0/webhooks/:id

- [x] Time Sheet
  - [x] GET /1.0/time_sheets
  - [x] GET /1.0/time_sheets/:id
  - [x] POST /1.0/time_sheets
  - [x] PUT /1.0/time_sheets/:id
  - [x] DELETE /1.0/time_sheets/:id

- [x] Channels
  - [x] GET /1.0/channels
  - [x] GET /1.0/channels/:id
  - [x] PUT /1.0/channels/:id

- [x] Settings
  - [x] GET /1.0/settings
  - [x] PUT /1.0/settings

- [x] Locales
  - [x] GET /1.0/locales

- [x] Timezones
  - [x] GET /1.0/timezones

- [x] Presence statuses
  - [x] GET /1.0/presence_status
  - [x] GET /1.0/presence_status/:id
  - [x] POST /1.0/presence_status
  - [x] PUT /1.0/presence_status/:id
  - [x] DELETE /1.0/presence_status/:id

- [x] Tasks
  - [x] GET /1.0/tasks
  - [x] GET /1.0/tasks/:id
  - [x] POST /1.0/tasks/:id/transfer
  - [x] POST /1.0/tasks/:id/move

- [ ] Reply Assistant - Entries
  - [ ] GET /1.0/reply_assistant/entries
  - [ ] GET /1.0/reply_assistant/entries/:id
  - [ ] POST /1.0/reply_assistant/entries
  - [ ] PUT /1.0/reply_assistant/entries/:id
  - [ ] DELETE /1.0/reply_assistant/entries/:id

- [ ] Reply Assistant - Versions
  - [ ] GET /1.0/reply_assistant/versions
  - [ ] GET /1.0/reply_assistant/versions/:id
  - [ ] POST /1.0/reply_assistant/versions
  - [ ] PUT /1.0/reply_assistant/versions/:id
  - [ ] DELETE /1.0/reply_assistant/versions/:id

- [ ] Reply Assistant - Groups
  - [ ] GET /1.0/reply_assistant/groups
  - [ ] GET /1.0/reply_assistant/groups/:id
  - [ ] POST /1.0/reply_assistant/groups
  - [ ] PUT /1.0/reply_assistant/groups/:id
  - [ ] DELETE /1.0/reply_assistant/groups/:id

- [ ] Survey Response
  - [ ] GET /1.0/survey_responses/:id

There are 127 endpoints. To count, use the following on OS-X:

```
$ grep ' [ ]' README.md  | wc -l
```

## Building the SDK

You won't normally need to do this unless you want to modify the SDK, like adding endpoints via the OpenAPI 2.0 / Swagger 2.0 specification.

This SDK is auto-generated from the [OpenAPI 2.0 / Swagger 2.0 spec](codegen/openapi-spec.json) using [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator).

> **Note:** This SDK uses a merged OpenAPI spec so do not edit the `codegen/openapi-spec.json` file. Instead, edit files in and add files to the `codegen/partial-specs` folder.

Run:

```
$ cd codegen
$ go run merge.go -v 3
$ sh generate.sh
$ rm -rf ../engagedigital
$ mv engagedigital ..
```
# Credits

Thanks to the following apps for making this possible:

* https://github.com/OpenAPITools/openapi-generator
* https://github.com/getkin/kin-openapi/tree/master/openapi3
* https://apidevtools.org/swagger-parser/online/
* https://mermade.org.uk/openapi-converter