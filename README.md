# mydocker

## Generating rootfs on ubuntu x86_64

```bash
sudo mkdir -p rootfs/{bin,lib,lib64,etc,proc,sys,dev,tmp}

# Copy bash 
sudo cp /bin/bash rootfs/bin/

# inspect ldd /bin/bash

# Copy required libraries into the corresponding paths
# for ex. on a typical x86_64 ubuntu system:

sudo cp /lib/x86_64-linux-gnu/libtinfo.so.6 rootfs/lib/
sudo cp /lib/x86_64-linux-gnu/libc.so.6 rootfs/lib/
sudo cp /lib64/ld-linux-x86-64.so.2 rootfs/lib64/
```