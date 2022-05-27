{
  description = "Drop-in replacement for the standard library errors package";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    goflake.url = "github:sagikazarmark/go-flake";
    goflake.inputs.nixpkgs.follows = "nixpkgs";
    gobin.url = "github:sagikazarmark/go-bin-flake";
    gobin.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = { self, nixpkgs, flake-utils, goflake, gobin, ... }:
    flake-utils.lib.eachDefaultSystem (
      system:
        let
          pkgs = import nixpkgs {
            inherit system;

            overlays = [
              goflake.overlay

              (
                final: prev: {
                  golangci-lint = gobin.packages.${system}.golangci-lint-bin;
                  dagger = prev.buildGo118Module rec {
                    pname = "dagger";
                    version = "0.2.12";

                    src = prev.fetchFromGitHub {
                      owner = "dagger";
                      repo = "dagger";
                      rev = "v${version}";
                      sha256 = "sha256-t58+dfsf6A38RG4uT8SJPi07gkC9dGZo0WpVwN9N31k=";
                    };

                    vendorSha256 = "sha256-7YKuOApIw4SG39BLb4kh7ZuZjhCBduzKo3iS8v8KZZU=";

                    proxyModule = true;

                    subPackages = [
                      "cmd/dagger"
                    ];

                    ldflags = [ "-s" "-w" "-X go.dagger.io/dagger/version.Revision=${version}" ];
                  };
                }
              )
            ];
          };

          buildDeps = with pkgs; [ git go_1_17 gnumake ];
          devDeps = with pkgs; buildDeps ++ [
            golangci-lint
            gotestsum
            dagger
            go-task
          ];
        in
          { devShell = pkgs.mkShell { buildInputs = devDeps; }; }
    );
}
