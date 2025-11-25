{
    description = "Horizon - A Lisp interpreter made using Go and packaged using Nix";

    inputs = {
        nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
        flake-utils.url = "github:numtide/flake-utils";
    };

    outputs = { self, nixpkgs, flake-utils }:
        flake-utils.lib.eachDefaultSystem(system:
            let
                pkgs = nixpkgs.legacyPackages.${system};

                goPackagePath = "github.com/technik97/horizon";

                version = builtins.substring 0 8 self.lastModifiedDate;

                horizon = pkgs.buildGoModule {
                    pname = "horizon";
                    inherit version;

                    src = pkgs.lib.cleanSourceWith {
                        src = ./.;
                        filter = path: type:
                        !(pkgs.lib.hasSuffix ".nix" path) && 
                        !(pkgs.lib.hasInfix "/.github" path) && 
                        !(pkgs.lib.hasInfix "/testdata" path);  
                    };

                    vendorHash = "sha256-12Xmgwnk/QYfQuXs3WSCiV58BRFb9TiSfW5HtxoQbwc=";

                    meta = with pkgs.lib; {
                        description = "Horizon - A Lisp interpreter made using Go";
                        homepage = "https://github.com/technik97/horizon";
                        license = licenses.mit;
                        maintainers = with maintainers; [ Technik ];
                        mainProgram = "horizon";
                        platforms = platforms.linux ++ platforms.darwin;
                    };
                };

                container = pkgs.dockerTools.buildLayeredImage {
                    name = "horizon";
                    tag = version;

                    contents = [ horizon ];

                    config = {
                        Cmd = [ "horizon" ];  
                    };    
                };
            in {
                packages.horizon = horizon;
                packages.default = horizon;

                packages.container = container;

                apps.horizon = flake-utils.lib.mkApp { drv = horizon; };
                apps.default = self.apps.${system}.horizon;


                devShells.default = pkgs.mkShell {
                    name = "horizon-dev";
                    packages = with pkgs; [
                        go 
                        gopls
                        go-tools
                        golangci-lint  
                    ];

                    shellHook = ''
                        echo "Horizon dev shell - Go $(go version)"
                        export GOPATH=$PWD/.gopath
                        export GOBIN=$GOPATH/bin
                        export PATH=$PATH:$GOBIN
                        mkdir -p $GOPATH 
                    '';
                };
            });
}
