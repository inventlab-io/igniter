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
- Templating
- Partitioned values

### Primary and data storage
There are 2 logical storage within Igniter, `Prmary` and `Data` storage. System data for Igniter operations are saved in the Primary storage while user data such as templates and values are stored in the secondary storage. The default secondary storage is set to the same as primary storage unless configured otherwise.

### Path partitioned storage
Different data storage can be used on different paths. If data storage is specified in the `:store` parameter of `GET`, `PUT` API, the specified storage will be used. Otherwise, the default storage will be used. The default storage will be the same primary storage configured during startup.

### Render

Merges the values specified and render the merged values into template and values into a single response.

`POST ` `/render/k/:path` - Renders template of `path` from the default store

`POST ` `/render/:store/k/:path` - Renders template of `path` from specified `store`

The request body of the `render` method in JSON takes several overloaded form. The body specifies the values to be retrieved to be merged and then rendered into the template. Important to note that the conflicting values keys will be overwritten, the precedence in index order i.e the earlier values will overwrite conflicting values key in the later values

---
**Full Form**

This form allows fine grain control on which values is retrieved.
```jsonc
{
    "values": [{
            //optional, overload (string or string array)
            "storeKeys": ["etcd","boltdb"],
            "path": "/myvalue1"
        },
        {
            //define a single key
            "storeKeys": "etcd",
            "path": "/myvalue2"
        },
        {
            //no storeKey set, will fetch from the same store location as template
            "path": "/myvalue3"
        }]
}
```

**example**
```bash
curl --location --request POST 'localhost:8080/render/k/mytemplate' \
--header 'Content-Type: application/json' \
--data-raw '
{
    "values": [{
            "storeKeys":["etcd","boltdb"],
            "path": "/myvalue1"
        },
        {
            "storeKeys":"etcd",
            "path": "/myvalue2"
        },
        {
            "path": "/myvalue3"
        }]'
```

---

**String Array Form**

A convinience form for multiple values from the same storage location as the template. Values will be fetched from the same store location as template.
```jsonc
{
    "values": [ "/myvalue1", "/myvalue2", "/myvalue3" ]
}
```

**example**
```bash
curl --location --request POST 'localhost:8080/render/k/mytemplate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "values": [ "/myvalue1",  "/myvalue2", "/myvalue3" ]
}'
```
---
**Simple String Form**

A convinience form for single value from the same storage location as the template. will be fetched from the same store location as template.
```jsonc
{
    "values": "/myvalue"
}
```

**example**
```bash
curl --location --request POST 'localhost:8080/render/k/mytemplate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "values": "/myvalue"
}'
```

## API
- `/options/store/:store`
- `/options/secrets/:engine`
- `/template/k/:path`
- `/template/:store/k/:path`
- `/values/k/:path`
- `/values/:store/k/:path`
- `/secrets/k/:path` (TODO)
- `/secrets/:engine/k/:path` (TODO)
- `/render/k/:path`
- `/policy/k/:path` (TODO)