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
