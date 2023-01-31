@echo off

set cnt = 0

for %%x in (%*) do Set /A cnt+=1

if %cnt% equ 2 (
    docker login
    docker build -t ansurfen/%1:%2 .
    docker push ansurfen/%1:%2
) else (
    echo "The argument must is 2"
)
