package main

import (
	"dagger.io/dagger"

	"universe.dagger.io/go"

	"github.com/sagikazarmark/dagger/go/golangci"
)

dagger.#Plan & {
	client: filesystem: ".": read: exclude: [
		".github",
		"bin",
		"build",
		"tmp",
	]
	// client: filesystem: "./build": write: contents: actions.build.debug.output
	client: network: "unix:///var/run/docker.sock": connect: dagger.#Socket

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

					[v=string]: go.#Test & {
						source:  _source
						name:    "go_test_\(v)"
						package: "./..."

						_image: go.#Image & {
							version: v
						}

						input: _image.output
						command: flags: "-race": true
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
