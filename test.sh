#!/bin/bash

create="POST"
read="GET"
update="PUT"
remove="DELETE"

base_url="http://localhost:8080"

key="status"
value="happy"

echo "Methodius testing"
echo "============================="

echo "Create"
curl -X ${create} "${base_url}/${key}" -d "${value}"
echo -e "\n"

echo "Read list"
curl -X ${read} "${base_url}"
echo -e "\n"

echo "Read detail"
curl -X ${read} "${base_url}/${key}"
echo -e "\n"

echo "Update"
curl -X ${update} "${base_url}/${key}" -d "${value} the new way"
echo -e "\n"

echo "Read detail"
curl -X ${read} "${base_url}/${key}"
echo -e "\n"

echo "Remove"
curl -X ${remove} "${base_url}/${key}"
echo -e "\n"

echo "Read list"
curl -X ${read} "${base_url}"
echo -e "\n"

echo "Quit"
curl -X QUIT "${base_url}"
echo -e "\n"
