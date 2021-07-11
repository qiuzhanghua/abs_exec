# Absolute Path of Executed File

## Build

```bash
go build
```

## Usage

```bash
#!/bin/bash
pushd `pwd` > /dev/null
cd `dirname $0`
exec_path=$("./abs_exec")
if [ "$?" != "0" ]; then
  echo "Can't find executed file path."
  exit 1
fi
echo "$exec_path"
popd > /dev/null
```