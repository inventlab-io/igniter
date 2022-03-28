# igniter

Igniter is a tool for managing app configuration to support complex deployment reqiurements.

## Features

- Support for multi configuration storage
    - etcd
    - BoltDb (TODO)
    - Filesystem (TODO)
    - Github (TODO)
- Support for dynamic secrets engines (TODO)
    - Vault
    - Generic (curl) 

## API

- `/template/:store/k/:path`
- `/data/:store/k/:path`
- `/render/k/:path`
- `/meta/k/:path`
- `/policy/k/:path`
- `/secret/:engine/setting/default`
- `/secret/:engine/setting/k/:path`