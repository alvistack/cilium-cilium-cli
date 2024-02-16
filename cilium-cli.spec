# Copyright 2024 Wong Hoi Sing Edison <hswong3i@pantarei-design.com>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

%global debug_package %{nil}

%global source_date_epoch_from_changelog 0

Name: cilium-cli
Epoch: 100
Version: 0.16.11
Release: 1%{?dist}
Summary: CLI to install, manage & troubleshoot Kubernetes clusters running Cilium
License: Apache-2.0
URL: https://github.com/cilium/cilium-cli/tags
Source0: %{name}_%{version}.orig.tar.gz
BuildRequires: golang-1.22
BuildRequires: glibc-static

%description
Cilium is open source software for providing and transparently securing
network connectivity and loadbalancing between application workloads
such as application containers or processes. Cilium operates at Layer
3/4 to provide traditional networking and security services as well as
Layer 7 to protect and secure use of modern application protocols such
as HTTP, gRPC and Kafka. Cilium is integrated into common orchestration
frameworks such as Kubernetes.

%prep
%autosetup -T -c -n %{name}_%{version}-%{release}
tar -zx -f %{S:0} --strip-components=1 -C .

%build
mkdir -p bin
set -ex && \
    export CGO_ENABLED=1 && \
    go build \
        -mod vendor -buildmode pie -v \
        -ldflags "-s -w \
            -X 'github.com/cilium/cilium-cli/internal/cli/cmd.Version=0.16.11' \
        " \
        -o ./cilium ./cmd/cilium

%install
install -Dpm755 -d %{buildroot}%{_bindir}
install -Dpm755 -d %{buildroot}%{_prefix}/share/bash-completion/completions
install -Dpm755 -t %{buildroot}%{_bindir}/ cilium
./cilium completion bash > %{buildroot}%{_prefix}/share/bash-completion/completions/cilium

%files
%license LICENSE
%{_bindir}/*
%{_prefix}/share/bash-completion/completions/*

%changelog
