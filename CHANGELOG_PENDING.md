### Improvements

- [sdk/dotnet] - Add `PluginDownloadURL` as a resource option. When provided by
  the schema, `PluginDownloadURL` will be baked into `new Resource` and `Invoke`
  requests in generated SDKs. 
  [#8739](https://github.com/pulumi/pulumi/pull/8739)

- [sdk] - Allow property paths to accept `[*]` as sugar for `["*"]`.
  [#8743](https://github.com/pulumi/pulumi/pull/8743)

### Bug Fixes

