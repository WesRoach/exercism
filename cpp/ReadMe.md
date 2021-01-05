# C++ Overview

## Running Tests in Linux

Assuming current exercise is `bob` and we're in the exercise directory:

```bash
touch bob.{h,cpp}
mkdir build
cd build
cmake -G "Unix Makefiles" ..
make
```
