#/bin/bash
mkdir -p ./$2
for file in $(ls ./$1/????-??-??)
do
    filename=$(basename $file)
    sed "s/^\([^ ]*\) \([^ ]*\) \([^ ]*\) .*/\3 \2/" "$file" | sort -u -n > "./$2/$filename.sort"
done
sort -u -m -n ./$2/*.sort > ./$2/users
