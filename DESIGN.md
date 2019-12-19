# Iridium Design

## Background

Google and Apple offer wonderful services which allow people to share and store photos/videos with relative ease. Unfortunately, they also scrape the shit out of them with complete disregard for you or anyone elses privacy.

Where we go, what we see, what we buy, who we're with... we are the commodity.

Even for the technically inclined, it's not easy to break out of their grasp.

## Problem Space & Scope

The goal of *Iridium* is to provide a self-hosted platform that enables users to break away from Google Photos and Apple iCloud.

*Iridium* will do the following:

* Provide a free and open alternative to Google Photos
* Provide a way for end users to quickly deploy this software to infrastructure of their choosing
* Provide a location to store all of your photos/videos
* Provide a way to move all of your media from Google/Apple into Iridium
* Provide user management
* Provide an IOS and Android app (IRONY\*) that allows you to automatically upload photos you've taken to your *Iridium* installation

\* Can't win 'em all.

## End-Users

There are 2 groups of users that we will be supporting. The first are fellow nerds that will be deploying and maintaining their own *Iridium* installations. The second are the users of those installations which are quite similar to the proverbial end-user, ie: non-technical.

Both groups should feel equally comfortable.

## Success Criteria

A user from the nerd group should be able to standup a new *Iridium* installation and get it to a usable state within 30 minutes.

A user from the proverbial end-user group should be able to start using the new installation with nothing more than a URL given to them by a member of the nerd group.

## System Technical Design

* Terraform for infra
* Golang for services
* JavaScript for UI
* nginx to serve the UI
* Postgres
* Redis for temp storage such as tokens
* S3 or similar for file storage

### Overview

An Iridium deployment will have 4 things active:

1. Service container(s) running the Golang services
1. Nginx container to serve the UI
1. Redis container to handle temp storage (ie: auth tokens)
1. Postgres container/RDS

Media storage will be handled, for now, by S3.

### Data Flow

![Data flow][https://raw.githubusercontent.com/bradj/iridium/design/doc-updates/assets/flow.svg]


### Service Container(s)

TODO
