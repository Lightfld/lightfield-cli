# Changelog

## 0.4.0 (2026-04-18)

Full Changelog: [v0.3.1...v0.4.0](https://github.com/Lightfld/lightfield-cli/compare/v0.3.1...v0.4.0)

### Features

* allow `-` as value representing stdin to binary-only file parameters in CLIs ([9835756](https://github.com/Lightfld/lightfield-cli/commit/9835756dfcbb016b0be9fe43c190e5f0292ca6eb))
* **api:** add meetings and files endpoints ([b2e8a34](https://github.com/Lightfld/lightfield-cli/commit/b2e8a348e6e475cf73e7950f166dd4402f1bf5c8))
* **api:** add notes and lists API ([2ff8ff7](https://github.com/Lightfld/lightfield-cli/commit/2ff8ff7bb0ae0bfe4f6b3f594bb8be8f15e5ed5c))
* **api:** manual updates ([c1878b9](https://github.com/Lightfld/lightfield-cli/commit/c1878b9f6f9051581adbd35ac6e3edcc2ec26376))
* **api:** tighten object API field schemas and named OpenAPI models ([a1c681f](https://github.com/Lightfld/lightfield-cli/commit/a1c681fea6221ae9a773b01fed7680e79ef222dc))
* better error message if scheme forgotten in CLI `*_BASE_URL`/`--base-url` ([e46a40d](https://github.com/Lightfld/lightfield-cli/commit/e46a40d5ed5f25f2bed19ab7e2f33c16fdb58209))
* binary-only parameters become CLI flags that take filenames only ([544d7a6](https://github.com/Lightfld/lightfield-cli/commit/544d7a654875a33b0732a9d94ae6d9fc28e6de98))
* **cli:** add `--raw-output`/`-r` option to print raw (non-JSON) strings ([0fc4308](https://github.com/Lightfld/lightfield-cli/commit/0fc43084a18b0b699a8718f197e7a5ddb24c5469))
* **cli:** alias parameters in data with `x-stainless-cli-data-alias` ([a7226f6](https://github.com/Lightfld/lightfield-cli/commit/a7226f6fca64ac1f8628022825483648074e57a7))
* **cli:** send filename and content type when reading input from files ([7af852d](https://github.com/Lightfld/lightfield-cli/commit/7af852d596c6f0d9f4212ba9f36e5df207334a0c))
* set CLI flag constant values automatically where `x-stainless-const` is set ([22817a4](https://github.com/Lightfld/lightfield-cli/commit/22817a48fbf2aac27665061202be7b0b4c88ba5b))


### Bug Fixes

* **api:** simplify object API schemas for SDK generation ([5245cd9](https://github.com/Lightfld/lightfield-cli/commit/5245cd9849be630fc8e4526cfa591e6735f881dc))
* **api:** untyped fields for typescript ([bc4ce3d](https://github.com/Lightfld/lightfield-cli/commit/bc4ce3d2a7fe0edf107df9f46883205be07470d8))
* cli no longer hangs when stdin is attached to a pipe with empty input ([260f676](https://github.com/Lightfld/lightfield-cli/commit/260f676b1d018c4d65973479a833e9fedba157da))
* fall back to main branch if linking fails in CI ([53eec86](https://github.com/Lightfld/lightfield-cli/commit/53eec863051e72027a183992500a6479ebae9125))
* fix for failing to drop invalid module replace in link script ([b85d467](https://github.com/Lightfld/lightfield-cli/commit/b85d467da694c3c18acc0c34dc84f21ccbdb0745))
* fix for off-by-one error in pagination logic ([9f8a73a](https://github.com/Lightfld/lightfield-cli/commit/9f8a73adac5630e540cdc7e31e443d35d01d7cc7))
* fix quoting typo ([a72d9b6](https://github.com/Lightfld/lightfield-cli/commit/a72d9b671765fc5fb6bdc4052116565ef76fb7e8))
* handle empty data set using `--format explore` ([ef00269](https://github.com/Lightfld/lightfield-cli/commit/ef00269145bec49ced9af3d922c1fdf3791cc536))
* use `RawJSON` when iterating items with `--format explore` in the CLI ([6fe1141](https://github.com/Lightfld/lightfield-cli/commit/6fe1141f4d5209506d027cd4081dc9f1d0f33964))


### Chores

* add documentation for ./scripts/link ([1b2125b](https://github.com/Lightfld/lightfield-cli/commit/1b2125be9bfa9981b5db9b4574995ac38b9da4b9))
* **api:** update file endpoint docs ([0b7b187](https://github.com/Lightfld/lightfield-cli/commit/0b7b187984460c010f49973a518a331419a8f9f5))
* **api:** update note documentation ([a0f34d9](https://github.com/Lightfld/lightfield-cli/commit/a0f34d9017f9a4b96204ff6c2e1b9c404c433228))
* **ci:** skip lint on metadata-only changes ([b944dc2](https://github.com/Lightfld/lightfield-cli/commit/b944dc2839a0d6de23057f9834b22a80220d1cfa))
* **ci:** support manually triggering release workflow ([bc5c362](https://github.com/Lightfld/lightfield-cli/commit/bc5c36216e720651b1722fbc2485aa5998dd978b))
* **cli:** additional test cases for `ShowJSONIterator` ([b22216f](https://github.com/Lightfld/lightfield-cli/commit/b22216f808bb32f9c71df848ffada98a9a3415a3))
* **cli:** fall back to JSON when using default "explore" with non-TTY ([fae8c94](https://github.com/Lightfld/lightfield-cli/commit/fae8c94308a212b40b54f1c630267389955a7b6a))
* **cli:** let `--format raw` be used in conjunction with `--transform` ([999ace8](https://github.com/Lightfld/lightfield-cli/commit/999ace810e95bf61825bb2ccaf0a656379c565ad))
* **cli:** switch long lists of positional args over to param structs ([ff4935a](https://github.com/Lightfld/lightfield-cli/commit/ff4935a0512872219a01706add1657ef356619fb))
* **internal:** update gitignore ([5efb55f](https://github.com/Lightfld/lightfield-cli/commit/5efb55f0d10af7b8b48174ffae57792b29491a0b))
* **internal:** update multipart form array serialization ([5cb4b6d](https://github.com/Lightfld/lightfield-cli/commit/5cb4b6dd1c18fc29b0c329b52a36c4fdb8e08750))
* mark all CLI-related tests in Go with `t.Parallel()` ([332e86c](https://github.com/Lightfld/lightfield-cli/commit/332e86ca7fabf264997be8af323a54de882020bc))
* modify CLI tests to inject stdout so mutating `os.Stdout` isn't necessary ([b9a501f](https://github.com/Lightfld/lightfield-cli/commit/b9a501f46c38fad7de446f510dab0edc6dce94cf))
* omit full usage information when missing required CLI parameters ([833bafc](https://github.com/Lightfld/lightfield-cli/commit/833bafc6fa80b6af1ed63148abe7eeb7e892dfcf))
* switch some CLI Go tests from `os.Chdir` to `t.Chdir` ([183e7d1](https://github.com/Lightfld/lightfield-cli/commit/183e7d1a2bc81d34cdd5c028d36fe3995700cc18))
* **tests:** bump steady to v0.19.4 ([1f311ab](https://github.com/Lightfld/lightfield-cli/commit/1f311ab18f2bedc2c9d2dd29e70049ca8afe8bb3))
* **tests:** bump steady to v0.19.5 ([9bdf5c8](https://github.com/Lightfld/lightfield-cli/commit/9bdf5c8a9a1793267ff36f3ff0a4b583c2756c56))
* **tests:** bump steady to v0.19.6 ([5e9695b](https://github.com/Lightfld/lightfield-cli/commit/5e9695ba7f63116fba4f6ae173ab50ccb9c1ec20))
* **tests:** bump steady to v0.19.7 ([d116364](https://github.com/Lightfld/lightfield-cli/commit/d11636429c5b564374864aea4230f10ff6c900a2))
* **tests:** bump steady to v0.20.1 ([a2429d2](https://github.com/Lightfld/lightfield-cli/commit/a2429d227eb6a045c18d4a4c91ef2727d2482aed))
* **tests:** bump steady to v0.20.2 ([73f721d](https://github.com/Lightfld/lightfield-cli/commit/73f721da2171fbb3f2a70254a4cbfc979a98b6cf))
* update SDK settings ([2e7f65e](https://github.com/Lightfld/lightfield-cli/commit/2e7f65ed998de991f6ad95f31ea37901df5e72b6))

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
