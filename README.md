# Absolute Path of Executed File

## Usage
```bash
#!/bin/bash
exec_path=$("./abs_exec")
if [ "$?" != "0" ]; then
  echo "Can't find executed file path!"
  exit 1
fi
echo "$exec_path"
```