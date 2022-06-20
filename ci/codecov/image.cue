package codecov

import (
	"universe.dagger.io/docker"
	// "universe.dagger.io/alpine"
	"universe.dagger.io/bash"
)

// Build a codecov base image
#Image: {
	_packages: [pkgName=string]: {
		// NOTE(samalba, gh issue #1532):
		//   it's not recommended to pin the version as it is already pinned by the major Alpine version
		//   version pinning is for future use (as soon as we support custom repositories like `community`,
		//   `testing` or `edge`)
		version: string | *""
	}
	_packages: {
		bash:         _
		curl:         _
		git:          _
		gnupg:        _
		coreutils:    _
		"perl-utils": _
	}
	docker.#Build & {
		steps: [
			// alpine.#Build & {
			//  packages: {
			//   bash:         _
			//   curl:         _
			//   gnupg:        _
			//   coreutils:    _
			//   "perl-utils": _
			//  }
			// },
			docker.#Pull & {
				source: "index.docker.io/alpine:3.15.0@sha256:21a3deaa0d32a8057914f36584b5288d2e5ecc984380bc0118285c70fa8c9300"
			},
			for pkgName, pkg in _packages {
				docker.#Run & {
					command: {
						name: "apk"
						args: ["add", "\(pkgName)\(pkg.version)"]
						flags: {
							"-U":         true
							"--no-cache": true
						}
					}
				}
			},
			bash.#Run & {
				script: contents: """
					curl https://keybase.io/codecovsecurity/pgp_keys.asc | gpg --no-default-keyring --keyring trustedkeys.gpg --import

					curl -Os https://uploader.codecov.io/latest/alpine/codecov

					curl -Os https://uploader.codecov.io/latest/alpine/codecov.SHA256SUM

					curl -Os https://uploader.codecov.io/latest/alpine/codecov.SHA256SUM.sig

					gpgv codecov.SHA256SUM.sig codecov.SHA256SUM

					shasum -a 256 -c codecov.SHA256SUM

					chmod +x codecov

					mv codecov /usr/local/bin
					rm codecov.SHA256SUM codecov.SHA256SUM.sig
					"""
			},
		]
	}
}
