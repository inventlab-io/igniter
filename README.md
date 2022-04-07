# igniter

Igniter is a tool for managing app configuration to support complex deployment reqiurements.

## Features

- Path partitioned storage
- Storage types
    - etcd
    - BoltDb (TODO)
    - Filesystem (TODO)
    - Github (TODO)
- Support for dynamic secrets engines (TODO)
    - Vault (TODO)
    - Generic (curl) (TODO)

### Path partitioned storage
Different storage can be used on different paths. If storage is specified in the `:store` parameter of `GET`, `PUT` API, the specified storage will be used. 

## API
- `/options/store/k/:store`
- `/template/k/:path`
- `/template/:store/k/:path`
- `/data/k/:path`
- `/data/:store/k/:path`
- `/render/k/:path`
- `/policy/k/:path`
- `/secret/:engine/options/default`
- `/secret/:engine/options/k/:path`
- `/options/[template|data]/k/:path`