{
  description = "gh-aipr: GitHub CLI extension to generate PR titles and descriptions";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    let 
      gh-aipr = { pkgs, ... }:
        pkgs.buildGoModule rec {
          pname = "gh-aipr";
          version = "unstable-" + (self.shortRev or "dev");

          src = ./.;

          vendorHash = "sha256-fUYTwKAjVV8g5Tc/bR1W2QoCownCRVr8zbZb77dg+YQ=";

          ldflags = [ "-s" "-w" ];
          env.CGO_ENABLED = 0;

          subPackages = [ "." ];
        };
    in
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        packages.default = gh-aipr { inherit pkgs; };

        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            gh
            git
            mise
          ];
          shellHook = ''
            echo "Entering gh-aipr dev shell (Go: $(go version | cut -d' ' -f3))"
          '';
        };

        formatter = pkgs.nixfmt-rfc-style;
      }
    ) // {
      overlays.pkgs = final: prev: {
        gh-aipr = prev.callPackage gh-aipr { pkgs = prev; };
      };
    };
}


