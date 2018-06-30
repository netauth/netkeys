NetKeys
=======

NetKeys is a convenient application to print out ssh keys for sshd to
consume.  To pull keys from the directory, place the following in your
sshd_config:

```
AuthorizedKeysCommandUser <some_dedicated_local_user>
AuthorizedKeysCommand /path/to/NetKeys --ID %u
```

Keep in mind that keys are not cached locally so if the network is
partitioned then no keys will be available.  Plan for this accordingly
with backup keys.
