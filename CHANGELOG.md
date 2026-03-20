# Changelog

## 0.3.1 (2026-03-20)

Full Changelog: [v0.3.0...v0.3.1](https://github.com/Lightfld/lightfield-cli/compare/v0.3.0...v0.3.1)

### Bug Fixes

* **api:** resolve system and custom attribute collision error ([a698bab](https://github.com/Lightfld/lightfield-cli/commit/a698babd7cf30f91aaa1e99412a5991fdc815e72))


### Chores

* remove custom code ([fda3047](https://github.com/Lightfld/lightfield-cli/commit/fda3047360436f3cf7f0c7c146564351f652aabc))

## 0.3.0 (2026-03-20)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/Lightfld/lightfield-cli/compare/v0.2.0...v0.3.0)

### Features

* **api:** add field/relationship param descriptions ([b1e7d8a](https://github.com/Lightfld/lightfield-cli/commit/b1e7d8ac50a1ad2fa7147487c3d3b69981ecea9f))
* **api:** add Go SDK and CLI targets to Stainless config ([267d53e](https://github.com/Lightfld/lightfield-cli/commit/267d53e120250c4a933a3a994f67c51a26534ac2))
* **api:** manual updates ([7144873](https://github.com/Lightfld/lightfield-cli/commit/7144873686ec8ef6cac117abc454fd10ba5c9a14))
* **api:** manual updates ([f61090e](https://github.com/Lightfld/lightfield-cli/commit/f61090e028cef0b57a4e685f0ba456de7a03e156))
* **api:** manual updates ([ed3d3b5](https://github.com/Lightfld/lightfield-cli/commit/ed3d3b59f7e44a50f322a01d96243d041503aa44))
* **api:** member and workflow run ([b26aeb9](https://github.com/Lightfld/lightfield-cli/commit/b26aeb9ab557131973c8a98c9149fe4d69cf64e6))
* **api:** shorten system attribute prefixes, add attribute definition endpoint ([f77d8ca](https://github.com/Lightfld/lightfield-cli/commit/f77d8ca0be2b663ce34ac38dee9632ac50d14004))
* **api:** update field descriptions, plural to singular keys, change READONLY_MARKDOWN type to MARKDOWN ([9370dcf](https://github.com/Lightfld/lightfield-cli/commit/9370dcf898cc92e8e20626afb71b1150adaa958b))


### Bug Fixes

* avoid reading from stdin unless request body is form encoded or json ([f258e74](https://github.com/Lightfld/lightfield-cli/commit/f258e745e049446373abe742edd8a2831d8ca86a))
* better support passing client args in any position ([2068bc5](https://github.com/Lightfld/lightfield-cli/commit/2068bc5206ecb21dec0213401d5b0f438f17bc15))
* improve linking behavior when developing on a branch not in the Go SDK ([62f426e](https://github.com/Lightfld/lightfield-cli/commit/62f426e8912391961b4a50266443c0b8499f991a))
* improved workflow for developing on branches ([80bf6b1](https://github.com/Lightfld/lightfield-cli/commit/80bf6b15a194010d71cfb489bda9b19b36c184d0))
* no longer require an API key when building on production repos ([83a0f9a](https://github.com/Lightfld/lightfield-cli/commit/83a0f9a35ab15c94b1fd533750528d33b63a96ad))


### Chores

* configure new SDK language ([3eaf042](https://github.com/Lightfld/lightfield-cli/commit/3eaf042a18ccde388f51a27b6b19bcf30e635e4e))
* configure new SDK language ([1cfd876](https://github.com/Lightfld/lightfield-cli/commit/1cfd876291292d28998136ce34cb61745b06feac))
* **internal:** codegen related update ([e4af2cf](https://github.com/Lightfld/lightfield-cli/commit/e4af2cf9dc07e0c2a0123faf68ab1fab88a681ad))
* **internal:** tweak CI branches ([94f8209](https://github.com/Lightfld/lightfield-cli/commit/94f8209b38d1e0f4063f52983880c2d1e2e57020))
* sync repo ([c0e4846](https://github.com/Lightfld/lightfield-cli/commit/c0e4846bca28195b945b7e3ce6a72fe3883ea8a9))
* update SDK settings ([2908846](https://github.com/Lightfld/lightfield-cli/commit/2908846990b130c73cf3fa4c1b3b141d7559d8d1))


### Refactors

* **tests:** switch from prism to steady ([5477f71](https://github.com/Lightfld/lightfield-cli/commit/5477f719e6a3401829b9bc960336c988d435d01d))

## 0.2.0 (2026-03-20)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/Lightfld/lightfield-cli/compare/v0.1.0...v0.2.0)

### Features

* **api:** add field/relationship param descriptions ([b1e7d8a](https://github.com/Lightfld/lightfield-cli/commit/b1e7d8ac50a1ad2fa7147487c3d3b69981ecea9f))
* **api:** manual updates ([7144873](https://github.com/Lightfld/lightfield-cli/commit/7144873686ec8ef6cac117abc454fd10ba5c9a14))
* **api:** manual updates ([f61090e](https://github.com/Lightfld/lightfield-cli/commit/f61090e028cef0b57a4e685f0ba456de7a03e156))
* **api:** member and workflow run ([b26aeb9](https://github.com/Lightfld/lightfield-cli/commit/b26aeb9ab557131973c8a98c9149fe4d69cf64e6))
* **api:** update field descriptions, plural to singular keys, change READONLY_MARKDOWN type to MARKDOWN ([9370dcf](https://github.com/Lightfld/lightfield-cli/commit/9370dcf898cc92e8e20626afb71b1150adaa958b))


### Bug Fixes

* avoid reading from stdin unless request body is form encoded or json ([f258e74](https://github.com/Lightfld/lightfield-cli/commit/f258e745e049446373abe742edd8a2831d8ca86a))
* better support passing client args in any position ([2068bc5](https://github.com/Lightfld/lightfield-cli/commit/2068bc5206ecb21dec0213401d5b0f438f17bc15))
* improve linking behavior when developing on a branch not in the Go SDK ([62f426e](https://github.com/Lightfld/lightfield-cli/commit/62f426e8912391961b4a50266443c0b8499f991a))
* improved workflow for developing on branches ([80bf6b1](https://github.com/Lightfld/lightfield-cli/commit/80bf6b15a194010d71cfb489bda9b19b36c184d0))
* no longer require an API key when building on production repos ([83a0f9a](https://github.com/Lightfld/lightfield-cli/commit/83a0f9a35ab15c94b1fd533750528d33b63a96ad))


### Chores

* **internal:** codegen related update ([e4af2cf](https://github.com/Lightfld/lightfield-cli/commit/e4af2cf9dc07e0c2a0123faf68ab1fab88a681ad))
* **internal:** tweak CI branches ([94f8209](https://github.com/Lightfld/lightfield-cli/commit/94f8209b38d1e0f4063f52983880c2d1e2e57020))


### Refactors

* **tests:** switch from prism to steady ([5477f71](https://github.com/Lightfld/lightfield-cli/commit/5477f719e6a3401829b9bc960336c988d435d01d))

## 0.1.0 (2026-03-17)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/Lightfld/lightfield-cli/compare/v0.0.1...v0.1.0)

### Features

* **api:** add Go SDK and CLI targets to Stainless config ([267d53e](https://github.com/Lightfld/lightfield-cli/commit/267d53e120250c4a933a3a994f67c51a26534ac2))
* **api:** manual updates ([ed3d3b5](https://github.com/Lightfld/lightfield-cli/commit/ed3d3b59f7e44a50f322a01d96243d041503aa44))
* **api:** shorten system attribute prefixes, add attribute definition endpoint ([f77d8ca](https://github.com/Lightfld/lightfield-cli/commit/f77d8ca0be2b663ce34ac38dee9632ac50d14004))


### Chores

* configure new SDK language ([3eaf042](https://github.com/Lightfld/lightfield-cli/commit/3eaf042a18ccde388f51a27b6b19bcf30e635e4e))
* configure new SDK language ([1cfd876](https://github.com/Lightfld/lightfield-cli/commit/1cfd876291292d28998136ce34cb61745b06feac))
* sync repo ([c0e4846](https://github.com/Lightfld/lightfield-cli/commit/c0e4846bca28195b945b7e3ce6a72fe3883ea8a9))
* update SDK settings ([2908846](https://github.com/Lightfld/lightfield-cli/commit/2908846990b130c73cf3fa4c1b3b141d7559d8d1))
