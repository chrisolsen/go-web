#!/usr/bin/bash

echo "What is your app's name?:"
read app_name

echo "What is your database's name?:"
read db_name

echo "What is your package's name? [foo.com/bar]:"
read package_name

echo "What is the http port?"
read port 

echo "Proceed with naming? ${app_name}:${package_name} (y/N):"
read confirm

function findAndReplace () {
    regex="s/$1/$2/g"
    find . -type f -not -path "./.git/*" -not -path "./init.sh" -exec sed -i $regex {} \;
}

if [ "$confirm" = "y" ]; then
    findAndReplace "__APP__" $app_name
    findAndReplace "__DB__" $db_name
    findAndReplace "__PKG__" $package_name
    findAndReplace "__PORT__" $port

    echo "Complete!"
else
    echo "Cancelled!"
fi
