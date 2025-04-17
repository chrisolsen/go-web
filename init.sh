#!/usr/bin/bash

echo "What is your app's name?:"
read app_name

function findAndReplace () {
    regex="s/$1/$2/g"
    find . -type f -not -path "./.git/*" -not -path "./init.sh" -exec sed -i $regex {} \;
}

# remove original .env file used for template dev
git rm .env
git commit -m "remove original .env file"

# re-add env file
touch .env
echo "# Env variables" > .env

findAndReplace "chrisolsen-goweb" $app_name
echo "Complete!"