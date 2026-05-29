# [1.3.0](https://github.com/BjoernKarma/gitctl/compare/v1.2.0...v1.3.0) (2026-05-29)


### Features

* add-branch-command ([#100](https://github.com/BjoernKarma/gitctl/issues/100)) ([b5661f0](https://github.com/BjoernKarma/gitctl/commit/b5661f03e3e44e3c579a437083340621998047b2))
* add-fetch-command ([#98](https://github.com/BjoernKarma/gitctl/issues/98)) ([ade5624](https://github.com/BjoernKarma/gitctl/commit/ade56242b05c1df71797e8fa18c3155da28e941d))
* add-stash-command ([#99](https://github.com/BjoernKarma/gitctl/issues/99)) ([b0d2263](https://github.com/BjoernKarma/gitctl/commit/b0d22630cef4b956d096bb716cfb44a4ff551466))

# [1.2.0](https://github.com/BjoernKarma/gitctl/compare/v1.1.1...v1.2.0) (2026-05-28)


### Features

* concurrency feature ([#97](https://github.com/BjoernKarma/gitctl/issues/97)) ([b833525](https://github.com/BjoernKarma/gitctl/commit/b8335256afc91093ccce6cfac38c05c0fed3d5e7))
* openspec document existing capabilities ([#96](https://github.com/BjoernKarma/gitctl/issues/96)) ([27ebe2a](https://github.com/BjoernKarma/gitctl/commit/27ebe2a2ee6f0884a1144a32df1384958686d844))

## [1.1.1](https://github.com/BjoernKarma/gitctl/compare/v1.1.0...v1.1.1) (2026-04-28)


### Bug Fixes

* restrict top-level workflow token permissions to read-only ([#89](https://github.com/BjoernKarma/gitctl/issues/89)) ([9f3ff5f](https://github.com/BjoernKarma/gitctl/commit/9f3ff5fc28407bab8f4b655928ea1f91033368a5)), closes [#9](https://github.com/BjoernKarma/gitctl/issues/9) [#18-23](https://github.com/BjoernKarma/gitctl/issues/18-23)

# [1.1.0](https://github.com/BjoernKarma/gitctl/compare/v1.0.0...v1.1.0) (2026-04-13)


### Features

* add a dedicated "Git Command Failures" Section ([#87](https://github.com/BjoernKarma/gitctl/issues/87)) ([b31ce8d](https://github.com/BjoernKarma/gitctl/commit/b31ce8d322e0db3bc054996cf813d43090c9b309))

# 1.0.0 (2026-04-13)


### Bug Fixes

* Adjust test setup for TestGitRepoRunGitPull ([#85](https://github.com/BjoernKarma/gitctl/issues/85)) ([1459c8d](https://github.com/BjoernKarma/gitctl/commit/1459c8dbcb8097483802d169acea34e79f4cedab))
* Change test assertion ([bf023a0](https://github.com/BjoernKarma/gitctl/commit/bf023a0409419afaecb041a475e5eb5c6b5ed20a))
* disable test case ([201e2f7](https://github.com/BjoernKarma/gitctl/commit/201e2f7bd44f6a337d837e8b5524f543b27bea9a))
* Fix asserts for TestGitRepoRunGitPull ([46b6cc2](https://github.com/BjoernKarma/gitctl/commit/46b6cc2b67e476fec3a221a1df1f806c803b4d90))


### Features

* **actions:** Use golangci-lint version v2.1 ([af0f8f0](https://github.com/BjoernKarma/gitctl/commit/af0f8f0e2019474de4c572e2b9f72b057e091a30))
* **actions:** Use stable go-version in workflows and update to golangci-lint v7 ([3f71d6b](https://github.com/BjoernKarma/gitctl/commit/3f71d6b6121a0750b1e8cc45045c76214855eae7))
* Add config with files and flags ([5d2784f](https://github.com/BjoernKarma/gitctl/commit/5d2784fdee57c2168e723bb4932cc08646727e69))
* Add formatted and colored output ([474ab32](https://github.com/BjoernKarma/gitctl/commit/474ab32be64608e0f9e902bdc2a4f2bea6c6d5d0))
* Add lipgloss and convert output as a tree ([f3318a8](https://github.com/BjoernKarma/gitctl/commit/f3318a8e28d12b33be328730f0d222047ba5b954))
* Add output format and colors ([80023cd](https://github.com/BjoernKarma/gitctl/commit/80023cd930872fae073ad6539aee4f8cd5f1242d))
* Extract commands ([5ef84e9](https://github.com/BjoernKarma/gitctl/commit/5ef84e9e07f10afbfba5194fdd295d8f8482f48c))
* Only add valid paths as base dir ([f633031](https://github.com/BjoernKarma/gitctl/commit/f633031e66098a393cbb5cc07b610d9b0c8adee3))
* Remove default codeql.yml ([63c2f99](https://github.com/BjoernKarma/gitctl/commit/63c2f995c8838e40c264cc8389d09358c14d4ea0))
* Remove fixed test resources ([c2589f8](https://github.com/BjoernKarma/gitctl/commit/c2589f856c7522ed461acaf59f6a06f46925ef63))
* Update dependencies ([e606924](https://github.com/BjoernKarma/gitctl/commit/e606924b5eced5c7eeae03e6f33451cd77ae7227))
* Update dependency versions ([6231cf3](https://github.com/BjoernKarma/gitctl/commit/6231cf36c0ee5187f089084ff10b10e4208c414d))
* Update gitignore ([2c32862](https://github.com/BjoernKarma/gitctl/commit/2c32862dbf0537008d93bf0fb626604fa495773e))
* Update README.md ([679befa](https://github.com/BjoernKarma/gitctl/commit/679befa4606b3903c966f680d5fe0dc11a0bb3e1))
* Update to go 1.23.7 ([b0034d8](https://github.com/BjoernKarma/gitctl/commit/b0034d835c462459d4e419665a3dcb7a4fcedd3e))
* Update to go 1.24.1 ([974f4c0](https://github.com/BjoernKarma/gitctl/commit/974f4c0f2f1225f2e12c7bb87e5e43fe0dfafaff))
