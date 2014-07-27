#!/usr/bin/env bash
#

pushd /home/cjlyth/go-work/src/github.com/cjlyth &> /dev/null

goapp deploy -oauth  go-resume-web/app.yaml go-resume-api/api.yaml
appcfg.py --oauth2 update_dispatch go-resume-web