# igniter

Igniter is a tool for managing app configuration to support complex deployment reqiurements.

## Features

- Support for multi configuration storage
    - ETCD
    - BoltDb (TODO)
    - Filesystem (TODO)
    - Github (TODO)
- Support for dynamic secrets engines (TODO)
    - Vault
    - Generic (curl) 

## API

- `/template/:repo/k/:path`
- `/data/:repo/k/:path`
- `/render/k/:path`
- `/meta/k/:path`
- `/policy/k/:path`
- `/secret/:engine/setting/default`
- `/secret/:engine/setting/k/:path`