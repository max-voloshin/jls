jls
===

JSON edition of ``ls`` util

Usage of ``jls``:

```
jls [file ...]
```

Example:

```
$ jls
{"name":".git","path":"/Users/maxvoloshin/golang/src/github.com/max-voloshin/jls/.git","size":510,"dir":true,"file":false,"append_only":false,"exclusive":false,"temporary":false,"symlink":false,"device":false,"named_pipe":false,"socket":false,"setuid":false,"setgid":false,"char_device":false,"sticky":false,"owner_read":true,"owner_write":true,"owner_execure":true,"group_read":true,"group_write":false,"group_execute":true,"other_read":true,"other_write":false,"other_execute":true,"permissions":"-rwxr-xr-x","mtime":1402320791,"ctime":1402320791,"atime":1402320615,"uid":501,"gid":20}
{"name":".gitignore","path":"/Users/maxvoloshin/golang/src/github.com/max-voloshin/jls/.gitignore","size":259,"dir":false,"file":true,"append_only":false,"exclusive":false,"temporary":false,"symlink":false,"device":false,"named_pipe":false,"socket":false,"setuid":false,"setgid":false,"char_device":false,"sticky":false,"owner_read":true,"owner_write":true,"owner_execure":false,"group_read":true,"group_write":false,"group_execute":false,"other_read":true,"other_write":false,"other_execute":false,"permissions":"-rw-r--r--","mtime":1402320789,"ctime":1402320789,"atime":1402320791,"uid":501,"gid":20}
{"name":"jls.go","path":"/Users/maxvoloshin/golang/src/github.com/max-voloshin/jls/jls.go","size":3960,"dir":false,"file":true,"append_only":false,"exclusive":false,"temporary":false,"symlink":false,"device":false,"named_pipe":false,"socket":false,"setuid":false,"setgid":false,"char_device":false,"sticky":false,"owner_read":true,"owner_write":true,"owner_execure":false,"group_read":true,"group_write":false,"group_execute":false,"other_read":true,"other_write":false,"other_execute":false,"permissions":"-rw-r--r--","mtime":1402320789,"ctime":1402320789,"atime":1402320791,"uid":501,"gid":20}
{"name":"LICENSE","path":"/Users/maxvoloshin/golang/src/github.com/max-voloshin/jls/LICENSE","size":1078,"dir":false,"file":true,"append_only":false,"exclusive":false,"temporary":false,"symlink":false,"device":false,"named_pipe":false,"socket":false,"setuid":false,"setgid":false,"char_device":false,"sticky":false,"owner_read":true,"owner_write":true,"owner_execure":false,"group_read":true,"group_write":false,"group_execute":false,"other_read":true,"other_write":false,"other_execute":false,"permissions":"-rw-r--r--","mtime":1402320789,"ctime":1402320789,"atime":1402320791,"uid":501,"gid":20}
{"name":"README.md","path":"/Users/maxvoloshin/golang/src/github.com/max-voloshin/jls/README.md","size":92,"dir":false,"file":true,"append_only":false,"exclusive":false,"temporary":false,"symlink":false,"device":false,"named_pipe":false,"socket":false,"setuid":false,"setgid":false,"char_device":false,"sticky":false,"owner_read":true,"owner_write":true,"owner_execure":false,"group_read":true,"group_write":false,"group_execute":false,"other_read":true,"other_write":false,"other_execute":false,"permissions":"-rw-r--r--","mtime":1402320916,"ctime":1402320918,"atime":1402320916,"uid":501,"gid":20}
```

Properties
==========

General
-------

* ``name``
* ``path``
* ``size`` (bytes)

Mode
----

* ``dir`` (bool)
* ``file`` (bool)
* ``append_only`` (bool)
* ``exclusive`` (bool)
* ``temporary`` (bool)
* ``symlink`` (bool)
* ``device`` (bool)
* ``named_pipe`` (bool)
* ``socket`` (bool)
* ``char_device`` (bool)

Permissions
-----------

* ``uid`` (int) – owner id
* ``gid`` (int) – group id
* ``permissions`` (string) – example: ``-rwxr-xr-x``
* ``sticky`` (bool)
* ``setuid`` (bool)
* ``setgid`` (bool)
* ``owner_read`` (bool)
* ``owner_write`` (bool)
* ``owner_execure`` (bool)
* ``group_read`` (bool)
* ``group_write`` (bool)
* ``group_execute`` (bool)
* ``other_read`` (bool)
* ``other_write`` (bool)
* ``other_execute`` (bool)

Time
----

* ``mtime`` (int) – UNIX time of modification
* ``ctime`` (int) - UNIX time of creation
* ``atime`` (int) - UNIX time of access

Acknowledgements
================

This project uses idea of [json4shell](https://github.com/amarao/json4shell) project.

License
=======

MIT
