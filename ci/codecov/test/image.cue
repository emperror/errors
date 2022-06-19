package codecov

import (
	"dagger.io/dagger"

	"emperror.dev/errors/ci/codecov"
)

dagger.#Plan & {
	actions: test: codecov.#Image & {}
}
