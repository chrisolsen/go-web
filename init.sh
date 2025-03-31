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

if [ "$confirm" = "y" ]; then
    find . -type f -not -path "./.git/*" -not -path "./init.sh" -exec sed -i "s/__APP__/${app_name}/g" {} \;
    find . -type f -not -path "./.git/*" -not -path "./init.sh" -exec sed -i "s/__DB__/${app_name}/g" {} \;
    find . -type f -not -path "./.git/*" -not -path "./init.sh" -exec sed -i "s/__PKG__/${app_name}/g" {} \;
    find . -type f -not -path "./.git/*" -not -path "./init.sh" -exec sed -i "s/__PORT__/${port}/g" {} \;

    echo "Complete!"
else
    echo "Cancelled!"
fi

