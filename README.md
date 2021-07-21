myCLI is both a library for creating your own CLI  easily of your project.

# Table of Contents
- [Table of Contents](#table-of-contents)
- [Overview](#overview)
- [Concepts](#concepts)
- [Installing](#installing)
- [Usage](#usage)
- [License](#license)

# Overview

myCLI is a library providing a simple CLI create your own based in YAML.

myCLI provides:
* Easy Setup.
* Easy generation files
  - Prefix files name
  - Suffix files name
  - Initial template

myCLI is builded using [cobra](https://github.com/spf13/cobra)


# Concepts

myCli is built based on a  YAML file.

```yaml
project:
  extension: php 
modules:
  model:
    description: "create class models"
    path: "models/"
    prefix: ""
    suffix: "_model"
    template: "templates/start.ph2p"
  controllers:
    command: ctl
    description: "create class controller"
    path: "controllers/"
    prefix: ""
    suffix: "_controller"
``` 
- `extension`: the files created will be with this extension.
- `modules`: are the CLI commands 

**commands** 
```yaml
 model:
    command: mdl
    description: "create class models"
    path: "models/"
    prefix: ""
    suffix: "_model"
    template: "templates/start.ph2p"
``` 
- `model`: by default the command name is the structure name, in this case model. 
- `command`: customize command name in CLI. 
- `description`: command description. 
- `path`: path where files will be created. 
- `prefix`: the prefix will be added to the name of the created example file: prefix_filename.php 
- `suffix`: the suffix will be added to the name of the created example file: filename_suffix.php
- `template`: You can define an initial template for when the file is created
  


# Installing
coming soon

# Usage

coming soon

# License
myCLI is released MIT license. See [LICENSE.txt](https://github.com/izacximenes/mycli/blob/master/LICENSE.txt)
