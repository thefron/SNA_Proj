#/bin/bash
mkdir -p ./$2
for file in $(ls ./$1/????-??-??)
do
    filename=$(basename $file)
    cut -d " " -f "8,7" "$file" | grep -v "^$" | sed "s/^\([^ ]*\) \([^ ]*\)/\2 \1/" | sort -u > "./$2/$filename.sort"
done
sort -u -m ./$2/*.sort > ./$2/users
