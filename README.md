# Hive Plugin for Mach Composer 

This repository contains the hive plugin for Mach Composer. It requires Mach Composer 2.x

## Usage

```yaml
mach-composer:
  # ...
  plugins:
    hive:
      source: mach-composer/hive
      version: 0.1.0

global:
  # ...

sites:
  - identifier: my-site

    hive:
      token: your-token

    components:
        # ...
```
