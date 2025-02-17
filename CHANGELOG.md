# [1.24.0](https://github.com/weka/weka-k8s-api/compare/v1.23.0...v1.24.0) (2025-02-17)


### Bug Fixes

* add cluster overrides ([29f16f1](https://github.com/weka/weka-k8s-api/commit/29f16f1770256950651c7ffe148f21ab43668f82))
* add deepcopy funcs ([9790640](https://github.com/weka/weka-k8s-api/commit/9790640bea57a7d8a0e0bad0f72045127c7624d5))
* add descriptions ([199be91](https://github.com/weka/weka-k8s-api/commit/199be916be053970ae55e20b53744309992648d3))
* add machineIdentifier to resources ([338834e](https://github.com/weka/weka-k8s-api/commit/338834e6a6dabe455a4fd24f7901936395d8f521))
* allow alternative drivers-loader for cluster ([2a39277](https://github.com/weka/weka-k8s-api/commit/2a3927758fb08843c7e5c2a40616c3bfeeb4c9fc))
* allow for pre-start script on container, for custom setup before drivers build ([2372fd4](https://github.com/weka/weka-k8s-api/commit/2372fd49aad82ee5ca3b938cfc5a7d3cc1313863))
* container exposePorts, to be exposed as pod ports ([b7eb806](https://github.com/weka/weka-k8s-api/commit/b7eb806dc20d119d1211ff2b9b21110fe77c8660))
* delete legacy driveOptions field on wekacontainer ([f27cae3](https://github.com/weka/weka-k8s-api/commit/f27cae33071757ced752adc41e34aec59ffbad25))
* exposedPorts instead of exposePort as more advanced struct ([ca33edb](https://github.com/weka/weka-k8s-api/commit/ca33edb5087e476198d22106a53795279d9514b5))
* fix case of force resign payload ([57edcc2](https://github.com/weka/weka-k8s-api/commit/57edcc226f130abea1fdd93b6a152b157071e48b))
* limit SignDrivesPayload to certain type values ([88f9f0a](https://github.com/weka/weka-k8s-api/commit/88f9f0a0ff95f8e01e817b19fc20d75f4d145146))
* make manual op / policy image optional ([be660f2](https://github.com/weka/weka-k8s-api/commit/be660f2ec6dc877e08fc883c526350c70782764b))
* management ips plural variant. to be used for applicable network recognition ([2280d28](https://github.com/weka/weka-k8s-api/commit/2280d286621cd97f994832405ecab7c86df699b2))
* metrics string accessors ([d64e764](https://github.com/weka/weka-k8s-api/commit/d64e764ebbc1ca9a13955090ab50e63e6158f5ff))
* metrics via endpoint ([621c0a7](https://github.com/weka/weka-k8s-api/commit/621c0a7a441fe62ddd7ad1046ba3ef91011337f0))
* move disregardRedundancy to overrides ([2fdb01f](https://github.com/weka/weka-k8s-api/commit/2fdb01fead43fd5f3397b81833979dad7a050d9e))
* move DriversLoaderImage to overrides ([e2e81e4](https://github.com/weka/weka-k8s-api/commit/e2e81e46c92880e810247ec33b31a5f472bf03c2))
* move ForceAio to overrides ([46dcc1e](https://github.com/weka/weka-k8s-api/commit/46dcc1e0d8860721185e93524376feae58a1d216))
* replace allow options with generic optinos ([c346ad4](https://github.com/weka/weka-k8s-api/commit/c346ad4593d4c5c134b21700b0e860416725edc5))
* skipCleanupPersistentDir unsafe comment/explanation ([b2e7801](https://github.com/weka/weka-k8s-api/commit/b2e7801ccb3b4b46947329776cf01c54bd5efedc))
* support for envoy additional memory ([56be48b](https://github.com/weka/weka-k8s-api/commit/56be48be6f4999757dd63ad0d84e7e4d2f17ec70))
* support new sign utility flags ([0e5f458](https://github.com/weka/weka-k8s-api/commit/0e5f4580d18dd7033e5b9318191e501c258666cc))
* use duration type for wekapolicy interval ([d908fb4](https://github.com/weka/weka-k8s-api/commit/d908fb43e0760c1915b626ddd1f1d65b7d56ef93))


### Features

* add activeMounts to the wekacontainer stats ([013ef7a](https://github.com/weka/weka-k8s-api/commit/013ef7ae47255316d69aad18e27bc06fe4c6ce24))
* add allowForceUpgrade field in wekacontainer spec ([2601085](https://github.com/weka/weka-k8s-api/commit/26010857d77277b1201d2ea8c25db4e3aca1fb57))
* add allowS3ClusterDestroy option to WekaCluster spec ([7db220a](https://github.com/weka/weka-k8s-api/commit/7db220a5d97ae123f275b1c4151fe500724d5727))
* add destroying container state ([333bfdd](https://github.com/weka/weka-k8s-api/commit/333bfdda67112e653c3f2d69796104d87fb83aa7))
* add devicesSubnet field in network configuration ([9efddca](https://github.com/weka/weka-k8s-api/commit/9efddcabb943869d4a4ebe5fc0efd419e5285250))
* add explicit gateway param for clients and cluster ([76fbcb9](https://github.com/weka/weka-k8s-api/commit/76fbcb9ba768dc2959ee91ee3e8bdba7bb3256bd))
* add GetHostIps, print management ips list ([4175657](https://github.com/weka/weka-k8s-api/commit/41756576cc6329cefce2aa120ec0309f6f6b6a93))
* add leadershipRaftSize and bucketRaftSize ([44d45b1](https://github.com/weka/weka-k8s-api/commit/44d45b119cdd27de4af27a0acfe6b34ffd04e276))
* add managementIPsStr column for ips displaying ([e3b142d](https://github.com/weka/weka-k8s-api/commit/e3b142dc570fb0b634166d91d00e9b8ff5a0d23d))
* add net devices to support ensure nics ([e004538](https://github.com/weka/weka-k8s-api/commit/e0045383e1420f5ef854d58e5d5100a404e74071))
* add NormalizeTolerations function ([69da029](https://github.com/weka/weka-k8s-api/commit/69da0293afd9b815084dfe31975605e19984eb83))
* add postFormClusterScript override, a script to run before startio ([e32620a](https://github.com/weka/weka-k8s-api/commit/e32620af35944334e70080b93492f1490e7671f9))
* add StartingIO and WaitForDrives cluster statuses ([ebe10f0](https://github.com/weka/weka-k8s-api/commit/ebe10f0759899ced45e699d7acf05708d2be04b7))
* add startIoConditions for cluster, rm message from container ([354ba49](https://github.com/weka/weka-k8s-api/commit/354ba49bb04d4ae71c76c3d6fc50116ab0ebf540))
* add type for wekacontainer instructions ([4a2c3af](https://github.com/weka/weka-k8s-api/commit/4a2c3af46972731b85bdbcbb25ff102a07bb30d6))
* add unblock-drives manual operation action ([43bef8c](https://github.com/weka/weka-k8s-api/commit/43bef8cda8533af6fa45568de0b1401ec30e1f5b))
* add wekacontainer deleting state ([f67c7b4](https://github.com/weka/weka-k8s-api/commit/f67c7b4099c0c6feaf83fb8eece573671d9d9a24))
* allowForceUpgrade -> allowHotUpgrade in wekacontainer spec ([cdf601a](https://github.com/weka/weka-k8s-api/commit/cdf601a2ce054fd6bfc0b9a986ca944aaa73c734))
* bring back tombstones removal ([21f35e6](https://github.com/weka/weka-k8s-api/commit/21f35e6c7c1b48a94be82ff34ef4814f94e1114d))
* clients counters on printer and status ([7f301d8](https://github.com/weka/weka-k8s-api/commit/7f301d82053a204e0bd35a2ab42c21fd605de684))
* force upgrade flow ([b3c6507](https://github.com/weka/weka-k8s-api/commit/b3c6507f6799b56aba91ce8c7c1650bc147b62f1))
* generic CR timestamps for actions throttling ([46842dc](https://github.com/weka/weka-k8s-api/commit/46842dcb16d486b92b953732a005060bcfa9186e))
* machine identifier ref for clients ([23c6047](https://github.com/weka/weka-k8s-api/commit/23c6047e682bcdc8384ff38ae6b053013817d645))
* make all-at-once default upgrade policy for client ([50d40f2](https://github.com/weka/weka-k8s-api/commit/50d40f26e932900706e7b8f752656d96580a65a5))
* move container overrides to spec, remove cluster overrideGracefulDestroyDuration ([8f936eb](https://github.com/weka/weka-k8s-api/commit/8f936ebe2e5a7b6babc66a6ce4762d8f9d355bc0))
* remote viewer tls config with trusted by default ([f0f6f6e](https://github.com/weka/weka-k8s-api/commit/f0f6f6e7720b85c00862c953b8c8231d6b023e0d))
* remove tombstones ([3e94171](https://github.com/weka/weka-k8s-api/commit/3e9417120000cbbeff46a212b477dfbf8dd2d54b))
* support multiple network devices ([7a2c993](https://github.com/weka/weka-k8s-api/commit/7a2c99338f031dc9963d0d8ba5407cbe05b6dbb6))
* update failure domain config for weka containers, allow compositeLabels ([03237b9](https://github.com/weka/weka-k8s-api/commit/03237b9acacb9ea7e0810e5c7a51ddf97da4f261))


### Reverts

* Revert "feat: remove tombstones" ([69846b8](https://github.com/weka/weka-k8s-api/commit/69846b896933c45b34106fdb4c450e65112aaac4))

# [1.23.0](https://github.com/weka/weka-k8s-api/compare/v1.22.0...v1.23.0) (2024-11-26)


### Bug Fixes

* add loader mode to IsAdhocOpContainer ([b78e5a5](https://github.com/weka/weka-k8s-api/commit/b78e5a5cc54ddf21f6fd45b375c3bd74ba52e182))
* change nfsGateway to nfs ([e5e5edc](https://github.com/weka/weka-k8s-api/commit/e5e5edce3c184249d9e46ac9dca1f0c0c4a2558f))


### Features

* manual operation of openning remote trace session ([f5f7c1f](https://github.com/weka/weka-k8s-api/commit/f5f7c1fca6eee781c1513bd1894104cfabdc4225))
* metrics via crd ([d9302bd](https://github.com/weka/weka-k8s-api/commit/d9302bd4bb6ff92def29057db043960b70c2fb3f))
* use crds as metrics storage ([a3ca575](https://github.com/weka/weka-k8s-api/commit/a3ca575db721bdb9c81bfdcdcf5283abb4b4c7bf))

# [1.22.0](https://github.com/weka/weka-k8s-api/compare/v1.21.0...v1.22.0) (2024-11-25)


### Bug Fixes

* add loader mode to IsAdhocOpContainer ([0df651e](https://github.com/weka/weka-k8s-api/commit/0df651e5f305d4766515ea3c7c09d3375e3b77b3))
* add loader mode to IsAdhocOpContainer ([#43](https://github.com/weka/weka-k8s-api/issues/43)) ([f33a857](https://github.com/weka/weka-k8s-api/commit/f33a857ed6ca8d8a859972d9c9896002ffeb730e))


### Features

* add ensure-nics manual operation ([c6fe92c](https://github.com/weka/weka-k8s-api/commit/c6fe92c407ca19b9a645add41de6ecc06c1db1f0))
* add ensure-nics manual operation ([#46](https://github.com/weka/weka-k8s-api/issues/46)) ([312c122](https://github.com/weka/weka-k8s-api/commit/312c122a82387b45fd74d943987b7486325ed7ea))
* add message field to weka container status ([3b015b0](https://github.com/weka/weka-k8s-api/commit/3b015b0d2ab08bb695bf5c0874d40cc3690a1401))

# [1.22.0](https://github.com/weka/weka-k8s-api/compare/v1.21.0...v1.22.0) (2024-11-25)


### Bug Fixes

* add loader mode to IsAdhocOpContainer ([0df651e](https://github.com/weka/weka-k8s-api/commit/0df651e5f305d4766515ea3c7c09d3375e3b77b3))
* add loader mode to IsAdhocOpContainer ([#43](https://github.com/weka/weka-k8s-api/issues/43)) ([f33a857](https://github.com/weka/weka-k8s-api/commit/f33a857ed6ca8d8a859972d9c9896002ffeb730e))


### Features

* add ensure-nics manual operation ([c6fe92c](https://github.com/weka/weka-k8s-api/commit/c6fe92c407ca19b9a645add41de6ecc06c1db1f0))
* add ensure-nics manual operation ([#46](https://github.com/weka/weka-k8s-api/issues/46)) ([312c122](https://github.com/weka/weka-k8s-api/commit/312c122a82387b45fd74d943987b7486325ed7ea))
* add message field to weka container status ([3b015b0](https://github.com/weka/weka-k8s-api/commit/3b015b0d2ab08bb695bf5c0874d40cc3690a1401))

# [1.21.0](https://github.com/weka/weka-k8s-api/compare/v1.20.0...v1.21.0) (2024-11-11)


### Features

* add stripe size and data protection level configs ([d9021af](https://github.com/weka/weka-k8s-api/commit/d9021afaf86f45e3fbb15ac78c741c4d0f4c7770))

# [1.20.0](https://github.com/weka/weka-k8s-api/compare/v1.19.0...v1.20.0) (2024-10-30)


### Features

* mark fields on container with p0 ([d416ed8](https://github.com/weka/weka-k8s-api/commit/d416ed85474587995f659bc97f3ffd70dc1ec8d1))

# [1.19.0](https://github.com/weka/weka-k8s-api/compare/v1.18.0...v1.19.0) (2024-10-24)


### Features

* add hugepages offset setting by container role to weka config ([0d7b64f](https://github.com/weka/weka-k8s-api/commit/0d7b64fc8a9e88d6211f4fc30a858e62d6991f48))

# [1.18.0](https://github.com/weka/weka-k8s-api/compare/v1.17.0...v1.18.0) (2024-10-22)


### Features

* add tolerations to tombstone spec ([94a2253](https://github.com/weka/weka-k8s-api/commit/94a22538d55c8326c9e9e052dec980f8a41f4bcf))

# [1.17.0](https://github.com/weka/weka-k8s-api/compare/v1.16.0...v1.17.0) (2024-10-22)


### Features

* add failureDomain to container allocations ([76d4527](https://github.com/weka/weka-k8s-api/commit/76d4527162697057058e49bb3487b88f522645f0))

# [1.16.0](https://github.com/weka/weka-k8s-api/compare/v1.15.0...v1.16.0) (2024-10-22)


### Features

* add forceAio setting for weka cluster (disabled by default) ([0bf9c4e](https://github.com/weka/weka-k8s-api/commit/0bf9c4ea94167c31f76b7407d2e9de31c9901f05))


### Reverts

* support iommuEnabled option for weka cluster ([9dcec57](https://github.com/weka/weka-k8s-api/commit/9dcec57f619bf5bbe863ce99989768e0bcb59c9b))

# [1.15.0](https://github.com/weka/weka-k8s-api/compare/v1.14.0...v1.15.0) (2024-10-21)


### Features

* support iommuEnabled option for weka cluster ([db01869](https://github.com/weka/weka-k8s-api/commit/db0186964838baa57481bf01268f59ebfb54707a))

# [1.14.0](https://github.com/weka/weka-k8s-api/compare/v1.13.0...v1.14.0) (2024-10-21)


### Features

* add PCI devices filtering option for signDrivesPayload ([4097a5f](https://github.com/weka/weka-k8s-api/commit/4097a5ffe86b1fc10c1782e10882fbd9b01f227b))

# [1.13.0](https://github.com/weka/weka-k8s-api/compare/v1.12.0...v1.13.0) (2024-10-21)


### Features

* support hotSpare setting for weka cluster (default 0) ([e4fac51](https://github.com/weka/weka-k8s-api/commit/e4fac512ee9ba66948801cd6ce567eb81a279c2e))

# [1.12.0](https://github.com/weka/weka-k8s-api/compare/v1.11.0...v1.12.0) (2024-10-21)


### Features

* add condition for JoinIpsSet ([d34d587](https://github.com/weka/weka-k8s-api/commit/d34d587d1aaf7bbcf7520d7a02c666d688cab5d8))

# [1.11.0](https://github.com/weka/weka-k8s-api/compare/v1.10.0...v1.11.0) (2024-10-20)


### Features

* add failure domain label and pod config (scheduling) for cluster ([70e01bc](https://github.com/weka/weka-k8s-api/commit/70e01bcd9b529f12646991aa4b91b22691577378))

# [1.10.0](https://github.com/weka/weka-k8s-api/compare/v1.9.0...v1.10.0) (2024-10-09)


### Features

* add labels on WekaContainerDetails ([8cc5340](https://github.com/weka/weka-k8s-api/commit/8cc5340d6f8d19f57a53e33344c17657a5ec1233))

# [1.9.0](https://github.com/weka/weka-k8s-api/compare/v1.8.0...v1.9.0) (2024-10-09)


### Features

* add GetForMode func on AdditionalMemory ([c049681](https://github.com/weka/weka-k8s-api/commit/c04968124a61296d51600ff080f2037ade720767))

# [1.8.0](https://github.com/weka/weka-k8s-api/compare/v1.7.0...v1.8.0) (2024-10-09)


### Features

* add portRange to wekaclient / wekacontainer spec ([bc5f508](https://github.com/weka/weka-k8s-api/commit/bc5f508fd06b54233863a2f58633dc21073d97a4))

# [1.7.0](https://github.com/weka/weka-k8s-api/compare/v1.6.0...v1.7.0) (2024-10-08)


### Features

* add gracefulDestroyTimeout field on WekaCluster ([69bb1f5](https://github.com/weka/weka-k8s-api/commit/69bb1f5408ff222c8b94d03acabeb4e415ceb0a6))

# [1.6.0](https://github.com/weka/weka-k8s-api/compare/v1.5.2...v1.6.0) (2024-10-08)


### Features

* add min length validation for tombstone CrId ([8d01ef9](https://github.com/weka/weka-k8s-api/commit/8d01ef99b1aba50fc8d0d98788704b0ee9125e77))

## [1.5.2](https://github.com/weka/weka-k8s-api/compare/v1.5.1...v1.5.2) (2024-10-07)


### Bug Fixes

* envoy weka container mode ([3616870](https://github.com/weka/weka-k8s-api/commit/3616870740082d9631985f5e663994cbe980ff05))

## [1.5.1](https://github.com/weka/weka-k8s-api/compare/v1.5.0...v1.5.1) (2024-10-07)


### Bug Fixes

* IsWekaContainer condition - include dist containers ([2bce35c](https://github.com/weka/weka-k8s-api/commit/2bce35cbc9f6c5e30e93f7c03b7b2f7480305d1b))

# [1.5.0](https://github.com/weka/weka-k8s-api/compare/v1.4.0...v1.5.0) (2024-10-03)


### Features

* support tolerations for WekaPolicy ([781457f](https://github.com/weka/weka-k8s-api/commit/781457f880bbc4e447aa9b267459001a8f9582c7))

# [1.4.0](https://github.com/weka/weka-k8s-api/compare/v1.3.0...v1.4.0) (2024-10-01)


### Features

* ability to specify custom image for drivers loader ([e09bdcb](https://github.com/weka/weka-k8s-api/commit/e09bdcb42ca0d3bf6dfdc8184e9f093d9cde778d))

# [1.3.0](https://github.com/weka/weka-k8s-api/compare/v1.2.2...v1.3.0) (2024-10-01)


### Features

* wekahome config api improvements ([f92633f](https://github.com/weka/weka-k8s-api/commit/f92633f10ace4ceba01d2346dfee84ec37459bed))

## [1.2.2](https://github.com/weka/weka-k8s-api/compare/v1.2.1...v1.2.2) (2024-09-30)


### Bug Fixes

* upgradePolicyType on container, to propagate desired behaviour ([9b60151](https://github.com/weka/weka-k8s-api/commit/9b601514f83115226f0dcf982a5653ba687ae97d))
* upgradePolicyType on container, to propagate desired behaviour ([#11](https://github.com/weka/weka-k8s-api/issues/11)) ([af75b24](https://github.com/weka/weka-k8s-api/commit/af75b24eb4f41584f1694f55ad6679e02da2a628))

## [1.2.1](https://github.com/weka/weka-k8s-api/compare/v1.2.0...v1.2.1) (2024-09-26)


### Bug Fixes

* drivers-dist in allowed values ([fd6ee7b](https://github.com/weka/weka-k8s-api/commit/fd6ee7b9ecd7f2c2f2b7e159bae3a9437f22bd39))

# [1.2.0](https://github.com/weka/weka-k8s-api/compare/v1.1.1...v1.2.0) (2024-09-26)


### Features

* policies should run only if pre-requisite policies have completed ([1a26f1c](https://github.com/weka/weka-k8s-api/commit/1a26f1c34d21a51cd9ae476376e2b227b7e89326))

## [1.1.1](https://github.com/weka/weka-k8s-api/compare/v1.1.0...v1.1.1) (2024-09-26)


### Bug Fixes

* rename new set of driver modes, keeping legacy intact ([5622d96](https://github.com/weka/weka-k8s-api/commit/5622d96a99ccfdf1a47893c1a498f5058cbd741d))

# [1.1.0](https://github.com/weka/weka-k8s-api/compare/v1.0.2...v1.1.0) (2024-09-26)


### Features

* upgrade policies for clients ([7c3371f](https://github.com/weka/weka-k8s-api/commit/7c3371f77f91fa86a24310f507b247c2b7d56a07))

## [1.0.2](https://github.com/weka/weka-k8s-api/compare/v1.0.1...v1.0.2) (2024-09-17)


### Bug Fixes

* IsDriversBuilder container mode condition ([959e2ee](https://github.com/weka/weka-k8s-api/commit/959e2eed8cb03235e0f08f7136d3b2ca1e89c6f9))

## [1.0.1](https://github.com/weka/weka-k8s-api/compare/v1.0.0...v1.0.1) (2024-09-16)


### Bug Fixes

* clean api - remove internal / unused logic ([63564c4](https://github.com/weka/weka-k8s-api/commit/63564c4e710e28364be0b485d7f33ac137585992))

# 1.0.0 (2024-09-15)


### Bug Fixes

* generate zz_ types from weka operator side ([5d0b9fd](https://github.com/weka/weka-k8s-api/commit/5d0b9fdea5f9f02ef8f8d95ae67928a7e6bed34f))


### Features

* new drivers flow integration ([1e98dd2](https://github.com/weka/weka-k8s-api/commit/1e98dd26dfb7fdfd5108d8cc053efe2b0ed45de9))
