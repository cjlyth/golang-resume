#!/usr/bin/env bash
#

pushd /home/cjlyth/go-work/src/github.com/cjlyth &> /dev/null

goapp serve go-resume-web/dispatch.yaml go-resume-web/app.yaml go-resume-api/api.yaml


