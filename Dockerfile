# Copyright 2020 OpenSSF Scorecard Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# golang:1.19
FROM golang@sha256:c416ceeec1cdf037b80baef1ccb402c230ab83a9134b34c0902c542eb4539c82 AS base
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download
COPY . ./

FROM base AS build
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 make build-scorecard

FROM gcr.io/distroless/base:nonroot@sha256:99133cb0878bb1f84d1753957c6fd4b84f006f2798535de22ebf7ba170bbf434
COPY --from=build /src/scorecard /
ENTRYPOINT [ "/scorecard" ]
