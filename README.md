# File Upload Streaming Demo

## Purpose

The purpose of this repo is to hold the example code for an backend server listening for a file upload and saving the file to disk

It holds two implementations

### NodeJS with Fastify

The NodeJS implementation uses the Fastify framework to handle the file upload.

The Fastify framework does the following:

1. Streams the file directly to the final destination.

### Go with Gin

The Go implementation uses the Gin framework to handle the file upload.

The Gin framework does the following:

1. Streams the file to a tmp file on disk, located in $TMPDIR.
1. Then it copies the file to the final destination.

It seems that there is no way stream the file directly to the final destination.

## Running the code

Start the servers and post a file to the servers, and compare the results.

## Results from testing

The tests are executed with a 10GB file, generated with the following command:

```bash
mkfile -n 10g big-image.img
```

### Saving to disk

- Fastify: ~ 15 seconds
- Gin: ~ 22 seconds

### Saving to an NFS mount

- Fastify: ~ 30 seconds
- Gin: ~ 50 seconds

## Oberservations

The Go implementation becomes slower, because it will wait to the whole file to be downloaded to the tmp directory, and then the NFS Copy will be started.
While the NodeJS implementation will start the copy to the NFS mount as soon as the first bytes are received.
