{
    description = "Jsonnet Extended Renderer";

    inputs = {
        nixpkgs.url = "nixpkgs/nixos-21.11";
        flake-utils.url = "github:numtide/flake-utils";
    };

    outputs = inputs@{ self, nixpkgs, flake-utils, ... }: let
        inherit (flake-utils.lib) eachDefaultSystem flattenTree;
    in eachDefaultSystem (system: let
        pkgs = nixpkgs.legacyPackages.${system};
    in rec {
        packages = flattenTree {
            jrender = let
                inherit (pkgs) lib buildGoModule;
            in pkgs.buildGoModule {
                pname = "jrender";
                version = "1.0.0";
                src = ./.;

                vendorSha256 = "sha256-nvp5hgVu/0VzdzeSgQsl1a4nvJ61JPisZTwM7OKLi0c=";
            };
        };

        defaultPackage = packages.jrender;
        defaultApp = packages.jrender;
    });
}
