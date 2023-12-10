# Clients

These are clients generated from one or more open-api specs and the [openapi-generator](https://openapi-generator.tech)

## Overseerr

| Client for | [overseerr](https://github.com/sct/overseerr) |
| OpenAPI Spec | [https://api-docs.overseerr.dev/overseerr-api.yml] |
| Command | `openapi-generator-cli generate -o overseerr -i https://api-docs.overseerr.dev/overseerr-api.yml -g go --skip-validate-spec --additional-properties="withGoMod=false,packageName=overseerr_go,structPrefix=true"` |

This was on hold because overseerr is using an "invalid" spec.

```log
Exception in thread "main" org.openapitools.codegen.SpecValidationException: There were issues with the specification. The option can be disabled via validateSpec (Maven/Gradle) or --skip-validate-spec (CLI).
 | Error count: 2, Warning count: 2
Errors:
        -attribute paths.'/user'(get).responses.200.content.'application/json'.schema.items is not of type `object`
        -paths.'/settings/discover/{sliderId}'. Declared path parameter sliderId needs to be defined as a path parameter in path or operation level
Warnings:
        -attribute paths.'/user'(get).responses.200.content.'application/json'.schema.items is not of type `object`
        -paths.'/settings/discover/{sliderId}'. Declared path parameter sliderId needs to be defined as a path parameter in path or operation level

        at org.openapitools.codegen.config.CodegenConfigurator.toContext(CodegenConfigurator.java:684)
        at org.openapitools.codegen.config.CodegenConfigurator.toClientOptInput(CodegenConfigurator.java:711)
        at org.openapitools.codegen.cmd.Generate.execute(Generate.java:511)
        at org.openapitools.codegen.cmd.OpenApiGeneratorCommand.run(OpenApiGeneratorCommand.java:32)
        at org.openapitools.codegen.OpenAPIGenerator.main(OpenAPIGenerator.java:66)
```

The first issue is addressed in [https://github.com/OpenAPITools/openapi-generator/issues/14317]

For now I've side-stepped this by adding the `--skip-validate-spec` flag. I've also included a yaml of the spec used for generation.
