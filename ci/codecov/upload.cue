package codecov

import (
	"dagger.io/dagger"

	"universe.dagger.io/docker"
)

#Upload: {
	// Source code
	source: dagger.#FS

	// Coverage files
	file: string

	// Codecov token (required for local runs and private repos)
	token?: dagger.#Secret

	_image: #Image

	_sourcePath: "/src"

	docker.#Run & {
		input: *_image.output | docker.#Image
		command: {
			name: "codecov"
			flags: {
				"--file": file
			}
		}
		env: {
			if token != _|_ {
				CODECOV_TOKEN: token
			}
		}
		workdir: _sourcePath
		mounts: {
			"source": {
				dest:     _sourcePath
				contents: source
			}
		}
	}
}
