gosubrename
=============

This is a simple project to rename subtitles files when they are not aligned with their video files.


Goal
-------
We can see a test directory here:
![Alt Text](https://github.com/benoitantelme/gosubrename/raw/master/screenshots/initialDir.png)

The application starts by getting a directory full path and optionally two extensions as arguments.
It will then rename the subtitles if they are not equals to the corresponding episode title and create a backup for the initial subtitles:
![Alt Text](https://github.com/benoitantelme/gosubrename/raw/master/screenshots/finalDir.png)

It also backs up the subtitles in a directory:
![Alt Text](https://github.com/benoitantelme/gosubrename/raw/master/screenshots/backupDir.png)

Command line example:
![Alt Text](https://github.com/benoitantelme/gosubrename/raw/master/screenshots/cli.png)

Tech stack
-------
It uses golang and the packages:

* "fmt"
* "os"
* "io"
* "io/ioutil"
* "path/filepath"
* "regexp"
* "errors"
* "testing"