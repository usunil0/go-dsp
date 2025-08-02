#!/bin/sh

# entrypoint.sh
set -e

air --root "." \
  --build.args_bin="" \
  --build.bin="$1" \
  --build.cmd="go build -o $1 $2" \
  --build.delay="1000" \
  --build.exclude_dir="assets,tmp,vendor,testdata,test,mocks" \
  --build.exclude_file="" \
  --build.exclude_regex="_test.go" \
  --build.exclude_unchanged="false" \
  --build.follow_symlink="false" \
  --build.full_bin="" \
  --build.include_dir="" \
  --build.include_ext="go,tpl,tmpl,html" \
  --build.kill_delay="0s" \
  --build.log="build-errors.log" \
  --build.send_interrupt="false" \
  --build.stop_on_error="true" \
  --color.app="" \
  --color.build="yellow" \
  --color.main="magenta" \
  --color.runner="green" \
  --color.watcher="cyan" \
  --log.time="false" \
  --misc.clean_on_exit="false" \
  --screen.clear_on_rebuild="false" \
  --tmp_dir="tmp" \
  --testdata_dir="testdata"
