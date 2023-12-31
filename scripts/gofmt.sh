#!/usr/bin/env bash

## green to echo
function green(){
    echo -e "\033[32m $1 \033[0m"
}

## Error
function bred(){
    echo -e "\033[31m\033[01m $1 \033[0m"
}

files=()
files_need_gofmt=()
files_need_goimports=()

if [[ $# -ne 0 ]]; then
    for arg in "$@"; do
        if [ -f "$arg" ];then
            files+=("$arg")
        fi

        if [ -d "$arg" ];then
            for file in `find $arg -type f | grep "\.go$" | grep -v -e mock -e proto`; do
                files+=("$file")
            done
        fi
    done
fi

# Check for files that fail gofmt and goimports.
if [[ "${#files[@]}" -ne 0 ]]; then
    for file in "${files[@]}"; do
        diff="$(gofmt -s -d ${file} 2>&1)"
        if [[ -n "$diff" ]]; then
            files_need_gofmt+=("${file}")
        fi

        diff2="$(goimports -s -d ${file} 2>&1)"
        if [[ -n "$diff2" ]]; then
          files_need_goimports+=("${file}")
        fi
    done
fi

# gofmt
if [[ "${#files_need_gofmt[@]}" -ne 0 ]]; then
    bred "ERROR!"
    for file in "${files_need_gofmt[@]}"; do
        gofmt -s -d ${file} 2>&1
    done
    echo ""
    bred "Some files have not been gofmt'd. To fix these errors, "
    bred "copy and paste the following:"
    for file in "${files_need_gofmt[@]}"; do
        bred "  gofmt -s -w ${file}"
    done
    exit 1
else
    green "OK"
    exit 0
fi

# goimports
if [[ "${#files_need_goimports[@]}" -ne 0 ]]; then
    bred "ERROR!"
    for file in "${files_need_goimports[@]}"; do
        gofmt -s -d ${file} 2>&1
    done
    echo ""
    bred "Some files have not been goimports'd. To fix these errors, "
    bred "copy and paste the following:"
    for file in "${files_need_goimports[@]}"; do
        bred "  goimports -s -w ${file}"
    done
    exit 1
else
    green "OK"
    exit 0
fi