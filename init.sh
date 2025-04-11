#!/usr/bin/bash

echo "What is your app's name?:"
read app_name

function findAndReplace () {
    regex="s/$1/$2/g"
    find . -type f -not -path "./.git/*" -not -path "./init.sh" -exec sed -i $regex {} \;
}

findAndReplace "chrisolsen-goweb" $app_name
echo "Complete!"