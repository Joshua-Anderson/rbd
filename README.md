### RBD - Remote build tool

`rbd` is a tool for developing code on your laptop but building that code on your workstation.

When you run `rbd make build`, `rbd`:

- Copies the current directory from your laptop to a specified location on your workstation, using rsync over ssh.
- Runs the given command, `make build`, in this directory on your workstation, using ssh.
- Copies the directory back from your workstation to your laptop, using rsync over ssh.

:warning: This tool is not very fancy, it's specifically designed to fit my use case and little more. :warning:

### Configuration:

`rbd` is configured by the file `~/.rbd/config.json`

```json
{
  "workers": [{
      "id": "ws",
      "host": "example.com",
      "user": "test",
      "port": 2222 // optional
  }],
  "mappings": [
    {
      "worker": "ws",
      "remote": "/remote/builder",
      "local": "/home/test/testprogram"
    },
  ]
}
```

You define a worker, which a ssh connection information to your remote build machine.
You can then map a local directory to that remote directory.  `rbd` will error if running
in a non-mapped directory.
