// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package config

const InitialGlobalConfig = `repo:
  # cron usage: https://godoc.org/github.com/robfig/cron#hdr-Usage
  #
  #   "@every 1h30m" means Every hour thirty
  #   "@hourly" means Every hour
  #   "0 30 * * * *" means Every hour on the half hour
  #
  #	  Field name   | Mandatory? | Allowed values  | Allowed special characters
  #	  ----------   | ---------- | --------------  | --------------------------
  #	  Seconds      | Yes        | 0-59            | * / , -
  #	  Minutes      | Yes        | 0-59            | * / , -
  #	  Hours        | Yes        | 0-23            | * / , -
  #	  Day of month | Yes        | 1-31            | * / , - ?
  #	  Month        | Yes        | 1-12 or JAN-DEC | * / , -
  #	  Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
  #
  cron: "0 30 4 * * *"
cluster:
  image_url: https://openpitrix.pek3a.qingstor.com/image/ubuntu.tar.gz
  plugins:
    - qingcloud
    - kubernetes
  frontgate_conf: '{"app_id":"app-ABCDEFGHIJKLMNOPQRST","version_id":"appv-ABCDEFGHIJKLMNOPQRST","name":"frontgate","description":"OpenPitrixbuilt-infrontgateservice","subnet":"","nodes":[{"container":{"type":"docker","image":"openpitrix/openpitrix:metadata"},"count":1,"cpu":1,"memory":1024,"volume":{"size":10,"mount_point":"/data","filesystem":"ext4"}}]}'
pilot:
  ip: 127.0.0.1
runtime:
  qingcloud_pek3a:
    api_server: api.qingcloud.com
    zone: pek3a
    image_id: xenial4x64a
  qingcloud_sh1a:
    api_server: api.qingcloud.com
    zone: sh1a
    image_id: xenial4x64a
`
