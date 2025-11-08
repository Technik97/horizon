{
    description = "Development environment for Horizon";

    inputs = {
        nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
        flake-utils.url = "github:numtide/flake-utils";
    };

    outputs = { self, nixpkgs, flake-utils }:
        flake-utils.lib.eachDefaultSystem (system:
            let 
                pkgs = import nixpkgs {
                    inherit system;
                };
            in {
                devShells.default = pkgs.mkShell {
                    name = "horizon-dev-shell";

                    packages = with pkgs; [
                        go
                        go-tools
                        gopls
                    ];

                    shellHook = ''
                        echo "Entering Polaris dev shell..."
                        echo "Go: $(go version)"
                        export GOPATH=$PWD/.gopath
                        export GOBIN=$GOPATH/bin
                        export PATH=$PATH:$GOBIN
                        mkdir -p $GOPATH
                    '';
                };
            });
}
