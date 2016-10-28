#! /bin/sh

if [ $# -lt 1 ]; then
    echo "Usage ./start.sh package"
    exit 0
fi

package="$(echo $1 |sed 's/\//\\\//g')"

FILES="$(find . -type f -name \*.go)"
for file in $FILES; do 
	sed -i '' "s/{{.package}}/$package/g" $file
done;

echo 'done'