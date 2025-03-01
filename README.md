# Octant <!-- Future Logo -->

[![Build Status](https://cloud.drone.io/api/badges/vmware/octant/status.svg)](https://cloud.drone.io/vmware/octant)
![GitHub release](https://img.shields.io/github/release/vmware/octant.svg)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

> A web-based, highly extensible platform for developers to better understand the complexity of Kubernetes clusters.

Octant is a tool for developers to understand how applications run on a Kubernetes cluster. It aims to be part of the developer's toolkit for gaining insight and approaching complexity found in Kubernetes. Octant offers a combination of introspective tooling, cluster navigation, and object management along with a plugin system to further extend its capabilities.

## Features

* **Resource Viewer**

    Graphically visualizate relationships between objects in a Kubernetes cluster. The status of individual objects are represented by color to show workload performance.

* **Summary View**

    Consolidated status and configuration information in a single page aggregated from output typically found using multiple kubectl commands.

* **Port Forward**

    Forward a local port to a running pod with a single button for debugging applications and even port forward multiple pods across namespaces.
 
* **Log Stream**

    View log streams of pod and container activity for troubleshooting or monitoring without holding multiple terminals open.

* **Label Filter**

    Organize workloads with label filtering for inspecting clusters with a high volume of objects in a namespace.

* **Cluster Navigation**

   Easily change between namespaces or contexts across different clusters. Multiple kubeconfig files are also supported.

 * **Plugin System**

   Highly extensible plugin system for users to provide additional functionality through gRPC. Plugin authors can add components on top of existing views.

## Usage

![Octant demo](docs/octant-demo.gif)

## Installation

<!-- TODO: brew and choco install -->
### Package (Linux only)

1. Download the `.deb` or `.rpm` from the [releases page](https://github.com/vmware/octant/releases).

2. Install with either `dpkg -i` or `rpm -i` respectively.

### Download a Pre-built Binary (Linux, macOS, Windows)

1. Open the [releases page](https://github.com/vmware/octant/releases) from a browser and download the latest tarball.

2. Extract the tarball:

    ```sh
    $ tar -xzvf ~/Downloads/octant_0.4.1_Linux-64bit.tar.gz
    octant_0.4.1_Linux-64bit/README.md
    octant_0.4.1_Linux-64bit/octant
    ```

3. Verify it runs:

    ```sh
    $ ./octant_0.4.1_Linux-64bit/octant version
    ```

## Getting Started

Before starting Octant, make sure you have access to a healthy cluster. If kubectl is installed, test using `kubectl cluster-info`.

Start running Octant:

`$ octant`

Octant should immediately launch your default web browser to an available high level port on 127.0.0.1. (i.e. http://127.0.0.1:51234)

For configuring Octant, setting up a development environment, or running tests, refer to the documentation [here](docs/getting-started.md).

## Plugins

Plugins are a core part of Octant in the Kubernetes ecosystem. A plugin can read objects and allows users to add components to Octant's views.

An example plugin can be found in [`cmd/octant-sample-plugin`](cmd/octant-sample-plugin) and installed to the default plugin path with `make install-test-plugin`.

Documentation for plugin components can be found in [`docs/plugins`](docs/plugins).

## Discussion

Feature requests, bug reports, and enhancements are welcome. Contributors, maintainers, and users are encouraged to collaborate through these communication channels:

 - [Twitter](https://twitter.com/projectoctant)
 - [Google group](https://groups.google.com/forum/#!forum/project-octant/)
 - [GitHub issues](https://github.com/vmware/octant/issues)

## Contributing

New contributors will need to sign a contributor license agreement before code changes can be merged. Follow the instructions given by `vmwclabot` after opening a pull request.

Pull requests should also include a changelog with the naming convention described [here](CONTRIBUTING.md).

See our [roadmap](ROADMAP.md) for tentative features in a 1.0 release.

## License

Octant is available under the [Apache License, Version 2.0](LICENSE)
