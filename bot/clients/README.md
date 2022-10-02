# Clients

These are clients generated from one or more open-api specs and the [openapi-generator](https://openapi-generator.tech)

## Overseerr

| Client for | [overseerr](https://github.com/sct/overseerr) |
| OpenAPI Spec | [https://api-docs.overseerr.dev/overseerr-api.yml] |
| Command | `openapi-generator-cli generate -o overseerr -i https://api-docs.overseerr.dev/overseerr-api.yml -g go` |

This is on hold because overseerr is using an invalid spec

```log
Exception in thread "main" org.openapitools.codegen.SpecValidationException: There were issues with the specification. T
he option can be disabled via validateSpec (Maven/Gradle) or --skip-validate-spec (CLI).
 | Error count: 1, Warning count: 2
Errors:
        -attribute paths.'/user'(get).responses.200.content.'application/json'.schema.items is not of type `object`
Warnings:
        -attribute paths.'/user'(get).responses.200.content.'application/json'.schema.items is not of type `object`

        at org.openapitools.codegen.config.CodegenConfigurator.toContext(CodegenConfigurator.java:604)
        at org.openapitools.codegen.config.CodegenConfigurator.toClientOptInput(CodegenConfigurator.java:631)
        at org.openapitools.codegen.cmd.Generate.execute(Generate.java:457)
        at org.openapitools.codegen.cmd.OpenApiGeneratorCommand.run(OpenApiGeneratorCommand.java:32)
        at org.openapitools.codegen.OpenAPIGenerator.main(OpenAPIGenerator.java:66)
```
