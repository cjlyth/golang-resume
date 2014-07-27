#!/usr/bin/env bash
#

pushd $HOME/go-work/src/github.com/cjlyth &> /dev/null

goapp serve go-resume-web/dispatch.yaml go-resume-web/app.yaml go-resume-api/api.yaml


