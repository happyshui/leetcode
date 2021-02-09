#!/bin/bash
# 1.
sed "s/ /\n/g" words.txt | sort -n |uniq -c | sort -k 1 -r | awk '{print $2" "$1}' | grep -v ^" "
# 2.
cat words.txt| tr -s " " "\n" | sort -n |uniq -c | sort -k 1 -r | awk '{print $2" "$1}'
