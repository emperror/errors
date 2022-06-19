package main

import (
	"dagger.io/dagger"
	"dagger.io/dagger/core"

	"universe.dagger.io/go"

	"github.com/sagikazarmark/dagger/go/golangci"
	"emperror.dev/errors/ci/codecov"
)

dagger.#Plan & {
	client: filesystem: ".": read: exclude: [
		".github",
		"bin",
		// "build",
		"tmp",
	]
	// client: filesystem: "./build": write: contents: actions.build.debug.output
	client: network: "unix:///var/run/docker.sock": connect: dagger.#Socket
	client: filesystem: "./build/test/coverage_1.18.out": write: contents: actions.check.test.go."1.18".export.files["/coverage.out"]
	client: env: {
		GITHUB_ACTIONS:    string | *""
		GITHUB_HEAD_REF:   string | *""
		GITHUB_REF:        string | *""
		GITHUB_REPOSITORY: string | *""
		GITHUB_RUN_ID:     string | *""
		GITHUB_SHA:        string | *""
		GITHUB_WORKFLOW:   string | *""
		CODECOV_TOKEN?:    dagger.#Secret
	}

	actions: {
		_source: client.filesystem["."].read.contents

		check: {
			test: {
				"go": {
					"1.14": _
					"1.15": _
					"1.16": _
					"1.17": _
					"1.18": _

					[v=string]: {
						_test: go.#Test & {
							source:  _source
							name:    "go_test_\(v)" // necessary to keep cache for different versions separate
							package: "./..."

							_image: go.#Image & {
								version: v
							}

							input: _image.output
							command: flags: {
								"-race":         true
								"-covermode":    "atomic"
								"-coverprofile": "/coverage.out"
							}

							export: files: "/coverage.out": _
						}

						_coverage: codecov.#Upload & {
							_write: core.#WriteFile & {
								input:    _source
								path:     "/coverage.out"
								contents: _test.export.files."/coverage.out"
							}

							source: _write.output
							file:   "/src/coverage.out"

							// Fixes https://github.com/dagger/dagger/issues/2680
							_env: client.env

							if _env.CODECOV_TOKEN != _|_ {
								token: _env.CODECOV_TOKEN
							}

							env: {
								GITHUB_ACTIONS:    client.env.GITHUB_ACTIONS
								GITHUB_HEAD_REF:   client.env.GITHUB_HEAD_REF
								GITHUB_REF:        client.env.GITHUB_REF
								GITHUB_REPOSITORY: client.env.GITHUB_REPOSITORY
								GITHUB_RUN_ID:     client.env.GITHUB_RUN_ID
								GITHUB_SHA:        client.env.GITHUB_SHA
								GITHUB_WORKFLOW:   client.env.GITHUB_WORKFLOW
							}
						}

						export: files: "/coverage.out": _test.export.files."/coverage.out"
					}
				}
			}

			lint: {
				"golangci": golangci.#Lint & {
					source:  _source
					version: "1.46"
				}
			}
		}
	}
}