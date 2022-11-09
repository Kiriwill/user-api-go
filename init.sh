projectName=${PWD##*/}

echo "creating folders"
mkdir cmd cmd/$projectName pkg misc

echo "renaming templates with project names"
find . -type f -exec sed -i "s/<PROJECT.NAME>/$projectName/g" {} +

go mod init

echo "removing the init file"
rm -rf ./init.sh
