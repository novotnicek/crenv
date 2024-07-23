# What's is crenv?

**crenv** is simple tool that merges
environment variables from some .env file
with variables from parent shell process
and then save it into destination file.

## Example
```sh
$ env | grep USER
USERNAME=peter
USER=peter

$ cat .env
APP_ENV=dev
USER=foobar

$ crenv .env .env.local
Source: .env
Destination: .env.local
Key: APP_ENV, Value: dev
Key: USER, Value: peter
====================================
.env.local created successfully
====================================

$ cat .env.local
APP_ENV=dev
USER=peter
```
