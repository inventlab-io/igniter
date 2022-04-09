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
- Templated config (TODO)
- Partitioned values (TODO)

### Primary and data storage
There are 2 logical storage within Igniter, `Prmary` and `Data` storage. System data for Igniter operations are saved in the Primary storage while user data such as templates and values are stored in the secondary storage. The default secondary storage is set to the same as primary storage unless configured otherwise.

### Path partitioned storage
Different data storage can be used on different paths. If data storage is specified in the `:store` parameter of `GET`, `PUT` API, the specified storage will be used. Otherwise, the default storage will be used. The default storage will be the same primary storage configured during startup.

## API
- `/options/store/k/:store`
- `/template/k/:path`
- `/template/:store/k/:path`
- `/values/k/:path`
- `/values/:store/k/:path`
- `/render/k/:path`
- `/policy/k/:path` (TODO)
- `/secret/engine/k/:engine` (TODO)
- `/secret/:engine/k/:path` (TODO)